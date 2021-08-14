FROM node:14.17.5-alpine

WORKDIR /home/node
RUN mkdir front
COPY . ./front
WORKDIR /home/node/front
RUN yarn --pure-lockfile
