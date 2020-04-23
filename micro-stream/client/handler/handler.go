package handler

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/client"
	ss "go-micro/micro-stream/proto"
	"log"
	"net/http"
	"os"
)

// 双向通信
func Stream1(w http.ResponseWriter, r *http.Request) {

	logger := log.New(os.Stdout, "[Client] ", 1)

	svr := ss.NewSteamService("go.micro.golango.svr", client.DefaultClient)
	stream1, _ := svr.GetStream1(context.Background())

	stream1.Send(&ss.StreamRequest{
		Id:        12323,
		FirstName: "From Client",
	})

	recv, err := stream1.Recv()
	logger.Println("客户端接收数据", recv, err)

	fmt.Fprintln(w, "操作成功")

}

// 服务端到客户端
func Stream2(w http.ResponseWriter, r *http.Request) {

	logger := log.New(os.Stdout, "[Client] ", 1)

	svr := ss.NewSteamService("go.micro.golango.svr", client.DefaultClient)
	stream1, _ := svr.GetStream2(context.Background(), &ss.StreamRequest{
		Id:        123,
		FirstName: "From Client",
	})

	recv, err := stream1.Recv()
	logger.Println("客户端接收数据", recv, err)

	fmt.Fprintln(w, "操作成功")

}

// 客户端到服务端
func Stream3(w http.ResponseWriter, r *http.Request) {

	logger := log.New(os.Stdout, "[Client] ", 1)

	svr := ss.NewSteamService("go.micro.golango.svr", client.DefaultClient)
	stream1, _ := svr.GetStream3(context.Background())

	err := stream1.Send(&ss.StreamRequest{
		Id:        123,
		FirstName: "From Client",
	})

	if err != nil {
		logger.Fatal(err)
	}

	fmt.Fprintln(w, "操作成功")

}
