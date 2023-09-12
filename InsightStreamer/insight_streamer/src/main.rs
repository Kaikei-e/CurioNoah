use crate::rest::server::api_server;
use std::env;

mod domain;
mod rest;
mod usecase;
mod driver;

#[tokio::main]
async fn main() {
    let loaded = dotenvy::dotenv();

    match loaded {
        Ok(_) => {
            println!("Loaded .env file");
        }
        Err(e) => {
            panic!("Failed to load .env file: {}", e)
        }
    }
    let var = env::var("DATABASE_URL").expect("DATABASE_URL must be set");
    let pool = sqlx::mysql::MySqlPool::connect(&var).await.unwrap();

    api_server(pool).await;

    println!("Hello, world!");
}
