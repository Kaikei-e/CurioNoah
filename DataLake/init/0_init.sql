CREATE DATABASE IF NOT EXISTS datalake;

CREATE SCHEMA IF NOT EXISTS news_generation;

CREATE TABLE news_generation.todays_news (
    news_id uuid PRIMARY KEY,
    news_title text,
    news_content text,
    news_date date
    search_vector tsvector,
    index search_vector_idx USING gin(search_vector)
    index news_date_title_content_idx (news_date, news_title, news_content)
);

UPDATE news_generation.todays_news SET search_vector = to_tsvector(news_title || ' ' || news_content);

