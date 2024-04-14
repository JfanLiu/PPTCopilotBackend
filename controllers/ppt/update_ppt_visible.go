package ppt

import (
	"backend/controllers"
	"backend/models"
)

func (this *Controller) UpdatePptVisible() {

	fileId, err := this.GetInt("file_id")
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	visible, err := this.GetBool("visible")
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	// 更新file
	err = models.UpdateFileVisible(fileId, visible)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "更新记录失败", nil)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", nil)
	this.ServeJSON()
}
