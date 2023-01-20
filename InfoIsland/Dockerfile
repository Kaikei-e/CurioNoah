FROM node:18.13

ADD ./InfoIsland /usr/src/app

RUN yarn set version stable && \
  cd /usr/src/app && \
  yarn install && \
  yarn build

EXPOSE 5173