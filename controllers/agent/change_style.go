package agent

import (
	"backend/conf"
	"backend/controllers"
	"backend/controllers/gpt"
	"encoding/json"
	"fmt"
	"strings"
)

type Style struct {
	Id      string `json:"id"`
	Content string `json:"content"`
}

type ChangeStyleRequest struct {
	Prompt string  `json:"prompt"`
	Slide  []Style `json:"slide"`
}

func (this *Controller) ChangeStyle() {

	var request ChangeStyleRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&request)

	// 生成任务列表
	template := conf.GetChangeStylePromptTemplate()
	template = strings.ReplaceAll(template, "{{prompt}}", request.Prompt)

	slideJson, _ := json.Marshal(request.Slide)
	slideStr := string(slideJson)
	slideStr = strings.Replace(slideStr, "\\u003c", "<", -1)
	slideStr = strings.Replace(slideStr, "\\u003e", ">", -1)
	slideStr = strings.Replace(slideStr, "\\u0026", "&", -1)

	template = strings.ReplaceAll(template, "{{slide}}", slideStr)

	changedSlideStr, err := gpt.RequestGpt(template)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	var styles []Style
	err = json.Unmarshal([]byte(changedSlideStr), &styles)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	fmt.Println(styles)

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", styles)
	this.ServeJSON()
}
