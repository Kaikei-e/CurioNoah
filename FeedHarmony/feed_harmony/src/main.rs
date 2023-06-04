use std::env;

mod api_handler;
mod domain;
mod driver;

fn main() {
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
    let pool = driver::repository::connect::initialize_connection(var);

    match pool {
        Ok(_) => {
            println!("Connected to database");
        }
        Err(e) => {
            println!("Failed to connect to database: {}", e);
        }
    }

    api_handler::handler::handler();
}
