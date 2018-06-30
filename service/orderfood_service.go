package service

import (
	"github.com/moandy/canyonsysu/entity"
	"github.com/moandy/canyonsysu/loghelper"
	//"fmt"
	simplejson "github.com/bitly/go-simplejson"
	"github.com/moandy/canyonsysu/db"
)

func OrderfoodRegister(customer_phone string, table_id int, order_contain *simplejson.Json, total float64, order_num int, time string) (bool, error) {
	var v entity.Orders
	v.Customer_phone = customer_phone
	v.Table_id = table_id
	v.Restaurant_id = 0
	v.Total = total
	v.Time = time
	id, err := db.CreateOrder(&v)
	if err != nil {
		loghelper.Error.Println("Order Register: Already exist Order")
		return false, nil
	}
	t := make([]entity.Orderfood, order_num)
	for i := 0; i < order_num; i++ {
		t[i].Order_id = id
		t[i].MenufoodID, _ = order_contain.GetIndex(i).Get("id").Int()
		t[i].Order_num, _ = order_contain.GetIndex(i).Get("order_num").Int()
		err := db.CreateOrderfood(&t[i])
		if err != nil {
			DeleteOrderBy(id) //delete the order
			loghelper.Info.Println("Orderfood Register:Fail!")
			loghelper.Error.Println("Orderfood Register:Fail!")
			return false, nil
		}
	}
	return true, nil
}

func ListAllOrders() []entity.Order_ins {
	order := db.QueryOrder(func(u *entity.Orders) bool {
		return true
	})
	t := make([]entity.Order_ins, len(order))
	for i := 0; i < len(order); i++ {
		t[i].ID = order[i].ID
		t[i].Table_id = order[i].Table_id
		t[i].Total = order[i].Total
		t[i].Time = order[i].Time
		t[i].Customer_phone = order[i].Customer_phone
		t[i].Order_contain = *db.QueryOrderfoodByID(order[i].ID)
		t[i].Order_num = len(t[i].Order_contain)
	}
	return t
}

func DeleteOrderBy(id int) int {
	del_ordernum := db.DeleteOrder(func(m *entity.Orders) bool {
		return m.ID == id
	})
	db.DeleteOrderfood(func(m *entity.Orderfood) bool {
		return m.Order_id == id
	})
	return del_ordernum
}

func GetOrderByID(id int) *entity.Order_ins {
	order := db.QueryOrderByID(id)
	var t entity.Order_ins
	t.ID = order.ID
	t.Table_id = order.Table_id
	t.Total = order.Total
	t.Time = order.Time
	t.Customer_phone = order.Customer_phone
	t.Order_contain = *db.QueryOrderfoodByID(order.ID)
	t.Order_num = len(t.Order_contain)
	return &t
}

func GetOrderByPhone(phone string) []entity.Order_ins {
	order := *db.QueryOrderByPhone(phone)
	t := make([]entity.Order_ins, len(order))
	for i := 0; i < len(order); i++ {
		t[i].ID = order[i].ID
		t[i].Table_id = order[i].Table_id
		t[i].Total = order[i].Total
		t[i].Time = order[i].Time
		t[i].Customer_phone = order[i].Customer_phone
		t[i].Order_contain = *db.QueryOrderfoodByID(order[i].ID)
		t[i].Order_num = len(t[i].Order_contain)
	}
	return t
}
