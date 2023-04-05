package service

import (
	"L0/cache"
	"L0/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type PostgresDB struct {
	Db *sqlx.DB
}

var Cache *cache.Cache

func InitializeDB(uri string) (*PostgresDB, error) {
	db, err := sqlx.Open("postgres", uri)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return &PostgresDB{db}, nil
}

func (pgDb *PostgresDB) CreateOrder(order model.Order) error {

	insertOrderSchema := `INSERT INTO orders 
	VALUES (:order_uid, :track_number, :entry,
:locale, :internal_signature, :customer_id, :delivery_service,:shardkey,
:sm_id, :date_created, :oof_shard)`
	insertDeliverySchema := `INSERT INTO delivery 
	VALUES (:order_uid, :name, :phone, :zip, :city, :address, :region, :email)
`
	insertPaymentSchema := `INSERT INTO payments 
	VALUES (:order_uid, :request_id, :currency, :provider, :amount,
:payment_dt, :bank, :delivery_cost, :goods_total, :custom_fee)
`
	insertItemSchema := `INSERT INTO items 
	VALUES (:order_uid, :chrt_id, :track_number, :price, :rid,
:name, :sale, :size, :total_price, :nm_id, :brand, :status)
`
	_, err := pgDb.Db.Exec(insertOrderSchema, order)
	if err != nil {
		return err
	}

	_, err = pgDb.Db.Exec(insertDeliverySchema, order.Delivery)
	if err != nil {
		return err
	}
	_, err = pgDb.Db.Exec(insertPaymentSchema, order.Payment)
	if err != nil {
		return err
	}
	_, err = pgDb.Db.Exec(insertItemSchema, order.Items)
	if err != nil {
		return err
	}

	Cache.AddOrder(order)

	return nil

}

func GetById(orderId string) model.Order {

	return Cache.GetOrderById(orderId)
}
