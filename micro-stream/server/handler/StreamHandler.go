package handler

import (
	"context"
	ss "go-micro/micro-stream/proto"
	"log"
	"os"
	"time"
)

type StreamHandler struct {
}

// 双向通信
func (s StreamHandler) GetStream1(ctx context.Context, stream ss.SteamService_GetStream1Stream) error {

	logger := log.New(os.Stdout, "[Server] ", 1)

	recv, err := stream.Recv()
	logger.Println("服务端收到数据", recv, err)

	stream.Send(&ss.StreamResponse{
		Id:      112321,
		Message: "From Server",
	})

	return nil
}

// 服务端到客户端
func (s StreamHandler) GetStream2(ctx context.Context, request *ss.StreamRequest, stream ss.SteamService_GetStream2Stream) error {

	logger := log.New(os.Stdout, "[Server] ", 1)
	logger.Println("服务端接收数据", request)

	ch := make(chan int)
	go func() {
		for {
			ch <- 1
			time.Sleep(time.Second * 5)
		}
	}()

	for {
		i := <-ch
		stream.Send(&ss.StreamResponse{
			Id:      int32(i),
			Message: "From Server",
		})
	}
	return nil

}

// 客户端到服务端
func (s StreamHandler) GetStream3(ctx context.Context, stream ss.SteamService_GetStream3Stream) error {

	logger := log.New(os.Stdout, "[Server] ", 1)
	recv, err := stream.Recv()
	logger.Println("服务端接收数据", recv, err)

	return nil
}
