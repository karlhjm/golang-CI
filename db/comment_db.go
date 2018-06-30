package db

import (
	_ "github.com/go-sql-driver/mysql" // for init
	"github.com/karl-jm-huang/golang-CI/entity"
	"github.com/karl-jm-huang/golang-CI/loghelper"
	//"github.com/go-xorm/xorm"
)

func insertComment(v *entity.Comment) error {
	//fmt.Println("3.RestaurantRegisterTest:")
	if affected, err := engine.Insert(v); err != nil {
		loghelper.Error.Println("insertComment Error:", affected, err)
		return err
	}
	return nil
}

func findAllComments() []entity.Comment {
	vec := make([]entity.Comment, 0)
	if err := engine.Find(&vec); err != nil {
		loghelper.Error.Println("findAllComments Error:", err)
	}
	return vec
}

func findCommentCountsByTag(tag string) int {
	comment := new(entity.Comment)
	total, err := engine.Where("tags =?", tag).Count(comment)
	if err != nil {
		loghelper.Error.Println(err)
	}
	return int(total)
}

func findAllMenufoodByTag() (int, []string) {
	menufood := new(entity.Menufood)
	tagnum, err := engine.Where("").Count(menufood)
	if err != nil {
		loghelper.Error.Println(err)
	}
	vec := make([]entity.Menufood, 0)
	if err := engine.Find(&vec); err != nil {
		loghelper.Error.Println(err)
	}
	var str []string
	for _, value := range vec {
		str = append(str, value.Categorys)
	}
	return int(tagnum), str
}

func findCommentByCount(begin int, offset int) []entity.Comment {
	vec := make([]entity.Comment, 0)
	sql := "select * from comment order by i_d desc limit ?, ?"
	err := engine.Sql(sql, begin, offset).Find(&vec)
	if err != nil {
		loghelper.Error.Println(err)
	}
	return vec
}
