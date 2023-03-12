use axum::{routing::{get, post},  Router};
use crate::fiber::server::summarizer::presentation::summarize_api::summarizer;


#[tokio::main]
pub async fn main() {
    let app = Router::new()
        .route("/", get(|| async { "Hello, World!" }))
        .route("/api/v1/textSummarize", post(summarizer));

    // run it with hyper on localhost:3000
    axum::Server::bind(&"0.0.0.0:5020".parse().unwrap())
        .serve(app.into_make_service())
        .await
        .unwrap();
}

