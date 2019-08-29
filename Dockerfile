FROM alpine

LABEL maintainer="sunnydog0826@gmail.com"

RUN apk update \
    && apk add --no-cache bash git curl \
    && rm -rf /var/cache/apk/*
    && curl -L https://github.com/sunny0826/drone-git/releases/download/v0.0.1/drone-git_0.0.1_Linux_x86_64.tar.gz -o /usr/local \
    && cd /usr/local \
    && tar zxvf drone-git_0.0.1_Linux_x86_64.tar.gz \
    && cd drone-git_0.0.1_Linux_x86_64 \
    && mv drone-git /bin/ \
    && chmod +x /bin/drone-git \

#ADD drone-git /bin/
ENTRYPOINT ["/bin/drone-git"]