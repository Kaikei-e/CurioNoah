FROM rust:1.79-bookworm

# Install dependencies
ADD ./feed_harmony /usr/src/app

RUN apt-get update && apt-get install -y \
curl build-essential

RUN touch /usr/src/app/harmony.log

WORKDIR /usr/src/app

ENV SQLX_OFFLINE true

EXPOSE 5100