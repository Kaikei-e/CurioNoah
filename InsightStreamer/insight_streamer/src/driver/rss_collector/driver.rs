use std::error::Error;
use rss::Channel;
use std::net::TcpStream;
use hyper::Request;
use bytes::Bytes;
use tokio::io::{self, AsyncWriteExt as _};

async fn fetch_rss(url: hyper::Uri) -> Result<Channel, Box<dyn Error>> {
    todo!("implement fetch_rss")

}