package ppt

import (
	"backend/controllers"
	"backend/models"
	"strconv"
)

func (this *Controller) UpdateHistory() {

	userId_ := this.Ctx.Input.Param(":user_id")
	userId, err := strconv.Atoi(userId_)

	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	fileId, err := this.GetInt("file_id")

	err = models.UpdateHistory(userId, fileId)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "更新记录失败", nil)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", nil)
	this.ServeJSON()
}
