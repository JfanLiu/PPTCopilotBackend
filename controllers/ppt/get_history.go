package ppt

import (
	"backend/controllers"
	"backend/models"
)

func (this *Controller) GetHistory() {

	token, err := this.Ctx.Request.Cookie("token")
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "未登录", nil)
		this.ServeJSON()
		return
	}
	userId := models.GetUserId(token.Value)

	files, err := models.GetHistoryPpt(userId)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "获取历史记录失败", nil)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", files)
	this.ServeJSON()
}
