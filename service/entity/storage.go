package entity

//CustomerFilter ...
type CustomerFilter func(*Customer) bool  
type RestaurantFilter func(*Restaurant) bool
type MenufoodFilter func(*Menufood) bool
type OrderFilter func(*Orders) bool
type OrderfoodFilter func(*Orderfood) bool
type CommentFilter func(*Comment) bool

func CreateCustomer(v *Customer) error {
	return insertCustomer(v)
	// uData = append(uData, *v.Copy())
	// dirty = true
}

func QueryCustomer(filter CustomerFilter) []Customer {
	var customer []Customer
	cData := findAllCustomers()
	for _, v := range cData {
		if filter(&v) {
			customer = append(customer, v)
		}
	}
	return customer
}

func QueryCustomerByName(v string) *Customer {
	return findCustomerByName(v)
}

//CreateRestaurant ...
func CreateRestaurant(v *Restaurant) error {
	return insertRestaurant(v)
	// uData = append(uData, *v.Copy())
	// dirty = true
}

func QueryRestaurant(filter RestaurantFilter) []Restaurant {
	var restaurant []Restaurant
	rData := findAllRestaurants()
	for _, v := range rData {
		if filter(&v) {
			restaurant = append(restaurant, v)
		}
	}
	return restaurant
}

func QueryRestaurantByName(v string) *Restaurant {
	return findRestaurantByName(v)
}

func UpdateRestaurant(filter RestaurantFilter, switcher func(*Restaurant)) int {
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

func CreateMenufood(v *Menufood) error {
	return insertMenufood(v)
	// uData = append(uData, *v.Copy())
	// dirty = true
}

func QueryMenufood(filter MenufoodFilter) []Menufood {
	var menufood []Menufood
	mfData := findAllMenufoods()
	for _, v := range mfData {
		if filter(&v) {
			menufood = append(menufood, v)
		}
	}
	return menufood
}

func QueryMenufoodByName(v string) *Menufood {
	return findMenufoodByName(v)
}

func UpdateMenufood(filter MenufoodFilter, switcher func(*Menufood)) int {
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

func CreateOrder(v *Orders) (int, error) {
	return insertOrder(v)
	// uData = append(uData, *v.Copy())
	// dirty = true
}

func CreateOrderfood(v *Orderfood) error {
	return insertOrderfood(v)
	// uData = append(uData, *v.Copy())
	// dirty = true
}

func QueryOrder(filter OrderFilter) []Orders {
	var order []Orders
	mfData := findAllOrders()
	for _, v := range mfData {
		if filter(&v) {
			order = append(order, v)
		}
	}
	return order
}

func QueryOrderfood(filter OrderfoodFilter) []Orderfood {
	var orderfood []Orderfood
	mfData := findAllOrderfoods()
	for _, v := range mfData {
		if filter(&v) {
			orderfood = append(orderfood, v)
		}
	}
	return orderfood
}

func QueryOrderfoodByID(id int) *[]Orderfood {
	return findOrderfoodByID(id)
}

func QueryOrderByID(id int) *Orders {
	return findOrderByID(id)
}

func QueryOrderByPhone(phone string) *[]Orders {
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

func CreateComment(v *Comment) error {
	return insertComment(v)
}

func QueryComment(filter CommentFilter) []Comment {
	var comment []Comment
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
	for _ , value := range comment {
		tags = append(tags, value.Tags)
	}
	return tags
}