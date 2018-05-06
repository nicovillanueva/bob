FROM golang:1.9.5-alpine3.6
RUN mkdir /app && \
    apk add --no-cache git && \
    go get github.com/golang/dep/cmd/dep
ADD . /go/src/app/
WORKDIR /go/src/app/
RUN dep ensure && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
COPY --from=0 /go/src/app/main /
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
CMD ["/main"]