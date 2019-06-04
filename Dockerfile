FROM golang:1.12-alpine as base
RUN apk add --no-cache libstdc++ gcc g++ git ca-certificates linux-headers
WORKDIR /go/src/github.com/hunterlong/smartedge
ADD . .
RUN go get && go install

FROM alpine:latest
COPY --from=base /go/bin/smartedge /usr/local/bin/smartedge
WORKDIR /app
VOLUME /app
CMD smartedge "hello world"
