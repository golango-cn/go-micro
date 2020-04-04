package learning

import (
	"fmt"
	learning "go-micro/rpc2http/proto"
	"golang.org/x/net/context"
)

type Handler struct {

}

func (h Handler) Hi(ctx context.Context, requ *learning.Request, resp *learning.Response) error {

	resp.Msg = fmt.Sprintf( "Hello %s, Welcome to Learning from V2!!!", requ.Name)
	return nil
	
}

