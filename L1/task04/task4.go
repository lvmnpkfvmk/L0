package main

// Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.

import (
    "context"
    "fmt"
    "os"
    "os/signal"
    "sync"
    "syscall"
    "time"
)

func producer(ctx context.Context, dataChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
    defer close(dataChan)

    for {
        select {
        case <-ctx.Done():
            fmt.Println("Producer done")
            return
        default:
            data := "data"
            dataChan <- data
            time.Sleep(time.Second)
        }
    }
}

func worker(ctx context.Context, id int, dataChan <-chan string, wg *sync.WaitGroup) {
    defer wg.Done()

    for {
        select {
        case <-ctx.Done():
            fmt.Println("Worker done")
            return
        case data := <-dataChan:
            fmt.Println(id, "Received: ", data)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    go func() {
        c := make(chan os.Signal, 1)
        signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
        <-c
        cancel()
    }()
	var n int
	fmt.Scan(&n)

    dataChan := make(chan string, n)

    var wg sync.WaitGroup
    for i := 0; i < n; i++ {
        wg.Add(1)
        go worker(ctx, i, dataChan, &wg)
    }

    wg.Add(1)
    go producer(ctx, dataChan, &wg)

    wg.Wait()
}