#!/bin/sh

INSIGHIT_STREAM_CONTAINER_NAME=curionoah-insight_stream-1
INFO_ISLAND_CONTAINER_NAME=curionoah-info_island-1

# get the current directory
CURRENT_DIR=$(pwd)

# run Front end and Back end apps
docker exec -d $INSIGHIT_STREAM_CONTAINER_NAME /bin/bash -c "cd /usr/src/app && go run main.go" &&
echo "Insight Stream is running" &&

docker exec -d $INFO_ISLAND_CONTAINER_NAME /bin/bash -c "cd /usr/src/app && npm install && npm run build && npm run preview" &&
echo "Info Island is running" &&

# clean up the log file
echo "" > $CURRENT_DIR/Exec/feed_harmony.log &&

# run microservices
(
  cd ./FeedHarmony/feed_harmony &&
  cargo build --release &&
  ./target/release/feed_harmony > $CURRENT_DIR/Exec/feed_harmony.log 2>&1
) &&
echo "FeedHarmony is running" &&

echo "All services are running"

# follow the output of FeedHarmony
tail -f $CURRENT_DIR/Exec/feed_harmony.log