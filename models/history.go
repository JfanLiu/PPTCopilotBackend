package models

import (
	"errors"
	"time"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type History struct {
	Id      int
	User    *User     `orm:"rel(fk)"` // 多对多
	File    *File     `orm:"rel(fk)"` // 多对多
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func UpdateHistory(userId int, fileId int) error {
	// 获取记录
	o := orm.NewOrm()
	history := History{
		User: &User{Id: userId},
		File: &File{Id: fileId},
	}
	err := o.Read(&history, "User", "File")
	if err != nil && !errors.Is(err, orm.ErrNoRows) {
		return err
	}

	if err == nil {
		// 若记录存在，则更新时间
		history.Updated = time.Now()
		_, err := o.Update(&history)
		return err
	} else {
		// 如不存在，则新增记录
		newHistory := History{
			User:    &User{Id: userId},
			File:    &File{Id: fileId},
			Updated: time.Now(),
		}
		_, err = o.Insert(&newHistory)
	}

	return err
}

func GetHistoryPpt(userId int) ([]File, error) {
	o := orm.NewOrm()
	var history []History
	qs := o.QueryTable("history").Filter("User__Id", userId).OrderBy("-Updated")
	_, err := qs.RelatedSel().All(&history)
	if err != nil {
		return nil, err
	}

	var files []File
	for _, h := range history {
		files = append(files, *h.File)
	}

	return files, nil
}
