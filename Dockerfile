
FROM golang:1.7-alpine as builder

RUN apk --update add --no-cache --virtual .build-deps \
    gcc libc-dev linux-headers

ENV SRC=/go/src/github.com/infobloxopen/container-ipam-tool

COPY . ${SRC}
WORKDIR ${SRC}

RUN go build -o bin/create-ea-defs ./ea-defs


FROM alpine:3.5

ENV SRC=/go/src/github.com/infobloxopen/container-ipam-tool

COPY launcher/launch.sh launch.sh
COPY --from=builder ${SRC}/bin/create-ea-defs /usr/local/bin/create-ea-defs

ENTRYPOINT ["/launch.sh"]

ARG GIT_SHA
ARG BUILD_DATE

LABEL GIT_SHA=$GIT_SHA \
      BUILD_DATE=$BUILD_DATE
