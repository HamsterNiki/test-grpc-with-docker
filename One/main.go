package main

import (
	"fmt"
	pb "github.com/HamsterNiki/test-grpc-with-docker/pr"
	"google.golang.org/grpc"
	"time"
)

func main() {
	for {
		fmt.Println("One working")
		time.Sleep(2 * time.Second)
	}
}
