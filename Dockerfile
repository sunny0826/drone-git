# build step
FROM golang:1.12 as builder

LABEL maintainer="sunnydog0826@gmail.com"
COPY . /build/

WORKDIR /build

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

# run step
FROM drone/git

# copy bin from build step
COPY --from=builder /build/drone-git /bin/

ENTRYPOINT /bin/drone-git