package agent

import (
	"backend/conf"
	"backend/controllers"
	"backend/controllers/gpt"
	"backend/models"
	"encoding/json"
	"encoding/xml"
	"strings"
)

type TasksXML struct {
	XMLName xml.Name `xml:"tasks"`
	Task    []string `xml:"task"`
}

type GenTaskRequest struct {
	Prompt string `json:"prompt"`
	Slide  string `json:"slide"`
}

func (this *Controller) GenTasks() {

	var request GenTaskRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&request)

	// 生成任务列表
	template := conf.GetTasksGeneratePromptTemplate()
	template = strings.ReplaceAll(template, "{{prompt}}", request.Prompt)

	tasksXML, err := gpt.RequestGpt(template, TasksXML{}) // <tasks></tasks>
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	tasksXML = models.ReformatXML(tasksXML)

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", tasksXML)
	this.ServeJSON()

	// 将xml格式的任务列表转换为结构体

	// 执行每个任务

	// 进行任务评价

}
