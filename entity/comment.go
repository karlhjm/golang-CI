package entity

type Comment struct {
	ID            int    `xorm:"pk autoincr unique" json:"id"`
	Order_ID      int    `json:"order_id"`
	Username      string `json:"usr_name"`
	Buddha_src    string `json:"usr_photo"`
	Rate_at       string `json:"comment_at"`
	Rating_star   int    `json:"comment_star"`
	Tags          string `json:"tag"`
	Client_text   string `json:"client_text"`
	Merchant_text string `json:"merchant_text"`
}

type Tags struct {
	Tag   string `json:"tag"`
	Count int    `json:"count"`
}
