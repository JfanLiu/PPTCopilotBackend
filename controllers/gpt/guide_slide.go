package gpt

import (
	"backend/conf"
	"backend/controllers"
	"encoding/json"
	"strings"
)

type GuideSlideRequest struct {
	Outline string `json:"outline"`
}

// GuideSlide 用于生成GuideSlide
func (this *Controller) GuideSlide() {
	var request GuideSlideRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&request)

	// 获取指导单页提示模板，并替换模板中的 "{{outline}}" 占位符为请求中的大纲内容
	template := conf.GetGuideSinglePromptTemplate()
	template = strings.ReplaceAll(template, "{{outline}}", request.Outline)

	// 调用 RequestGptXml 函数生成GuideSlide内容
	guide_slide, err := RequestGptXml(template, SectionXML{}) // <section></section>
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	guide_slide = strings.ReplaceAll(guide_slide, "\n", "")

	// 返回生成的GuideSlide
	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", guide_slide)
	this.ServeJSON()
}
