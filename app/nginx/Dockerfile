FROM nginx:1.16.1-alpine

# remove default conf
RUN rm /etc/nginx/conf.d/default.conf

# create my conf
COPY ./nginx.conf /etc/nginx/nginx.conf
COPY ./sssignal3.conf /etc/nginx/conf.d/sssignal3.conf

# set Japanese time
RUN apk --update add tzdata
RUN cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
RUN apk del tzdata
RUN rm -rf /var/cache/apk/*
