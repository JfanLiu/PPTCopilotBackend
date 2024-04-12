package ppt

import (
	"backend/controllers"
	"backend/models"
	"strconv"
)

func (this *Controller) GetHistory() {
	userId_ := this.Ctx.Input.Param(":user_id")
	userId, err := strconv.Atoi(userId_)

	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	files, err := models.GetHistoryPpt(userId)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "获取历史记录失败", nil)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", files)
	this.ServeJSON()
}
