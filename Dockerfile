# 编译阶段
FROM golang:1.12
LABEL maintainer="sunnydog0826@gmail.com"
COPY . /build/

WORKDIR /build

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOARM=6 go build .

# 运行阶段
FROM alpine

# 从编译阶段的中拷贝编译结果到当前镜像中
COPY --from=0 /build/drone-git /bin

FROM alpine

LABEL maintainer="sunnydog0826@gmail.com"


RUN apk update \
    && apk add --no-cache bash git \
    && rm -rf /var/cache/apk/* \
    && chmod +x /bin/drone-git \

#ADD drone-git /bin/
ENTRYPOINT ["/bin/drone-git"]