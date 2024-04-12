package ppt

import (
	"backend/controllers"
	"backend/models"
	"strings"
)

func (this *Controller) SearchAllPublicPpt() {
	filterWords := this.GetString("filter_words")

	// 拆分关键词
	keywords := strings.Split(filterWords, " ")

	// 查询项目
	projects, err := models.SearchAllPublicPpt(keywords)

	if err != nil {

		this.Data["json"] = controllers.MakeResponse(controllers.Err, "获取ppt列表失败", nil)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "获取ppt列表成功", projects)
	this.ServeJSON()

}
