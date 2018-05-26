package entity

type Orderfood struct {
	ID         int `xorm:"pk autoincr" json:"id"`
	MenufoodID int `xorm:"notnull" json:"menufood_id"`
	Order_id   int `xorm:"notnull" json:"order_id"`
	Order_num  int `xorm:"notnull default(0)" json:"order_num"`
}
