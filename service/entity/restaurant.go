package entity

type Restaurant struct {
	ID           int    `xorm:"pk autoincr" json:"id"`
	Name         string `xorm:"pk notnull unique" json:"name"`
	Address      string `json:"address"`
	Certificates string `json:"certificates"`
	Servertime   string `json:"server_time"`
	//password string
	//CreateAt *time.Time
	//order_ID            int
	//order_Restaurant_ID int
}

func (u *Restaurant) Copy() *Restaurant {
	copy := *u
	return &copy
}

func Getres_ID(a Restaurant) int {
	return a.ID
}

func GetRestaurant_Name(a Restaurant) string {
	return a.Name
}

func GetRestaurant_Address(a Restaurant) string {
	return a.Address
}

// func GetRestaurant_CreateAt(a customer) *time.Time {
// 	return a.CreateAt
// }
