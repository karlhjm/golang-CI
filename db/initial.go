package db

import (
	"github.com/moandy/canyonsysu/entity"
	_ "github.com/go-sql-driver/mysql" // for init
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	db, err := xorm.NewEngine("mysql", "root:wsm971058171@tcp(119.23.243.149:3306)/canyonsysu?charset=utf8&parseTime=true")
	checkErr(err)
	engine = db
	db.Sync2(new(entity.Restaurant), new(entity.Menufood), new(entity.Orders), new(entity.Orderfood), new(entity.Comment), new(entity.Categorys))
}