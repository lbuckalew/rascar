package main

import "github.com/lbuckalew/blulang"
import "fmt"

func main() {

	fmt.Println("t1")
	a := &blulang.Adapter{}
	fmt.Println("t2")
	if err := a.Init(); err != nil {
		panic(err)
	}
	fmt.Println("t3")
}
