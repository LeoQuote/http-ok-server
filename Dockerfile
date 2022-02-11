FROM golang:alpine as build-env

RUN apk add git

# Copy source + vendor
COPY . /go/src/github.com/tekkamanendless/http-ok-server
WORKDIR /go/src/github.com/tekkamanendless/http-ok-server

# Build
ENV GOPATH=/go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -v -a -ldflags "-s -w" -o /go/bin/http-ok-server .

FROM scratch
COPY --from=build-env /go/bin/http-ok-server /usr/bin/http-ok-server
ENTRYPOINT ["http-ok-server"]
