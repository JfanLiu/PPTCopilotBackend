package ppt

import (
	"backend/controllers"
	"backend/models"
)

func (this *Controller) GetHistory() {

	token := this.Ctx.Request.Header.Get("token")
	err := models.CheckToken(token)

	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "未登录", nil)
		this.ServeJSON()
		return
	}
	userId := models.GetUserId(token)

	files, err := models.GetHistoryPpt(userId)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "获取历史记录失败", nil)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", files)
	this.ServeJSON()
}
