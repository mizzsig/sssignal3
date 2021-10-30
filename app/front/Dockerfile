FROM node:16.13.0-alpine

WORKDIR /home/node
RUN mkdir front
COPY . ./front
WORKDIR /home/node/front
RUN yarn --pure-lockfile
