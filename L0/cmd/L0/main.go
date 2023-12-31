package main

import (
	"context"
	"encoding/json"
	"log"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nats-io/nats.go"

	"github.com/lvmnpkfvmk/L0/config"
	"github.com/lvmnpkfvmk/L0/internal/model"
	"github.com/lvmnpkfvmk/L0/internal/repositories/orderrepo"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()
	cfg := config.Get()

	repo, err := orderrepo.NewOrderRepository(ctx, cfg)
	if err != nil {
		log.Fatalf("Error creating repository: %v", err)
	}

	conn, err := nats.Connect("nats://nats-streaming:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Subscribe("orders",
		func(msg *nats.Msg) {
			var order model.Order
			err := json.Unmarshal(msg.Data, &order)
			if err != nil {
				log.Printf("Error unmarshaling order: %v\n", err)
				return
			}
			log.Println(order)

			err = repo.CreateOrder(&order)
			if err != nil {
				log.Printf("Error unmarshaling order: %v\n", err)
				return
			}

			log.Printf("Received order: %v\n", order.OrderUID)
		},
	)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.Static("/", "../../web")
	// Routes
	e.GET("/get_order/:id", func(c echo.Context) error {
		order, err := repo.GetOrder(c.Param("id"))
		if err != nil {
			log.Printf("Error not found: %v\n", err)
			return c.NoContent(http.StatusNotFound)
		}

		log.Printf("Sending order id: %s\n", order.OrderUID)
		return c.JSON(http.StatusFound, order)
	})

	// Start server
	e.Logger.Fatal(e.Start(cfg.HTTPAddr))
	return nil
}
