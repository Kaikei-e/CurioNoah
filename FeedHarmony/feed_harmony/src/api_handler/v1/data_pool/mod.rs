mod update_word_pool;

use crate::api_handler::handler::DatabasePool;
use axum::extract::State;

pub async fn make_word_pool(State(pool): State<DatabasePool>) -> impl IntoResponse {
    todo!()
}
