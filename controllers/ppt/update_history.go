package ppt

import (
	"backend/controllers"
	"backend/models"
)

func (this *Controller) UpdateHistory() {

	token, err := this.Ctx.Request.Cookie("token")
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "未登录", nil)
		this.ServeJSON()
		return
	}
	userId := models.GetUserId(token.Value)

	fileId, err := this.GetInt("file_id")
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	err = models.UpdateHistory(userId, fileId)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "更新记录失败", nil)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", nil)
	this.ServeJSON()
}
