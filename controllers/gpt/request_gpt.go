package gpt

import (
	"backend/conf"
	"backend/models"
	"encoding/xml"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"unicode"

	"github.com/imroc/req"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestBody struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
}
type ResponseBody struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

type SlidesXML struct {
	XMLName  xml.Name     `xml:"slides"`
	Sections []SectionXML `xml:"section"`
}

type SectionXML struct {
	XMLName xml.Name `xml:"section"`
	Class   string   `xml:"class,attr"`
	Content []string `xml:"p"`
}

func ErrorScanner(gptResponse string, genXMLType interface{}) (string, error) {
	// 确保程序按照genXMLType的类型进行解析
	value := reflect.New(reflect.TypeOf(genXMLType)).Interface()
	err := xml.Unmarshal([]byte(gptResponse), value)
	if err != nil {
		return "", fmt.Errorf("genXMLType格式与gptResponse不匹配: %s", err.Error())
	}

	// 移除所有转义字符
	r := regexp.MustCompile(`\\.`)
	gptResponse = r.ReplaceAllString(gptResponse, "")

	// 移除所有不必要空格，"<>"标签内内容除外
	inScope := false
	var builder strings.Builder
	for _, ch := range gptResponse {
		if ch == '<' {
			inScope = true
			builder.WriteRune(ch)
			continue
		}
		if ch == '>' {
			inScope = false
			builder.WriteRune(ch)
			continue
		}
		if !inScope && unicode.IsSpace(ch) {
			continue
		}
		builder.WriteRune(ch)
	}

	return builder.String(), nil
}

func RequestGpt(prompt string, genXmlType interface{}) (string, error) {
	var apikey string

	// 当GetApiKey()得到合法的apikey时，才进行下一步
	for apikey == "" {
		apikey, _ = GetApiKey()
	}
	fmt.Println("apikey: " + apikey)
	defer ReleaseApiKey(apikey)

	var body RequestBody
	body.Model = conf.GetGptModel()
	body.Messages = append(body.Messages, Message{
		Role:    "user",
		Content: prompt,
	})
	body.Temperature = 0.7

	fmt.Println(body)

	req_url := conf.GetGptApiUrl()
	//req.SetProxyUrl("http://host.docker.internal:7890")
	req.SetProxyUrl(conf.GetGptProxy())

	// 最多尝试2次
	retryCount := 0

	for retryCount < 2 {
		// 进行http请求
		resp, err := req.Post(req_url, req.BodyJSON(&body), req.Header(map[string]string{
			"Authorization": "Bearer " + apikey,
			"Content-Type":  "application/json",
		}))
		if err != nil {
			return "", err
		}
		if resp.Response().StatusCode != 200 {
			return "", fmt.Errorf("GPT请求失败，状态码为%d", resp.Response().StatusCode)
		}

		var res ResponseBody
		resp.ToJSON(&res)

		// 扫描错误
		result, err := ErrorScanner(res.Choices[0].Message.Content, genXmlType)
		if err != nil {
			retryCount++
			continue
		}

		// 成功
		return result, nil
	}

	// 3次尝试均失败
	return "", fmt.Errorf("all retries failed")
}

// GuideContentSection 函数丰富每个内容页
func GuideContentSection(outline string) (string, error) {

	debug := 1 // 调试模式
	if debug == 1 {
		guide_slide := "<section class='content'><p>社交互动</p><p>1. 分享自己在游戏中的心得体会有助于与其他玩家建立更紧密的联系，增强游戏体验。</p><p>2. 参与游戏社区的互动活动，不仅可以赢取奖励，还能结交志同道合的朋友。</p><p>3. 玩家之间的互动是游戏中不可或缺的一部分，可以互相帮助、交流游戏心得、组队挑战副本等。</p><p>4. 玩家之间的互动是游戏中不可或缺的一部分，可以互相帮助、交流游戏心得、组队挑战副本等。</p><p>5. 玩家之间的互动是游戏中不可或缺的一部分，可以互相帮助、交流游戏心得、组队挑战副本等。</p></section>"
		return guide_slide, nil
	} else {
		promptTemplate := conf.GetGuideSinglePromptTemplate()
		prompt := strings.ReplaceAll(promptTemplate, "{{outline}}", outline)

		guide_slide, err := RequestGpt(prompt, SectionXML{}) // <section></section>
		if err != nil {
			return "", err
		}

		models.StrDeleteLineBreak(guide_slide)
		return guide_slide, nil
	}

}
