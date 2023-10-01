use rss::Channel;

async fn fetch_rss(url: url::Url) -> Result<Channel, anyhow::Error> {
    let resp = reqwest::blocking::get(url)?.text()?;
    let channel = resp.parse::<rss::Channel>()?;

    Ok(channel)
}