FROM alpine:latest
LABEL maintainer "ckatsak@gmail.com"

RUN apk update && apk add --no-cache rsync

ADD ./disttate_static /

ENTRYPOINT ["/disttate_static"]
