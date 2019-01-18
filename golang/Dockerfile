FROM golang:1.11

WORKDIR /home/
ADD . /home/

RUN go get github.com/UnnoTed/fileb0x && fileb0x extmodules.yaml
RUN go mod init github.com/scale17x/golang && \
    GOOS=linux; GOARCH=amd64; CGO_ENABLED=0; go build -o scale17x main.go
