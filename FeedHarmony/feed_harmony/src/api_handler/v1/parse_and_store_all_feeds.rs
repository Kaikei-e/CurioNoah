use crate::api_handler::handler::DatabasePool;
use crate::usecase::parse_and_store::parse_and_store_feeds;
use axum::http::StatusCode;
use axum::response::IntoResponse;
use axum::{extract::State, http::header, http::Response, response, Json};
use serde_json::json;

pub async fn parse_and_store(State(pool): State<DatabasePool>) -> impl IntoResponse {
    let result = parse_and_store_feeds(pool).await;
    match result {
        Ok(_) => {
            println!("Parsed and stored all feeds");
            (
                StatusCode::OK,
                Json(json!({"message": "Parsed and stored all feeds"})),
            )
        }
        Err(e) => {
            println!("Failed to parse and store all feeds: {}", e);
            let res_body = json!({"error": e.to_string()});
            (StatusCode::INTERNAL_SERVER_ERROR, Json(res_body))
        }
    }
}
