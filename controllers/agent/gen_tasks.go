package agent

import (
	"backend/conf"
	"backend/controllers"
	"backend/controllers/gpt"
	"encoding/json"
	"strings"
)

type GenTaskRequest struct {
	Prompt string `json:"prompt"`
}

func (this *Controller) GenTasks() {

	var request GenTaskRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&request)

	// 生成任务列表
	template := conf.GetTasksGeneratePromptTemplate()
	template = strings.ReplaceAll(template, "{{prompt}}", request.Prompt)

	tasksStr, err := gpt.RequestGpt(template)

	var tasks []string
	err = json.Unmarshal([]byte(tasksStr), &tasks)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", tasks)
	this.ServeJSON()

}
