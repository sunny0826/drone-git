# build step
FROM golang:1.12 as builder

LABEL maintainer="sunnydog0826@gmail.com"
COPY . /build/

WORKDIR /build

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

# run step
FROM alpine:3.7

RUN apk add --no-cache ca-certificates git git-lfs openssh curl perl

ADD script/clone.sh /bin/

# copy bin from build step
COPY --from=builder /build/drone-git /bin/

ENTRYPOINT ["/bin/clone.sh","/bin/drone-git"]