package main

import (
	"L0/cache"
	"L0/http"
	"L0/model"
	"L0/service"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
)

func main() {

	sc, err := stan.Connect("test-cluster", "test2", stan.NatsURL(
		fmt.Sprintf("nats://%s:%s", "localhost", "4222")))

	sub, err := sc.Subscribe("backend", func(m *stan.Msg) {
		var newOrder model.Order
		erro := json.Unmarshal(m.Data, &newOrder)
		if erro != nil {
			log.Printf("Error unmarshalling: %s\n", string(m.Data))
			return
		}
		log.Print(newOrder)
	})
	defer func(sub stan.Subscription) {
		err := sub.Unsubscribe()
		if err != nil {
			log.Printf("Error: %s\n", err)
		}
	}(sub)
	defer func(sub stan.Subscription) {
		err := sub.Close()
		if err != nil {
			log.Printf("Error: %s\n", err)
		}
	}(sub)
	Cache := cache.InitCache()
	service.Cache = Cache

	addr := "postgres://user:user@localhost:12312/db?sslmode=disable"
	db, err := service.InitializeDB(addr)
	if err != nil {
		log.Fatal(err)
		return
	}
	cache.DB = db.Db
	Cache.LoadCache()
	http.NewServer().Run("localhost:8080")
}
