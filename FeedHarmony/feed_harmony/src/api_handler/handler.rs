use crate::api_handler;
use axum::extract::State;
use axum::routing::post;
use axum::{async_trait, routing::get, Router};
use sqlx::{MySql, Pool};

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
        .route("/", get(|| async { "Hello, World!" }))
        .route(
            "/api/v1/store_feeds",
            post(api_handler::v1::store_feeds::store_feeds),
        )
        .with_state(state.pool);

    axum::Server::bind(&"0.0.0.0:5100".parse().unwrap())
        .serve(app.into_make_service())
        .await
        .unwrap();
}
