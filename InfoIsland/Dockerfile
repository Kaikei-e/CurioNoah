FROM node:20.6.1-bullseye

ADD ./InfoIsland /usr/src/app

WORKDIR /usr/src/app

RUN npm install -g vite

RUN cd /usr/src/app && npm install && npm run build

EXPOSE 4173 5173