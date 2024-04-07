package agent

import (
	"backend/conf"
	"backend/controllers"
	"backend/controllers/gpt"
	"encoding/json"
	"fmt"
	"strings"
)

type Task struct {
	TaskName string `json:"task_name"`
	Prompt   string `json:"prompt"`
}

type GenTaskRequest struct {
	Slide  string `json:"slide"`
	Prompt string `json:"prompt"`
}

func (this *Controller) GenTasks() {

	var request GenTaskRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&request)

	// 生成任务列表
	template := conf.GetTasksGeneratePromptTemplate()
	template = strings.ReplaceAll(template, "{{prompt}}", request.Prompt)
	template = strings.ReplaceAll(template, "{{slide}}", request.Slide)

	tasksStr, err := gpt.RequestGpt(template)

	fmt.Println(tasksStr)

	var tasks []Task
	err = json.Unmarshal([]byte(tasksStr), &tasks)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", tasks)
	this.ServeJSON()

}
