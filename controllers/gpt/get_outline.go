package gpt

import (
	"backend/controllers"
	"backend/models"
	"strconv"
)

// GetOutline 根据大纲id获取大纲信息
func (this *Controller) GetOutline() {
	id_ := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(id_)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	// 调用 models 包中的 GetOutline 方法获取指定 ID 的大纲信息
	outline, err := models.GetOutline(id)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	// 调用 models 包中的 RefactOutline 调整大纲格式
	outline = models.RefactOutline(outline)

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", outline)
	this.ServeJSON()

}
