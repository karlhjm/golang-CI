package service

import (
	"github.com/moandy/canyonsysu/entity"
	"github.com/moandy/canyonsysu/loghelper"
	//"fmt"
	//simplejson "github.com/bitly/go-simplejson"
)

func ListAllCategorys() []string {
	categorys := entity.FindAllCategorys()
	var str []string
	for _, value := range categorys {
		str = append(str, value.Categorys)
	}
	return str
}

func CategoryRegister(name string) (bool, error) {
	var v entity.Categorys
	v.Categorys = name
	if err := entity.CreateCategorys(&v); err != nil {
		loghelper.Error.Println(err)
		return false, err
	} else {
		return true, err
	}
}