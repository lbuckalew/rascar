package main

import "github.com/lbuckalew/blulang"
import "fmt"

func main() {

	fmt.Println("t1")
	a := &blulang.Adapter{}
	if err := a.Init(); err != nil {
		panic(err)
	}

	a.Enable()
}
