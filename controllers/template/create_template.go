package template

import (
	"backend/controllers"
	"backend/models"
	"encoding/json"
)

type JsonObject map[string]interface{}

type CreateRequest struct {
	Name       string     `json:"name"`
	Cover      JsonObject `json:"cover"`
	Thank      JsonObject `json:"thank"`
	Transition JsonObject `json:"transition"`
	Catalog_3  JsonObject `json:"catalog_3"`
	Catalog_4  JsonObject `json:"catalog_4"`
	Catalog_5  JsonObject `json:"catalog_5"`
	Content_2  JsonObject `json:"content_2"`
	Content_3  JsonObject `json:"content_3"`
	Content_4  JsonObject `json:"content_4"`
}

func getRet(this *Controller, obj JsonObject) string {
	ret, err := json.Marshal(obj)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return ""
	}
	return string(ret)
}
func (this *Controller) CreateTemplate() {
	var req CreateRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&req)

	template := models.Template{
		Name:       req.Name,
		Cover:      getRet(this, req.Cover),
		Thank:      getRet(this, req.Thank),
		Transition: getRet(this, req.Transition),
		Catalog_3:  getRet(this, req.Catalog_3),
		Catalog_4:  getRet(this, req.Catalog_4),
		Catalog_5:  getRet(this, req.Catalog_5),
		Content_2:  getRet(this, req.Content_2),
		Content_3:  getRet(this, req.Content_3),
		Content_4:  getRet(this, req.Content_4),
	}
	err := models.CreateTemplate(template)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "ok", nil)
	this.ServeJSON()
}
