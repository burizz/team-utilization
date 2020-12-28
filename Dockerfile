### Build stage
FROM golang:1.15-alpine AS go-build
#RUN apk add --no-cache git

ENV SRC_DIR=/go/src/github.com/burizz/team-utilization

WORKDIR $SRC_DIR

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go get -d -v ./...

COPY .env .env
COPY ./config/ ./config/
COPY ./server/ ./server/
COPY ./storage/ ./storage/
COPY ./teams/ ./teams/
COPY ./seed/ ./seed/

# Run Unit tests
#RUN CGO_ENABLED=0 go test -v test/tests.go

# Build binary
RUN go build -o bin/utilization-server server/utilization.go
# RUN go install -v ./...

### Run stage
FROM alpine:latest
#RUN apk add ca-certificates

ENV ENV_TYPE=DOCKER
ENV SRC_DIR=/go/src/github.com/burizz/team-utilization

WORKDIR /usr/local/bin

COPY --from=go-build ${SRC_DIR}/seed ./seed/
COPY --from=go-build ${SRC_DIR}/.env .
COPY --from=go-build ${SRC_DIR}/bin/utilization-server .

RUN ls -alh && pwd

#EXPOSE 8080

CMD [ "utilization-server" ]
