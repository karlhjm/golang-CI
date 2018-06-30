package db

import (
	"github.com/moandy/canyonsysu/loghelper"
	"github.com/moandy/canyonsysu/entity"
	_ "github.com/go-sql-driver/mysql" // for init
	//"github.com/go-xorm/xorm"
)


func insertCategorys(v *entity.Categorys) error {
	if affected, err := engine.Insert(v); err != nil {
		loghelper.Error.Println("insertCategorys Error:", affected, err)
		return err
	}
	return nil
}

func FindAllCategorys() []entity.Categorys {
	sql := "select * from categorys"
	vec := make([]entity.Categorys, 0)
	err := engine.Sql(sql).Find(&vec)
	if err != nil {
		loghelper.Error.Println(err)
	}
	return vec
}