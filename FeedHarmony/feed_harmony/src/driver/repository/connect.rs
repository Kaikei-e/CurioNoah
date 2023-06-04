use sqlx::mysql::MySqlPool;

#[tokio::main]
pub async fn initialize_connection(var: String) -> anyhow::Result<MySqlPool, sqlx::Error> {
    let pool = MySqlPool::connect(&var).await;

    match pool {
        Ok(_) => {
            println!("Connected to database");
            pool
        }
        Err(e) => {
            panic!("Failed to connect to database: {}", e);
        }
    }
}
