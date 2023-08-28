package main

import (
	// "context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"

	// "github.com/lvmnpkfvmk/L0/config"
	"github.com/lvmnpkfvmk/L0/internal/model"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// ctx := context.Background()
	// cfg := config.Get()

	conn, err := nats.Connect("nats://nats-streaming:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Subscribe("orders",
		func(msg *nats.Msg) {
			var order model.Order
			funcErr := json.Unmarshal(msg.Data, &order)
			if funcErr != nil {
				log.Println(funcErr)
				return
			}
			log.Println(order)

			// funcErr = stor.CreateOrder(order)
			// if funcErr != nil {
			// 	logger.Error(errors.Wrap(funcErr, "conn.Subscribe: stor.CreateOrder"))
			// 	return
			// }

			// logger.Debugf("Received order %v", order)
		},
	)
	r := gin.Default()
	r.GET("/orders/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	return nil
}
