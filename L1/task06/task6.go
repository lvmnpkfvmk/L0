package main

//Реализовать все возможные способы остановки выполнения горутины.

import (
	"context"
	"fmt"
	"time"
)

func cancelContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Cancelling")
			return
		default:
			fmt.Println("Working")
			time.Sleep(time.Second)
		}
	}
}

func cancelChannel(c chan struct{}) {
	for {
		select {
		case <-c:
			fmt.Println("Cancelling")
			return
		default:
			fmt.Println("Working")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go cancelContext(ctx)
	time.Sleep(time.Second * 2)
	cancel()

	fmt.Println("cancelContext cancelled")

	c := make(chan struct{})
	go cancelChannel(c)
	time.Sleep(time.Second * 2)
	close(c)

	fmt.Println("cancelChannel cancelled")

}
