package ppt

import (
	"backend/controllers"
	"backend/models"
	"fmt"
	"strings"
)

func (this *Controller) SearchUserPpt() {
	// 获取用户
	token := this.Ctx.Request.Header.Get("token")
	err := models.CheckToken(token)

	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "未登录", nil)
		this.ServeJSON()
		return
	}
	userId := models.GetUserId(token)

	filterWords := this.GetString("filter_words")

	// 拆分关键词
	keywords := strings.Split(filterWords, " ")

	// 获取用户的项目
	project, err := models.GetDefaultProjectByUser(userId)
	if err != nil {

		this.Data["json"] = controllers.MakeResponse(controllers.Err, "获取用户项目失败", nil)
		this.ServeJSON()
		return
	}

	// 查询项目
	files, err := models.SearchPptOfProj(project.Id, keywords)

	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "获取ppt列表失败", nil)
		fmt.Println(err.Error())
		this.ServeJSON()
		return
	}

	files = models.RefactFiles(files)

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "获取ppt列表成功", files)
	this.ServeJSON()

}
