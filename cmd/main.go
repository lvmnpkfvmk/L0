package main

import (
	// "context"
	"fmt"
	"log"

	"github.com/lvmnpkfvmk/L0/config"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// ctx := context.Background()
	cfg := config.Get()
	fmt.Println(cfg)

	return nil
}