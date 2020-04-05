package test

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T)  {

	mp := make(map[interface{}]interface{})
	mp["userId"] = 123

	fmt.Println(mp["userId"])

}
