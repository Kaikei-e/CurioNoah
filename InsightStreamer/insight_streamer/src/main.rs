use crate::rest::server::api_server;

mod domain;
mod driver;
mod rest;
mod usecase;

#[tokio::main]
async fn main() {
    api_server().await;

    println!("Hello, world!");
}
