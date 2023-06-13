use crate::api_handler::handler::DatabasePool;
use crate::usecase::store_word_network_usecase::store::store_word_net_usecase;
use axum::extract::State;
use axum::http::StatusCode;
use axum::response::IntoResponse;
use axum::Json;
use serde_json::json;

pub async fn group_and_store_by_url(State(pool): State<DatabasePool>) -> impl IntoResponse {
    let result = store_word_net_usecase(pool).await;
    match result {
        Ok(_) => {
            println!("Grouped and stored all urls");
            (
                StatusCode::OK,
                Json(json!({"message": "Grouped and stored all urls"})),
            )
        }
        Err(e) => {
            println!("Failed to group and store all urls: {}", e);
            let res_body = json!({"error": e.to_string()});
            (StatusCode::INTERNAL_SERVER_ERROR, Json(res_body))
        }
    }
}
