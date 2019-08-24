package db

import (
	_ "github.com/go-sql-driver/mysql" // for init
	"github.com/go-xorm/xorm"
	"github.com/karlhjm/golang-CI/entity"
)

var engine *xorm.Engine

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	db, err := xorm.NewEngine("mysql", "root:mysql@tcp(localhost:3306)/hello2?charset=utf8&parseTime=true")
	checkErr(err)
	engine = db
	db.Sync2(new(entity.Restaurant), new(entity.Menufood), new(entity.Orders), new(entity.Orderfood), new(entity.Comment), new(entity.Categorys))
}
