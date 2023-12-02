FROM golang:1.21.1-bullseye

ADD ./InsightStream /usr/src/app

RUN cd /usr/src/app && go run main.go

EXPOSE 9000