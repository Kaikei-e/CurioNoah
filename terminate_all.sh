#!/bin/sh

# To terminate all services, run this script from the root directory of the project

# terminate FeedHarmony
kill $(lsof -t -i:5100)
