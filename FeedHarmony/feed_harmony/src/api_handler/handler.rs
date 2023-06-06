use crate::api_handler;
use axum::routing::post;
use axum::{routing::get, Router};

#[tokio::main]
pub async fn handler() {
    let app = Router::new()
        .route("/", get(|| async { "Hello, World!" }))
        .route(
            "/store_feeds",
            post(api_handler::v1::store_feeds::store_feeds),
        );

    axum::Server::bind(&"0.0.0.0:5100".parse().unwrap())
        .serve(app.into_make_service())
        .await
        .unwrap();
}
