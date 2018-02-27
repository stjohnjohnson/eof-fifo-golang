FROM golang:1.10
WORKDIR /go/src/app
COPY . .
ENTRYPOINT /go/src/app/test.sh
