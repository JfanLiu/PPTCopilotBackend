package agent

import (
	"backend/conf"
	"backend/controllers"
	"backend/controllers/gpt"
	"encoding/json"
	"strings"
	"fmt"
)

// 添加文本框时，该页ppt内现有的文本框
type TextBox struct {
	Top     float64 `json:"top"`
	Left    float64 `json:"left"`
	Width   float64 `json:"width"`
	Height  float64 `json:"height"`
	Rotate  float64 `json:"rotate"`
	Content string  `json:"content"`
}

type AddTextRequest struct {
	Prompt  string    `json:"prompt"`
	TextNow []TextBox `json:"textnow"`
}

func (this *Controller) AddText() {

	var request AddTextRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&request)

	template := conf.GetAddTextPromptTemplate()
	template = strings.ReplaceAll(template, "{{prompt}}", request.Prompt)

	textsJson, _ := json.Marshal(request.TextNow)
	textsStr := string(textsJson)
	textsStr = strings.Replace(textsStr, "\\u003c", "<", -1)
	textsStr = strings.Replace(textsStr, "\\u003e", ">", -1)
	textsStr = strings.Replace(textsStr, "\\u0026", "&", -1)

	template = strings.ReplaceAll(template, "{{slide}}", textsStr)
	newTextStr, err := gpt.RequestGpt(template)
	fmt.Println(newTextStr, err)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	var newText []TextBox
	err = json.Unmarshal([]byte(newTextStr), &newText)
	fmt.Println(2,err)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	fmt.Println(newText)

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", newText)
	this.ServeJSON()

}
