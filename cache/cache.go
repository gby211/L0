package cache

import (
	"L0/model"
	"github.com/jmoiron/sqlx"
	"log"
)

var DB *sqlx.DB

type Cache struct {
	orders map[string]model.Order
}

func (c Cache) GetOrderById(id string) model.Order {

	order := c.orders[id]

	return order
}

func (c *Cache) AddOrder(order model.Order) {

	c.orders[order.OrderUID] = order

}

func InitCache() *Cache {
	var cache Cache
	cache.orders = make(map[string]model.Order)
	return &cache
}

func (c *Cache) LoadCache() {
	log.Printf("ggssasdqsaddsaasd - %s", c)
	orderQuerySchema := `SELECT * FROM orders`
	paymentQuerySchema := `SELECT * FROM payments
WHERE order_uid=($1)`
	deliveryQuerySchema := `SELECT * FROM delivery
WHERE order_uid=($1)`
	itemQuerySchema := `SELECT * FROM items
WHERE order_uid=($1)`
	rows, err := DB.Queryx(orderQuerySchema)
	if err != nil {
		log.Printf("orderQuerySchema - %s", err)
		return
	}
	log.Printf("rows0 - %s", rows)
	orders := make(map[string]model.Order)
	for rows.Next() {
		var order model.Order
		log.Printf("rows - %s", &order)
		err := rows.StructScan(&order)
		if err != nil {
			log.Printf("2 - %s", err)
			continue
		}
		row := DB.QueryRowx(paymentQuerySchema, order.OrderUID)
		log.Printf("222 - %s", row)
		err = row.StructScan(&order.Payment)
		if err != nil {
			log.Printf("3 - %s", err)
			continue
		}
		row = DB.QueryRowx(deliveryQuerySchema, order.OrderUID)
		err = row.StructScan(&order.Delivery)
		if err != nil {
			log.Printf("4 - %s", err)
			continue
		}
		rows, err := DB.Queryx(itemQuerySchema, order.OrderUID)
		if err != nil {
			log.Printf("5 - %s", err)
			continue
		}
		var item model.Item
		for rows.Next() {
			err := rows.StructScan(&item)
			if err != nil {
				log.Printf("6 - %s", err)
				continue
			}
			order.Items = append(order.Items, item)
		}
		orders[order.OrderUID] = order
	}
	c.orders = orders
}
