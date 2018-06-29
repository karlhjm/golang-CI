package entity

type Categorys struct {
	ID            int     `xorm:"pk autoincr" json:"id"`
	Categorys     string  `json:"categorys"`
}