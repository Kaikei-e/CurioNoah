use crate::driver::repository::rss_feeds_driver::DatabasePool;
use crate::usecase::feeds::fetch::follow_list_by_twenty;
use axum::extract::State;
use axum::http::StatusCode;
use axum::response::IntoResponse;

pub async fn fetch(State(pool): State<DatabasePool>) -> impl IntoResponse {
    let result = follow_list_by_twenty(pool).await;

    if let Err(e) = &result {
        println!("Error: {}", e);
        todo!("implement error handling")
    }

    match result {
        Ok(result) => (StatusCode::OK, serde_json::to_string(&result).unwrap()),
        Err(e) => (
            StatusCode::INTERNAL_SERVER_ERROR,
            serde_json::to_string(&e.to_string()).unwrap(),
        ),
    }
}
