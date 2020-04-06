FROM golang:1.14
WORKDIR /go/src/app

copy . .

RUN go env -w GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE on

RUN go get -d -v ./...
RUN go build -o app
RUN chmod +x ./app

EXPOSE 9600/tcp

CMD ./app