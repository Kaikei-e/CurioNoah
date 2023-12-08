FROM rust:1.74-bullseye

# Install dependencies
ADD ./feed_oracle /usr/src/app

RUN apt-get update && apt-get install -y \
curl build-essential

RUN touch /usr/src/app/feed.log

WORKDIR /usr/src/app

EXPOSE 5110