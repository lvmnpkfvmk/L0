package main

//Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

import (
	"fmt"
	"time"
)

func main() {
	dataChannel := make(chan int)
	N := 5

	go func() {
		for i := 1; i <= 10; i++ {
			dataChannel <- i
			time.Sleep(time.Second)
		}
		close(dataChannel)
	}()

	go func() {
		for num := range dataChannel {
			fmt.Println("Received:", num)
		}
	}()

	time.Sleep(time.Duration(N) * time.Second)

	fmt.Println("Time elapsed")
}
