package entity

type Orders struct {
	ID             int    `xorm:"pk autoincr"`
	Customer_phone string `xorm:"pk notnull"`
	Table_id       int    `xorm:"notnull"`
	Restaurant_id  int
	Total          float64
	Time           string
}

//"0000-00-00/00:00"

type Order_ins struct {
	ID             int         `json:"order_id"`
	Table_id       int         `json:"table_id"`
	Order_num      int         `json:"order_num"`
	Order_contain  []Orderfood `json:"order_contain"`
	Total          float64     `json:"total"`
	Customer_phone string      `json:"customer_phone"`
	Time           string      `json:"order_time"`
}

type Orderfood_ins struct {
	MenufoodID int `json:"id"`
	Order_num  int `json:"order_num"`
}
