package main

import "github.com/lbuckalew/blulang"

func main() {
	a := &blulang.Adapter{}
	if err := a.Init(); err != nil {
		panic(err)
	}
}
