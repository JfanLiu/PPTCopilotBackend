package ppt

import (
	"backend/controllers"
	"backend/models"
)

func (this *Controller) UpdateHistory() {

	token := this.Ctx.Request.Header.Get("token")
	err := models.CheckToken(token)

	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "未登录", nil)
		this.ServeJSON()
		return
	}
	userId := models.GetUserId(token)

	projectId, err := this.GetInt("project_id")
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	filename := this.GetString("filename")
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	file, err := models.GetFileOfProj(filename, projectId)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "获取文件失败", nil)
		this.ServeJSON()
		return
	}

	err = models.UpdateHistory(userId, file.Id)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "更新记录失败", nil)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", nil)
	this.ServeJSON()
}
