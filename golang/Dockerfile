FROM golang:1.11

ARG GOLANG_BINARY=scale17x
ENV GB=$GOLANG_BINARY

WORKDIR /home/
ADD . /home/

RUN go get github.com/UnnoTed/fileb0x && \
    fileb0x extmodules.yaml && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${GOLANG_BINARY} -a -ldflags '-extldflags "-static"' main.go

ENTRYPOINT mv /home/${GB} /tmp/${GB}