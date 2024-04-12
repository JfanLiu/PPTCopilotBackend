package ppt

import (
	"backend/controllers"
	"backend/models"
)

func (this *Controller) GetAllPublicPpt() {

	files, err := models.GetAllPublicPpt()
	if err != nil {

		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	files = models.RefactFiles(files)

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", files)
	this.ServeJSON()
}
