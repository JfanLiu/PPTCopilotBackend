package gpt

import (
	"backend/controllers"
	"backend/models"
	"encoding/json"
	"fmt"
)

// GenPPTRequest 结构体定义了一个用于生成 PPT 的请求结构体
type GenPPTRequest struct {
	OutlineId  int `json:"outline_id"`  // 大纲 ID
	TemplateId int `json:"template_id"` // 模板 ID
	//ProjectId  int    `json:"project_id"`  // 项目 ID
	FileName string `json:"file_name"` // 文件名
	Visible  bool   `json:"visible"`   // 可见性（公开或私有）
}

// GenPPT 方法用于生成 PPT
func (this *Controller) GenPPT() {
	var request GenPPTRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&request)

	// 获取操作用户
	token := this.Ctx.Request.Header.Get("token")
	fmt.Println(token)
	err := models.CheckToken(token)
	fmt.Println(err)

	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "未登录", nil)
		this.ServeJSON()
		return
	}
	// userId := models.GetUserId(token.Value)
	userId := models.GetUserId(token)

	// 获取请求中的大纲 ID、模板 ID、项目 ID 和文件名
	outlineId := request.OutlineId
	templateId := request.TemplateId
	//projectId := request.ProjectId
	project, err := models.GetDefaultProjectByUser(userId)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}
	projectId := project.Id
	fileName := request.FileName

	// 从数据库中根据id获取大纲outline和模板template
	outline, err := models.GetOutline(outlineId)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}
	template, err := models.GetTemplate(templateId)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), template)
		this.ServeJSON()
		return
	}

	debug := 0
	if debug == 1 {
		outlinexml := `<slides>
    	<section class='cover'>
        <p>我为什么玩明日方舟</p>
        <p>汇报人：dhf</p>
    	</section>
    	<section class='catalog'>
        <p>目录</p>
        <p>1. 游戏背景</p>
        <p>2. 独特的人设</p>
        <p>3. 挑战性的玩法</p>
        <p>4. 社区互动</p>
        <p>5. 结语</p>
    	</section>
    	<section class='content'>
        <p>游戏背景</p>
        <p>内容概要：介绍明日方舟的世界观，讲述感染者与感染病毒的斗争，以及玩家在游戏中扮演的医疗救援人员的角色。</p>
    	</section>
    	<section class='content'>
        <p>独特的人设</p>
        <p>内容概要：分享明日方舟各种不同种族，不同职业的角色形象，以及他们的个性、故事和能力等。</p>
    	</section>
    	<section class='content'>
        <p>挑战性的玩法</p>
        <p>内容概要：介绍明日方舟各种游戏模式和关卡，以及它们的难度和挑战性，包括如何提高玩家的战斗技巧和策略。</p>
		</section>
		<section class='content'>
        <p>社区互动</p>
        <p>内容概要：讲述明日方舟社区的特点和优势，包括玩家之间的互动、交流和创作，以及开发团队与玩家之间的沟通和回应。</p>
		</section>
		<section class='content'>
        <p>结语</p>
        <p>内容概要：总结明日方舟的优点和吸引力，以及我个人对它的喜爱和热爱，鼓励更多的人加入明日方舟的世界。</p>
		</section>
		</slides>`

		// 从大纲中获取所有的ContentSections
		contentSections, err := models.GetContentSections(outlinexml)
		if err != nil {
			this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
			this.ServeJSON()
			return
		}

		// 所有的ContentSection进行guide_slide
		// 结果以 string 形式存储在 guide_slides 中
		guide_slides := make([]string, 0)
		for _, contentSection := range contentSections {
			guide_slide, err := GuideContentSection(contentSection)
			fmt.Println(guide_slide)
			if err != nil {
				this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
				this.ServeJSON()
				return
			}
			guide_slides = append(guide_slides, guide_slide)
		}

		// 将outline.Outline中的所有的ContentSection替换为guide_slide
		// resultxml 是内容页经过替换的大纲
		resultxml, err := models.RefactContentSections(outline.Outline, guide_slides)
		if err != nil {
			this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
			this.ServeJSON()
			return
		}

		var res []string

		res, err = models.GenPPTWithTemplate(resultxml, template)
		if err != nil {
			this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), res)
			this.ServeJSON()
			return
		}

		JsonRes := make([]models.JsonObject, len(res))
		for i, _ := range res {
			JsonRes[i] = models.GetObj(res[i])
		}

		file, err := models.CreatePptFile(fileName, projectId, request.Visible)
		if err != nil {
			this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), file)
			this.ServeJSON()
			return
		}
		err = models.SaveJsonsToFile(JsonRes, fileName, projectId)
		if err != nil {
			this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), res)
			this.ServeJSON()
			return
		}

		// 更新操作历史
		userId := 1
		err = models.UpdateHistory(userId, file.Id)

		this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", JsonRes)
		this.ServeJSON()
		return
	}

	// 从大纲中获取所有的ContentSections
	contentSections, err := models.GetContentSections(outline.Outline)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	// 所有的ContentSection进行guide_slide
	// 结果以 string 形式存储在 guide_slides 中
	guide_slides := make([]string, 0)
	for _, contentSection := range contentSections {
		guide_slide, err := GuideContentSection(contentSection)
		fmt.Println(guide_slide)
		if err != nil {
			this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
			this.ServeJSON()
			return
		}
		guide_slides = append(guide_slides, guide_slide)
	}

	// 将outline.Outline中的所有的ContentSection替换为guide_slide
	// resultxml 是内容页经过替换的大纲
	resultxml, err := models.RefactContentSections(outline.Outline, guide_slides)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	// 使用模板生成 PPT，并保存到文件中
	// res 是最终结果，string格式表示的ppt
	var res []string
	res, err = models.GenPPTWithTemplate(resultxml, template)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), res)
		this.ServeJSON()
		return
	}

	// 将生成的 PPT 结果保存到文件，并返回响应
	JsonRes := make([]models.JsonObject, len(res))
	for i, _ := range res {
		JsonRes[i] = models.GetObj(res[i])
	}

	file, err := models.CreatePptFile(fileName, projectId, request.Visible)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), file)
		this.ServeJSON()
		return
	}
	err = models.SaveJsonsToFile(JsonRes, fileName, projectId)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), res)
		this.ServeJSON()
		return
	}

	// 更新操作历史
	err = models.UpdateHistory(userId, file.Id)

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", JsonRes)
	this.ServeJSON()

}
