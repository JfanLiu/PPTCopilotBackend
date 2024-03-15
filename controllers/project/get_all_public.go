package project

import (
	"backend/controllers"
	"backend/models"
)

func (this *Controller) GetAllPublic() {

	projects := models.GetAllPublicProjects()

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", projects)
	this.ServeJSON()
}
