package model

type Order struct {
	OrderUID          string   `db:"order_uid" json:"order_uid" `
	TrackNumber       string   `db:"track_number" json:"track_number"`
	Entry             string   `db:"entry" json:"entry"`
	Delivery          Delivery `json:"delivery"`
	Payment           Payment  `json:"payment"`
	Items             []Item   `json:"items"`
	Locale            string   `db:"locale" json:"locale"`
	InternalSignature string   `db:"internal_signature" json:"internal_signature"`
	CustomerID        string   `db:"customer_id" json:"customer_id"`
	DeliveryService   string   `db:"delivery_service" json:"delivery_service"`
	Shardkey          string   `db:"shardkey" json:"shardkey"`
	SmID              int      `db:"sm_id" json:"sm_id"`
	DateCreated       string   `db:"date_created" json:"date_created"`
	OofShard          string   `db:"oof_shard" json:"oof_shard"`
}

type Delivery struct {
	OrderUid string `db:"order_uid" json:",omitempty"`
	Name     string `db:"name" json:"name"`
	Phone    string `db:"phone" json:"phone"`
	Zip      string `db:"zip" json:"zip"`
	City     string `db:"city" json:"city"`
	Address  string `db:"address" json:"address"`
	Region   string `db:"region" json:"region"`
	Email    string `db:"email" json:"email"`
}

type Payment struct {
	OrderUid     string `db:"order_uid" json:",omitempty"`
	Transaction  string `db:"transaction" json:"transaction"`
	RequestID    string `db:"request_id" json:"request_id"`
	Currency     string `db:"currency" json:"currency"`
	Provider     string `db:"provider" json:"provider"`
	Amount       int    `db:"amount" json:"amount"`
	PaymentDt    int    `db:"payment_dt" json:"payment_dt"`
	Bank         string `db:"bank" json:"bank"`
	DeliveryCost int    `db:"delivery_cost" json:"delivery_cost"`
	GoodsTotal   int    `db:"goods_total" json:"goods_total"`
	CustomFee    int    `db:"custom_fee" json:"custom_fee"`
}

type Item struct {
	OrderUid    string `db:"order_uid" json:",omitempty"`
	ChrtID      int    `db:"chrt_id" json:"chrt_id"`
	TrackNumber string `db:"track_number" json:"track_number"`
	Price       int    `db:"price" json:"price"`
	Rid         string `db:"rid" json:"rid"`
	Name        string `db:"name" json:"name"`
	Sale        int    `db:"sale" json:"sale"`
	Size        string `db:"size" json:"size"`
	TotalPrice  int    `db:"total_price" json:"total_price"`
	NmID        int    `db:"nm_id" json:"nm_id"`
	Brand       string `db:"brand" json:"brand"`
	Status      int    `db:"status" json:"status"`
}
