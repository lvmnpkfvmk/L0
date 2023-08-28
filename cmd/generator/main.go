package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/lvmnpkfvmk/L0/internal/model"
)

const (
	testsNumber  = 10
	sendingDelay = 5 * time.Second
)

func main() {
	conn, err := nats.Connect("nats://nats-streaming:4222")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for i := 0; i < testsNumber; i++ {
		time.Sleep(sendingDelay)
		data := `{"order_uid":"b563feb7b2b84b6test","track_number":"WBILMTESTTRACK","entry":"WBIL","delivery":{"name":"Test Testov","phone":"+9720000000","zip":"2639809","city":"Kiryat Mozkin","address":"Ploshad Mira 15","region":"Kraiot","email":"test@gmail.com"},"payment":{"transaction":"b563feb7b2b84b6test","request_id":"","currency":"USD","provider":"wbpay","amount":1817,"payment_dt":1637907727,"bank":"alpha","delivery_cost":1500,"goods_total":317,"custom_fee":0},"items":[{"chrt_id":9934930,"track_number":"WBILMTESTTRACK","price":453,"rid":"ab4219087a764ae0btest","name":"Mascaras","sale":30,"size":"0","total_price":317,"nm_id":2389212,"brand":"Vivienne Sabo","status":202}],"locale":"en","internal_signature":"","customer_id":"test","delivery_service":"meest","shardkey":"9","sm_id":99,"date_created":"2021-11-26T06:22:19Z","oof_shard":"1"}`

		var order model.Order
		err := json.Unmarshal([]byte(data), &order)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		bytes, err := json.Marshal(order)
		if err != nil {
			log.Println("ERROR: json.Marshal:", err)
			continue
		}

		err = conn.Publish("orders", bytes)
		if err != nil {
			log.Println("ERROR: conn.Publish:", err)
			continue
		}

		log.Println("Published order id:", i)
	}
	// send trash data
	err = conn.Publish("orders", []byte("lolkek"))
	if err != nil {
		log.Println("ERROR: conn.Publish:", err)
	}
}