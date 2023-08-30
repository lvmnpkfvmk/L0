package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"github.com/nats-io/nats.go"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/lvmnpkfvmk/L0/config"
	"github.com/lvmnpkfvmk/L0/internal/model"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()
	cfg := config.Get()

	dbpool, err := pgxpool.New(ctx, cfg.PgURL)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}
	defer dbpool.Close()

	var greeting string
	err = dbpool.QueryRow(ctx, "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}
	log.Println(greeting)


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

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
	return nil
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
  }