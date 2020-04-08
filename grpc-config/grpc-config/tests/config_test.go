package tests

import (
	"fmt"
	"github.com/micro/go-plugins/config/source/grpc/v2"
	"testing"
)
import _ "github.com/micro/go-micro/config/source"

func TestConfig01(t *testing.T) {

	fmt.Println("123")

	source := grpc.NewSource(
		grpc.WithAddress("127.0.0.1:9600"),
		grpc.WithPath("micro"),
	)

	s1 := WithSource(source)
	s2 := WithApp("123")

	var opts []Option
	opts = append(opts, s1)
	opts = append(opts, s2)

	ops := Options{}
	for _, o := range opts {
		o(&ops)
	}

	fmt.Println(ops)

}
