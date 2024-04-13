package user

import (
	"backend/controllers"
	"backend/models"
	"strconv"
)

func (this *Controller) GetPpts() {
	// 获取路由参数
	id_ := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(id_)
	if err != nil {

		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}
	projects, err := models.GetProjects(id)
	if err != nil {

		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	var files []models.File
	for _, project := range projects {

		tempFiles, err := models.GetAllFilesOfProj(project.Id)

		if err != nil {
			this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
			this.ServeJSON()
			return
		}
		files = append(files, tempFiles...)
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", files)
	this.ServeJSON()

}
