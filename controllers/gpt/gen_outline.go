package gpt

import (
	"backend/conf"
	"backend/controllers"
	"backend/models"
	"encoding/json"
	"strings"
)

// GenOutlineRequest 用于生成大纲的请求结构体
type GenOutlineRequest struct {
	Topic   string `json:"topic"`   // 主题
	Sponsor string `json:"sponsor"` // 发起人
}

// GenOutline 方法用于生成大纲
func (this *Controller) GenOutline() {
	var request GenOutlineRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&request)

	// 获取大纲prompt模板，并根据请求中的主题和发起人信息替换模板中的占位符
	prompt := conf.GetOutlinePromptTemplate()
	prompt = strings.ReplaceAll(prompt, "{{topic}}", request.Topic)
	prompt = strings.ReplaceAll(prompt, "{{sponsor}}", request.Sponsor)

	outline_str := ``
	debug := 1 // 调试模式
	if debug == 1 {
		// 如果调试模式开启，则使用固定的大纲字符串
		// 粥p差不多得了
		outline_str = `<slides>
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
	} else {
		// 如果调试模式未开启，则请求 GPT 模型生成大纲
		var err error
		outline_str, err = RequestGpt(prompt, SlidesXML{}) //<slide></slide>
		if err != nil {
			this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), outline_str)
			this.ServeJSON()
			return
		}
	}

	// TODO：用缓存会不会更好？
	// 将生成的大纲字符串保存到数据库中
	outline, err := models.CreateOutline(outline_str)
	if err != nil {
		// 保存大纲到数据库失败，则返回错误响应
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), outline)
		this.ServeJSON()
		return
	}

	// 返回成功响应
	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", outline)
	this.ServeJSON()
}
