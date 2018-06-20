package entity

type Customer struct {
	//ID            int `xorm:"pk autoincr"`
	Restaurant_ID int
	Name          string `xorm:"pk"`
	Phone         string
	Password      string `xorm:"notnull"`
	//CreateAt      *time.Time
	//order_ID            int
	//order_Restaurant_ID int
}

func (m_User Customer) init(t_Restaurant_id int, t_Name, t_Password, t_Phone string) {
	m_User.Name = t_Name
	m_User.Password = t_Password
	//m_User.Email= t_Email
	m_User.Phone = t_Phone
	m_User.Restaurant_ID = t_Restaurant_id
}

// GetID
// func GetID(a Customer) int {
// 	return a.ID
// }

func GetName(a Customer) string {
	return a.Name
}

func GetPhone(a Customer) string {
	return a.Phone
}

func GetPassword(a Customer) string {
	return a.Password
}

func GetRestaurant_ID(a Customer) int {
	return a.Restaurant_ID
}
