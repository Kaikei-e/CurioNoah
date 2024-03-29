use crate::driver::repository::rss_feeds_driver::*;
use crate::rest::v1::fetch::fetch;
use axum::routing::get;
use axum::Router;
use axum::Server;
use std::net::SocketAddr;

#[derive(Clone)]
struct AppState {
    pool: DatabasePool,
}

pub async fn api_server() {
    let pool = RSSFeedRepository::initialize_connection().await;

    let state = AppState {
        pool: DatabasePool { pool },
    };

    let app = Router::new()
        .route("/system/alive", get(|| async { "Keep alive" }))
        .route("/v1/fetch", get(fetch))
        .with_state(state.pool);

    let port = 9200;
    let host = "0.0.0.0";

    let addr = format!("{}:{}", host, port);
    println!("Starting server on {}", addr);

    let addr_parsed: Result<SocketAddr, _> = addr.parse();

    if let Err(e) = addr_parsed {
        println!("Error parsing address: {}", e);
        return;
    }

    Server::bind(&addr_parsed.unwrap())
        .serve(app.into_make_service())
        .await
        .unwrap();

    println!("Server stopped");
}
