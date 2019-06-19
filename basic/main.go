package main

import (
	"fmt"

	"github.com/zeromq/goczmq"
)

func main() {
	fmt.Println("go")
	goczmq.NewPub("endpoint")
}
