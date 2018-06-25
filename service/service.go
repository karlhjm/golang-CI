package service

import (
	"github.com/moandy/canyonsysu/entity"
	"github.com/moandy/canyonsysu/loghelper"
	"fmt"

	simplejson "github.com/bitly/go-simplejson"
)

//var restaurant_id int

func CustomerRegister(name string, password string, restaurant_id int, phone string) (bool, error) {
	var v entity.Customer
	//v.ID = 0
	v.Name = name
	v.Password = password
	v.Restaurant_ID = restaurant_id
	v.Phone = phone
	if err := entity.CreateCustomer(&v); err != nil {
		loghelper.Error.Println("User Register: Already exist username")
		return false, nil
	}
	return true, nil
}

func ListAllCustomers() []entity.Customer {
	return entity.QueryCustomer(func(u *entity.Customer) bool {
		return true
	})
}

func GetCustomerByName(cname string) *entity.Customer {
	return entity.QueryCustomerByName(cname)
}

func RestaurantRegister(name string, address string, certificates string, servertime string) (bool, error) {
	fmt.Println("2.RestaurantRegisterTest:")
	var v entity.Restaurant
	//v.ID = 0
	v.Name = name
	v.Address = address
	v.Certificates = certificates
	v.Servertime = servertime
	if err := entity.CreateRestaurant(&v); err != nil {
		loghelper.Error.Println("Restaurant Register: Already exist Restaurant")
		return false, nil
	}
	return true, nil
}

func ListAllRestaurants() []entity.Restaurant {
	return entity.QueryRestaurant(func(u *entity.Restaurant) bool {
		return true
	})
}

func GetRestaurantByName(rname string) *entity.Restaurant {
	return entity.QueryRestaurantByName(rname)
}

func UpdateRestaurant(name string, address string, servertime string, certificates string) int {
	filter := func(m *entity.Restaurant) bool {
		return m.Name == name
	}
	return entity.UpdateRestaurant(filter, func(arg1 *entity.Restaurant) {
		arg1.Name = name
		arg1.Servertime = servertime
		arg1.Address = address
		arg1.Certificates = certificates
	})
}

func MenufoodRegister(name string, price float64, restaurant_id int, categorys string, detail string, src string) (bool, error) {
	var v entity.Menufood
	v.Name = name
	v.Price = price
	v.Restaurant_id = restaurant_id
	v.Categorys = categorys
	v.Detail = detail
	v.Src = src
	if err := entity.CreateMenufood(&v); err != nil {
		//fmt.Println("Menufood Register: Already exist Menufood")
		loghelper.Error.Println(err)
		return false, err
	}
	return true, nil
}

func ListAllMenufoods() []entity.Menufood {
	return entity.QueryMenufood(func(u *entity.Menufood) bool {
		return true
	})
}

func ListAllMenufoodsThroughTags() []entity.Menufood_ins {
	menufoods := entity.QueryMenufood(func(u *entity.Menufood) bool {
		return true
	})
	_, str := entity.QueryMenufoodTags()
	var res []entity.Menufood_ins
	for _, tag := range str {
		var a []entity.Menufood
		for _, value := range menufoods {
			if value.Categorys == tag {
				a = append(a, value)
			}
		}
		var b entity.Menufood_ins
		b.Categorys = tag
		b.Menufoods = a
		res = append(res, b)
	}
	return res
}

func GetMenufoodByName(rname string) *entity.Menufood {
	return entity.QueryMenufoodByName(rname)
}

func UpdateMenufood(id int, src string, name string, price float64, detail string, categorys string) int {
	filter := func(m *entity.Menufood) bool {
		return m.ID == id
	}
	return entity.UpdateMenufood(filter, func(arg1 *entity.Menufood) {
		arg1.Name = name
		arg1.Src = src
		arg1.Price = price
		arg1.Detail = detail
		arg1.Categorys = categorys
	})
}

func DeleteMenufood(id int) int {
	return entity.DeleteMenufood(func(m *entity.Menufood) bool {
		return m.ID == id
	})
}

func OrderfoodRegister(customer_phone string, table_id int, order_contain *simplejson.Json, total float64, order_num int, time string) (bool, error) {
	var v entity.Orders
	v.Customer_phone = customer_phone
	v.Table_id = table_id
	v.Restaurant_id = 0
	v.Total = total
	v.Time = time
	id, err := entity.CreateOrder(&v)
	if err != nil {
		loghelper.Error.Println("Order Register: Already exist Order")
		return false, nil
	}
	t := make([]entity.Orderfood, order_num)
	for i := 0; i < order_num; i++ {
		t[i].Order_id = id
		t[i].MenufoodID, _ = order_contain.GetIndex(i).Get("id").Int()
		t[i].Order_num, _ = order_contain.GetIndex(i).Get("order_num").Int()
		err := entity.CreateOrderfood(&t[i])
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
	order := entity.QueryOrder(func(u *entity.Orders) bool {
		return true
	})
	t := make([]entity.Order_ins, len(order))
	for i := 0; i < len(order); i++ {
		t[i].ID = order[i].ID
		t[i].Table_id = order[i].Table_id
		t[i].Total = order[i].Total
		t[i].Time = order[i].Time
		t[i].Customer_phone = order[i].Customer_phone
		t[i].Order_contain = *entity.QueryOrderfoodByID(order[i].ID)
		t[i].Order_num = len(t[i].Order_contain)
	}
	return t
}

func DeleteOrderBy(id int) int {
	del_ordernum := entity.DeleteOrder(func(m *entity.Orders) bool {
		return m.ID == id
	})
	entity.DeleteOrderfood(func(m *entity.Orderfood) bool {
		return m.Order_id == id
	})
	return del_ordernum
}

func GetOrderByID(id int) *entity.Order_ins {
	order := entity.QueryOrderByID(id)
	var t entity.Order_ins
	t.ID = order.ID
	t.Table_id = order.Table_id
	t.Total = order.Total
	t.Time = order.Time
	t.Customer_phone = order.Customer_phone
	t.Order_contain = *entity.QueryOrderfoodByID(order.ID)
	t.Order_num = len(t.Order_contain)
	return &t
}

func GetOrderByPhone(phone string) []entity.Order_ins {
	order := *entity.QueryOrderByPhone(phone)
	t := make([]entity.Order_ins, len(order))
	for i := 0; i < len(order); i++ {
		t[i].ID = order[i].ID
		t[i].Table_id = order[i].Table_id
		t[i].Total = order[i].Total
		t[i].Time = order[i].Time
		t[i].Customer_phone = order[i].Customer_phone
		t[i].Order_contain = *entity.QueryOrderfoodByID(order[i].ID)
		t[i].Order_num = len(t[i].Order_contain)
	}
	return t
}

func CommentRegister(order_id, rating_star int, rate_at,
	username, tags, buddha_src, client_text, merchant_text string) (bool, error) {
	var v entity.Comment
	v.Order_ID = order_id
	v.Rating_star = rating_star
	v.Rate_at = rate_at
	v.Username = username
	v.Merchant_text = merchant_text
	v.Client_text = client_text
	v.Buddha_src = buddha_src
	v.Tags = tags
	err := entity.CreateComment(&v)
	if err != nil {
		loghelper.Error.Println(err)
		return false, err
	}
	return true, nil
}

func ListAllComments() []entity.Comment {
	return entity.QueryComment(func(u *entity.Comment) bool {
		return true
	})
}

func GetCommentCountByTag(tag string) int {
	return entity.QueryCommentCountsByTag(tag)
}

func ListAllTags() []entity.Tags {
	tag_str := entity.QueryTag()	
	if len(tag_str) == 0 {
		return nil
	}
	var tag_map map[string]int
	tag_map = make(map[string]int)
	for _ , v := range tag_str {
		tag_map[v]++
	}
	var tags []entity.Tags
	for k, v := range tag_map {
		tags = append(tags, entity.Tags{Tag:k, Count:v})
	}
	return tags
}

func ListCommentsByCount(begin, offset int) []entity.Comment {
	return entity.QueryCommentByCount(begin, offset)
}

func ListAllCategorys() []string {
	str := entity.FindAllCategorys()
	return str
}