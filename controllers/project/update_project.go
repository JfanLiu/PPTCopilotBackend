package project

import (
	"backend/controllers"
	"backend/models"
	"encoding/json"
	"strconv"
)

type UpdateProjectRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

func (this *Controller) UpdateProject() {
	var request UpdateProjectRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&request)

	id_ := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(id_)
	if err != nil {

		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}
	if request.Name != nil {
		_, err := models.UpdateProjectName(id, *request.Name)
		if err != nil {

			this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
			this.ServeJSON()
			return
		}
	}
	if request.Description != nil {
		_, err := models.UpdateProjectDescription(id, *request.Description)
		if err != nil {

			this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
			this.ServeJSON()
			return
		}
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", nil)
	this.ServeJSON()
}
