#!/bin/sh

# To terminate all services, run this script from the root directory of the project
INSIGHIT_STREAM_CONTAINER_NAME=curionoah-insight_stream-1
INFO_ISLAND_CONTAINER_NAME=curionoah-info_island-1


# terminate Insight Stream
docker exec -d $INSIGHIT_STREAM_CONTAINER_NAME /bin/bash -c "kill $(cat /usr/src/app/pidfile.txt)" &&

# terminate Info Island
docker exec -d $INFO_ISLAND_CONTAINER_NAME /bin/bash -c "kill $(cat /usr/src/app/vite_pid.txt)" &&


# terminate FeedHarmony
kill $(lsof -t -i:5100) &&

echo "All services are terminated"
