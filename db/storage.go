package db

import "github.com/karlhjm/golang-CI/entity"

//CustomerFilter ...
type CustomerFilter func(*entity.Customer) bool
type RestaurantFilter func(*entity.Restaurant) bool
type MenufoodFilter func(*entity.Menufood) bool
type OrderFilter func(*entity.Orders) bool
type OrderfoodFilter func(*entity.Orderfood) bool
type CommentFilter func(*entity.Comment) bool

func CreateCategorys(v *entity.Categorys) error {
	return insertCategorys(v)
}

func CreateCustomer(v *entity.Customer) error {
	return insertCustomer(v)
}

func QueryCustomer(filter CustomerFilter) []entity.Customer {
	var customer []entity.Customer
	cData := findAllCustomers()
	for _, v := range cData {
		if filter(&v) {
			customer = append(customer, v)
		}
	}
	return customer
}

func QueryCustomerByName(v string) *entity.Customer {
	return findCustomerByName(v)
}

//CreateRestaurant ...
func CreateRestaurant(v *entity.Restaurant) error {
	return insertRestaurant(v)
}

func QueryRestaurant(filter RestaurantFilter) []entity.Restaurant {
	var restaurant []entity.Restaurant
	rData := findAllRestaurants()
	for _, v := range rData {
		if filter(&v) {
			restaurant = append(restaurant, v)
		}
	}
	return restaurant
}

func QueryRestaurantByName(v string) *entity.Restaurant {
	return findRestaurantByName(v)
}

func UpdateRestaurant(filter RestaurantFilter, switcher func(*entity.Restaurant)) int {
	count := 0
	mData := findAllRestaurants()
	for i := 0; i < len(mData); i++ {
		if v := &mData[i]; filter(v) {
			origin := v.Copy()
			switcher(v)
			updateRestaurant(origin, v)
			count++
		}
	}
	return count
}

func CreateMenufood(v *entity.Menufood) error {
	return insertMenufood(v)
}

func QueryMenufood(filter MenufoodFilter) []entity.Menufood {
	var menufood []entity.Menufood
	mfData := findAllMenufoods()
	for _, v := range mfData {
		if filter(&v) {
			menufood = append(menufood, v)
		}
	}
	return menufood
}

func QueryMenufoodByName(v string) *entity.Menufood {
	return findMenufoodByName(v)
}

func UpdateMenufood(filter MenufoodFilter, switcher func(*entity.Menufood)) int {
	count := 0
	mData := findAllMenufoods()
	for i := 0; i < len(mData); i++ {
		if v := &mData[i]; filter(v) {
			origin := v.Copy()
			switcher(v)
			updateMenufood(origin, v)
			count++
		}
	}
	return count
}

// DeleteUser : delete users
// @param a lambda function as the filter
// @return the number of deleted users
func DeleteMenufood(filter MenufoodFilter) int {
	count := 0
	uData := findAllMenufoods()
	length := len(uData)
	for i := 0; i < length; {
		if filter(&uData[i]) {
			length--
			deleteMenufood(&uData[i])
			uData[i] = uData[length]
			uData = uData[:length]
			count++
		} else {
			i++
		}
	}
	return count
}

func CreateOrder(v *entity.Orders) (int, error) {
	return insertOrder(v)
}

func CreateOrderfood(v *entity.Orderfood) error {
	return insertOrderfood(v)
}

func QueryOrder(filter OrderFilter) []entity.Orders {
	var order []entity.Orders
	mfData := findAllOrders()
	for _, v := range mfData {
		if filter(&v) {
			order = append(order, v)
		}
	}
	return order
}

func QueryOrderfood(filter OrderfoodFilter) []entity.Orderfood {
	var orderfood []entity.Orderfood
	mfData := findAllOrderfoods()
	for _, v := range mfData {
		if filter(&v) {
			orderfood = append(orderfood, v)
		}
	}
	return orderfood
}

func QueryOrderfoodByID(id int) *[]entity.Orderfood {
	return findOrderfoodByID(id)
}

func QueryOrderByID(id int) *entity.Orders {
	return findOrderByID(id)
}

func QueryOrderByPhone(phone string) *[]entity.Orders {
	return findOrderByPhone(phone)
}

func DeleteOrder(filter OrderFilter) int {
	count := 0
	uData := findAllOrders()
	length := len(uData)
	for i := 0; i < length; {
		if filter(&uData[i]) {
			length--
			deleteOrder(&uData[i])
			uData[i] = uData[length]
			uData = uData[:length]
			count++
		} else {
			i++
		}
	}
	return count
}

func DeleteOrderfood(filter OrderfoodFilter) int {
	count := 0
	uData := findAllOrderfoods()
	length := len(uData)
	for i := 0; i < length; {
		if filter(&uData[i]) {
			length--
			deleteOrderfood(&uData[i])
			uData[i] = uData[length]
			uData = uData[:length]
			count++
		} else {
			i++
		}
	}
	return count
}

func CreateComment(v *entity.Comment) error {
	return insertComment(v)
}

func QueryComment(filter CommentFilter) []entity.Comment {
	var comment []entity.Comment
	mfData := findAllComments()
	for _, v := range mfData {
		if filter(&v) {
			comment = append(comment, v)
		}
	}
	return comment
}

func QueryCommentCountsByTag(tag string) int {
	return findCommentCountsByTag(tag)
}

func QueryTag() []string {
	comment := findAllComments()
	var tags []string
	for _, value := range comment {
		tags = append(tags, value.Tags)
	}
	return tags
}

func QueryMenufoodTags() (int, []string) {
	return findAllMenufoodByTag()
}

func QueryCommentByCount(begin, offset int) []entity.Comment {
	return findCommentByCount(begin, offset)
}
