FROM node:22.12.0-bullseye AS base

ADD ./InfoIsland /usr/src/app

WORKDIR /usr/src/app

RUN npm install -g vite

RUN cd /usr/src/app && npm install && npm run build

FROM node:22.12.0-bullseye AS production

COPY --from=base /usr/src/app /usr/src/app

RUN npm install serve

EXPOSE 4173 5173