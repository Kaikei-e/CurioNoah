use axum::{
    routing::get,
    Router,
};


#[tokio::main]
pub async fn handler(){
    let app = Router::new().route("/", get(|| async { "Hello, World!" }));

    axum::Server::bind(&"0.0.0.0:5100".parse().unwrap())
        .serve(app.into_make_service())
        .await
        .unwrap();
}