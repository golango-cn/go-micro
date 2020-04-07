FROM golang:1.14 as builder

WORKDIR /go/src/app
COPY ./  ./
RUN export GO111MODULE=on \
 && go env -w GOPROXY=https://goproxy.cn,direct \
 && go get -d -v ./... \
 && CGO_ENABLED=0 GOOS=linux go build -o app main.go


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /go/src/app/app ./
COPY --from=builder /go/src/app/conf ./conf

EXPOSE 9600/tcp

#ENTRYPOINT ["/root/app"]
CMD ./app