use crate::api_handler::handler::DatabasePool;
use crate::usecase;
use crate::usecase::parse_and_store::parse_and_store_feeds;
use axum::extract::State;
use axum::response::IntoResponse;
use axum::Json;

pub async fn parse_and_store(State(pool): State<DatabasePool>) -> impl IntoResponse {
    let result = parse_and_store_feeds(pool).await;
    match result {
        Ok(_) => {
            println!("Parsed and stored all feeds");
            Json("Parsed and stored all feeds")
        }
        Err(e) => {
            println!("Failed to parse and store all feeds: {}", e);
            todo!("Return error response")
        }
    }
}
