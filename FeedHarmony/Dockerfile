FROM rust:1.74-bullseye

# Install dependencies
ADD ./feed_harmony /usr/src/app

RUN touch /usr/src/app/harmony.log

RUN cd /usr/src/app && cargo build --release
RUN chmod +x /usr/src/app/target/release/feed_harmony

# For development, now this section is disabled
RUN ./target/release/feed_harmony > harmony.log 2>&1 &&\
echo "FeedHarmony started"

EXPOSE 5100
