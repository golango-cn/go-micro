package tests

import (
	"fmt"
	"testing"
)

type functionModel struct {
	ints []int
	name string
}

type funcOptions func(m *functionModel)

func withInts(i int) funcOptions  {
	return func(m *functionModel) {
		m.ints = append(m.ints, i)
	}
}

func withName(name string) funcOptions  {
	return func(m *functionModel) {
		m.name = name
	}
}

func TestFunction(t *testing.T)  {

	m := &functionModel{}
	for _, opt := range []funcOptions{ withInts(1), withInts(2), withName("helloworld"), } {
		opt(m)
	}
	fmt.Println(m)

}
