use crate::api_handler;
use axum::routing::post;
use axum::Router;
use sqlx::{MySql, Pool};
use std::net::SocketAddr;

#[derive(Clone)]
pub struct DatabasePool {
    pub(crate) pool: Pool<MySql>,
}

#[derive(Clone)]
struct AppState {
    pool: DatabasePool,
}

pub async fn handler(pool: Pool<MySql>) {
    let state = AppState {
        pool: DatabasePool { pool },
    };

    let app = Router::new()
        .route(
            "/api/v1/store_feeds",
            post(api_handler::v1::store_feeds::fetch_all),
        )
        .route(
            "/api/v1/parse_and_store_feeds",
            post(api_handler::v1::parse_and_store_all_feeds::parse_and_store),
        )
        .route(
            "/api/v1/parse_and_store_latest_feeds",
            post(api_handler::v1::parse_and_store_latest_feeds::parse_and_store_latest),
        )
        .route(
            "/api/v1/group_and_store_by_url",
            post(api_handler::v1::data_pool::update_word_pool::group_and_store_by_url),
        )
        .with_state(state.pool);

    let addr = SocketAddr::from(([0, 0, 0, 0], 5100));

    axum::serve(tokio::net::TcpListener::bind(addr).await.unwrap(), app)
        .await
        .expect("Failed to start the server");
}
