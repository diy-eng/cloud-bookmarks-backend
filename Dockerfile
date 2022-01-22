FROM golang:alpine

MAINTAINER 92400538+diy-eng@users.noreply.github.com

ENV GIN_MODE=release
ENV PORT=3004

WORKDIR /go/src/app

COPY . /go/src/app/

RUN apk update && apk add --no-cache git
RUN go get ./...

RUN go build .

RUN ls -alh

EXPOSE $PORT

ENTRYPOINT ["./cloud-bookmarks"]