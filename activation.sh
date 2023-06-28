#!/bin/sh

INSIGHIT_STREAM_CONTAINER_NAME=curionoah-insight_stream-1
INFO_ISLAND_CONTAINER_NAME=curionoah-info_island-1

# run Front end and Back end apps
docker exec -d $INSIGHIT_STREAM_CONTAINER_NAME /bin/bash -c "cd /usr/src/app && go run main.go & echo $! > ./pidfile.txt" &&
echo "Insight Stream is running" &&

docker exec -d $INFO_ISLAND_CONTAINER_NAME /bin/bash -c "cd /usr/src/app && npm install && npm run build && npm run preview &echo $! > ./vite_pid.txt" &&
echo "Info Island is running" &&

# run microservices
cd ./FeedHarmony/feed_harmony && cargo build --release && ./target/release/feed_harmony &

echo "All services are running" &&
