package main

import (
	"L0/model"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"time"
)

func main() {
	sc, err := stan.Connect("test-cluster", "test")
	if err != nil {
		log.Fatalf("Failed to connect to NATS Streaming: %v", err)
	}

	defer func(sc stan.Conn) {
		err := sc.Close()
		if err != nil {
			log.Printf("Error: %s\n", err)
		}
	}(sc)

	for i := 3; i < 5; i++ {
		order := GetExampleOrder(fmt.Sprintf("test-%d", i))
		bytes, err := json.Marshal(order)
		if err != nil {
			if err != nil {
				log.Printf("Error: %s\n", err)
			}
		}
		if err = sc.Publish("backend", bytes); err != nil {
			log.Fatalf("Failed to publish message: %v", err)
		}
		log.Printf("Sent message: %s", bytes)
		time.Sleep(1 * time.Second)
	}
}

func GetExampleOrder(orderUID string) model.Order {
	return model.Order{
		OrderUID:    orderUID,
		TrackNumber: "WBILMTESTTRACK",
		Entry:       "WBIL",
		Delivery: model.Delivery{
			Name:    "Test Testov",
			Phone:   "+9720000000",
			Zip:     "2639809",
			City:    "Kiryat Mozkin",
			Address: "Ploshad Mira 15",
			Region:  "Kraiot",
			Email:   "test@gmail.com",
		},
		Payment: model.Payment{
			Transaction:  "b563feb7b2b84b6test",
			RequestID:    "",
			Currency:     "USD",
			Provider:     "wbpay",
			Amount:       1817,
			PaymentDt:    1637907727,
			Bank:         "alpha",
			DeliveryCost: 1500,
			GoodsTotal:   317,
			CustomFee:    0,
		},
		Items: []model.Item{
			{
				ChrtID:      9934930,
				TrackNumber: "WBILMTESTTRACK",
				Price:       453,
				Rid:         "ab4219087a764ae0btest",
				Name:        "Mascaras",
				Sale:        30,
				Size:        "0",
				TotalPrice:  317,
				NmID:        2389212,
				Brand:       "Vivienne Sabo",
				Status:      202,
			},
			{
				ChrtID:      9934931,
				TrackNumber: "WBILMTESTTRACK",
				Price:       555,
				Rid:         "dddd",
				Name:        "dsa",
				Sale:        20,
				Size:        "1",
				TotalPrice:  520,
				NmID:        2389212,
				Brand:       "Vivienne Sabo",
				Status:      202,
			},
		},
		Locale:            "en",
		InternalSignature: "",
		CustomerID:        "test",
		DeliveryService:   "meest",
		Shardkey:          "9",
		SmID:              99,
		DateCreated:       time.Now().String(),
		OofShard:          "1",
	}

}
