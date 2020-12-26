### Build stage
FROM golang:1.15-alpine AS go-build
#RUN apk add --no-cache git

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

# Run Unit tests
#RUN CGO_ENABLED=0 go test -v test/tests.go

# Build binary
RUN go build -o bin/utilization-server server/utilization.go
# RUN go install -v ./...

### Run stage
FROM alpine:latest
#RUN apk add ca-certificates

WORKDIR /usr/local/bin

COPY --from=go-build /go/src/github.com/burizz/team-utilization/bin/utilization-server .

RUN ls -alh && pwd

CMD [ "utilization-server" ]
