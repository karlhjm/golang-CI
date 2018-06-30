package entity

type Menufood struct {
	ID            int     `xorm:"pk autoincr" json:"id"`
	Name          string  `xorm:"unique" json:"name"`
	Restaurant_id int     `json:"restaurant_id"`
	Src           string  `json:"src"`
	Price         float64 `json:"price"`
	Detail        string  `json:"detail"`
	Categorys     string  `json:"categorys"`
}

type Menufood_ins struct {
	Categorys string `json:"categorys"`
	Menufoods []Menufood
}

func (u *Menufood) Copy() *Menufood {
	copy := *u
	return &copy
}
