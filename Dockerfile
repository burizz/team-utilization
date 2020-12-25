# Build stage
FROM golang:1.15-alpine AS build

WORKDIR /go/src/github.com/burizz/team-utilization

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go get -d -v ./...

COPY .env .env
COPY ./config/ ./config/
COPY ./server/ ./server/
COPY ./seed/ ./seed/
COPY ./storage/ ./storage/
COPY ./teams/ ./teams/

RUN go install -v ./...

# Run stage

FROM alpine:latest

WORKDIR /usr/local/bin

COPY --from=build /go/bin/utilization .

CMD [ "utilization" ]
