FROM golang:1.15.15-alpine

RUN apk add --no-cache git

WORKDIR /go
RUN go get -u github.com/labstack/echo/...
RUN go get github.com/joho/godotenv
RUN go get github.com/oxequa/realize
RUN go get gopkg.in/mgo.v2
RUN go get go.mongodb.org/mongo-driver/mongo
RUN mkdir /go/cmd
COPY . ./cmd/sssignal3-api

WORKDIR /go/cmd/sssignal3-api
RUN go build -o /go/bin/sssignal3-api
