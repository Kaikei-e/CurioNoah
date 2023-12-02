FROM node:20.6.1-bullseye

ADD ./InfoIsland /usr/src/app

RUN cd /usr/src/app && npm install && npm run build

RUN npm run preview

EXPOSE 5173 4173