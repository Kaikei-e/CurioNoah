#!/bin/sh

# To terminate all services, run this script from the root directory of the project
INSIGHIT_STREAM_CONTAINER_NAME=curionoah-insight_stream-1
INFO_ISLAND_CONTAINER_NAME=curionoah-info_island-1

# get the current directory
CURRENT_DIR=$(pwd)

# terminate Insight Stream
docker restart $INSIGHIT_STREAM_CONTAINER_NAME &&
echo "container curionoah-insight_stream-1 is terminated" &&

# terminate Info Island
docker restart $INFO_ISLAND_CONTAINER_NAME &&
echo "container curionoah-info_island-1 is terminated" &&

# terminate FeedHarmony
docker restart curionoah-feed_harmony-1 &&
echo "container curionoah-feed_harmony-1 is terminated" &&

echo "All services are terminated"
