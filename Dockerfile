# Compile stage
FROM golang:bullseye AS build-env

# Build Delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest

ADD . /dockerdev
WORKDIR /dockerdev
RUN go build -gcflags="all=-N -l" -o /server

#final stage
FROM busybox:glibc

EXPOSE 8000 40000

WORKDIR /
COPY --from=build-env /go/bin/dlv /
COPY --from=build-env /server /

CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/server"]