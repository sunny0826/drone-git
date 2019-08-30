FROM alpine

LABEL maintainer="sunnydog0826@gmail.com"

RUN apk update \
    && apk add --no-cache bash git \
    && rm -rf /var/cache/apk/*

ADD bin/drone-git /bin/

ENTRYPOINT ["/bin/drone-git"]