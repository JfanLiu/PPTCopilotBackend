package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Image struct {
	Id    int    `orm:"pk"`
	Image string `orm:"type(text)"`
}

func GetBase64ById(id int) (string, error) {
	o := orm.NewOrm()
	var image Image
	err := o.QueryTable("image").Filter("id", id).One(&image)
	return image.Image, err
}
