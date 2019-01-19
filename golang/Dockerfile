FROM golang:1.11

ARG GOLANG_BINARY=scale17x

WORKDIR /home/
ADD . /home/

RUN go get github.com/UnnoTed/fileb0x && \
    fileb0x extmodules.yaml && \
    GOOS=linux GOARCH=amd64 go build -o ${GOLANG_BINARY} main.go