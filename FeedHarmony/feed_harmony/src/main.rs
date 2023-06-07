use std::env;

mod api_handler;
mod domain;
mod driver;
mod usecase;

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
    let pool = driver::repository::feed_db::connect::initialize_connection(var).await;

    match pool {
        Ok(pool) => {
            println!("Established connection to database");
            api_handler::handler::handler(pool);
        }
        Err(e) => {
            panic!("Failed to initialize application : {}", e)
        }
    }
}
