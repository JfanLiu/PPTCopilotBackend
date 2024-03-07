package models

import (
	"encoding/xml"
	"regexp"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

// Outline 大纲结构体
type Outline struct {
	Id      int    // 大纲 id
	Outline string `orm:"type(text)"` // 大纲内容，类型为 text
}

// TODO：错误处理可以更优雅
// GetOutline 从数据库中获取指定 Id 的大纲
func GetOutline(id int) (Outline, error) {
	o := orm.NewOrm()
	outline := Outline{Id: id}
	err := o.Read(&outline)
	return outline, err
}

// CreateOutline 创建一个新的大纲，并将其保存到数据库中
func CreateOutline(outline_str string) (Outline, error) {
	o := orm.NewOrm()
	outline := Outline{Outline: outline_str}
	_, err := o.Insert(&outline)
	return outline, err
}

// UpdateOutline 函数更新指定 Id 的大纲内容
func UpdateOutline(id int, outline string) (Outline, error) {
	o := orm.NewOrm()
	_outline := Outline{Id: id}
	err := o.Read(&_outline)
	if err != nil {
		return Outline{}, err
	}
	_outline.Outline = outline
	_, err = o.Update(&_outline)
	return _outline, err
}

// TODO：util提出来/再封装一层
// RefactOutline 函数用于将outline中的"\n"字符串删去
func RefactOutline(outline Outline) Outline {
	outline.Outline = DeleteLineBreak(outline.Outline)
	return outline
}

// TODO：合并
// DeleteLineBreak 函数用于删除字符串中的换行符
func DeleteLineBreak(outline string) string {
	outline = strings.ReplaceAll(outline, "\n", "")
	return outline
}

// RefactXML 函数用于修改 XML 字符串，删除\n和\t，并返回第一个匹配的 XML 标签
func RefactXML(xmlStr string) string {
	xmlStr = strings.ReplaceAll(xmlStr, "\n", "")
	xmlStr = strings.ReplaceAll(xmlStr, "\t", "")
	regex := "(<.*>)" // 匹配 XML 标签的正则表达式
	re := regexp.MustCompile(regex)
	matches := re.FindAllString(xmlStr, -1)
	return matches[0] // 返回第一个匹配的 XML 标签
}

// P 结构体表示 XML 中的 <p> 元素
type P struct {
	XMLName xml.Name `xml:"p"`
	Content string   `xml:",innerxml"`
}

// Slide 结构体表示 XML 中的 <section> 元素
type Slide struct {
	XMLName xml.Name `xml:"section"`    // XML 元素名称为 section
	Class   string   `xml:"class,attr"` // class 属性
	P_arr   []P      `xml:"p"`          // 段落数组，对应 <p> 元素
	Content string   `xml:",innerxml"`  // 内容字段为 XML 元素的内部 XML 字符串
}

// Slides 结构体表示整个 XML 文档
type Slides struct {
	XMLName  xml.Name `xml:"slides"`  // XML 元素名称为 slides
	Sections []Slide  `xml:"section"` // slides数组，对应 <section> 元素
}

// GetContentSections 函数从 XML 字符串中获取内容部分（即class="content"）
func GetContentSections(xmlStr string) ([]string, error) {
	var slides Slides
	err := xml.Unmarshal([]byte(xmlStr), &slides) // 解析 XML 字符串到 Slides 结构体中
	if err != nil {
		return []string{}, err
	}

	// 查询<section class='content'>的项
	var contents []string
	for _, section := range slides.Sections {
		if section.Class == "content" {
			contents = append(contents, "<section class='content'>"+section.Content+"</section>")
		}
	}
	return contents, nil
}

// RefactContentSections 函数用于修改幻灯片内容，将其中的内容部分替换为指定的内容
func RefactContentSections(xmlStr string, guide_slides []string) (string, error) {
	var slides Slides
	err := xml.Unmarshal([]byte(xmlStr), &slides)
	if err != nil {
		return "", err
	}

	// 查询<section class='content'>的项
	// 并将其内容替换为指定的内容
	var resultxml string
	//初始化count=0
	count := 0
	for _, section := range slides.Sections {
		if section.Class == "content" {
			resultxml += guide_slides[count]
			count++
		} else {
			resultxml += "<section class='" + section.Class + "'>" + section.Content + "</section>"
		}
	}

	resultxml = "<slides>" + resultxml + "</slides>" // 包装成完整的 XML 结构
	return resultxml, nil
}

// coverReplace 函数替换cover页中的 title 和 description 信息
func coverReplace(section Slide, cover string) string {
	var title = section.P_arr[0].Content
	var description = section.P_arr[1].Content

	cover = strings.ReplaceAll(cover, "{{title}}", title)
	cover = strings.ReplaceAll(cover, "{{description}}", description)

	return cover
}

func catalogReplace(section Slide, template Template) string {
	var p_num = len(section.P_arr)
	var p_list = section.P_arr
	var catalog = p_list[0].Content
	var ret = ""

	if p_num-1 == 3 {
		ret = template.Catalog_3
		ret = strings.ReplaceAll(ret, "{{item1}}", p_list[1].Content)
		ret = strings.ReplaceAll(ret, "{{item2}}", p_list[2].Content)
		ret = strings.ReplaceAll(ret, "{{item3}}", p_list[3].Content)
	} else if p_num-1 == 4 {
		ret = template.Catalog_4
		ret = strings.ReplaceAll(ret, "{{item1}}", p_list[1].Content)
		ret = strings.ReplaceAll(ret, "{{item2}}", p_list[2].Content)
		ret = strings.ReplaceAll(ret, "{{item3}}", p_list[3].Content)
		ret = strings.ReplaceAll(ret, "{{item4}}", p_list[4].Content)
	} else if p_num-1 == 5 {
		ret = template.Catalog_5
		ret = strings.ReplaceAll(ret, "{{item1}}", p_list[1].Content)
		ret = strings.ReplaceAll(ret, "{{item2}}", p_list[2].Content)
		ret = strings.ReplaceAll(ret, "{{item3}}", p_list[3].Content)
		ret = strings.ReplaceAll(ret, "{{item4}}", p_list[4].Content)
		ret = strings.ReplaceAll(ret, "{{item5}}", p_list[5].Content)
	}
	ret = strings.ReplaceAll(ret, "{{catalog}}", catalog)
	return ret

}

func contentReplace(section Slide, template Template) []string {
	var p_num = len(section.P_arr)
	var p_list = section.P_arr
	var sub_title = p_list[0].Content
	var retList []string
	// 如果p_num-1<=4，只用生成一个页面
	var index = 1
	var ret string
	ret = template.Transition
	ret = strings.ReplaceAll(ret, "{{sub_title}}", sub_title)
	retList = append(retList, ret)

	for p_num-1 > 0 {
		if p_num-1 <= 4 {
			var ret string
			if p_num-1 == 1 {
				ret = template.Content_1
			} else if p_num-1 == 2 {
				ret = template.Content_2
			} else if p_num-1 == 3 {
				ret = template.Content_3
			} else if p_num-1 == 4 {
				ret = template.Content_4
			}
			ret = strings.ReplaceAll(ret, "{{sub_title}}", sub_title)
			for i := 1; i < p_num; i++ {
				ret = strings.ReplaceAll(ret, "{{sub_title_content"+strconv.Itoa(i)+"}}", p_list[index].Content)
				index++
			}
			retList = append(retList, ret)
			index += p_num - 1
			p_num = 1
		} else if p_num-1 == 5 {
			var ret1, ret2 string
			ret1 = template.Content_3
			ret2 = template.Content_2
			ret1 = strings.ReplaceAll(ret1, "{{sub_title}}", sub_title)
			ret2 = strings.ReplaceAll(ret2, "{{sub_title}}", sub_title)
			for i := 1; i < 4; i++ {
				ret1 = strings.ReplaceAll(ret1, "{{sub_title_content"+strconv.Itoa(i)+"}}", p_list[index].Content)
				index++
			}
			for i := 4; i < 6; i++ {
				ret2 = strings.ReplaceAll(ret2, "{{sub_title_content"+strconv.Itoa(i-3)+"}}", p_list[index].Content)
				index++
			}
			retList = append(retList, ret1)
			retList = append(retList, ret2)
			p_num = 1
		} else {
			// 生成一个页面 4
			var ret string
			ret = template.Content_4
			ret = strings.ReplaceAll(ret, "{{sub_title}}", sub_title)
			for i := 1; i < 5; i++ {
				ret = strings.ReplaceAll(ret, "{{sub_title_content"+strconv.Itoa(i)+"}}", p_list[index].Content)
				index++
			}
			retList = append(retList, ret)

			p_num -= 4
		}
	}
	return retList
}

// GenPPT 函数用于生成 PPT，根据 XML 中的内容生成对应的幻灯片，并将幻灯片内容替换为模板中的内容
func GenPPT(xmlStr string, template Template) ([]string, error) {
	var slides Slides
	err := xml.Unmarshal([]byte(xmlStr), &slides)
	if err != nil {
		return []string{}, err
	}
	var ppt []string
	res := ""

	for i, section := range slides.Sections {
		if i == 0 {
			res = coverReplace(section, template.Cover) // 封面页替换
			ppt = append(ppt, res)
		} else if i == 1 {
			res = catalogReplace(section, template) // 目录页替换
			ppt = append(ppt, res)
		} else {
			res_list := contentReplace(section, template) // 内容页替换
			for _, res_item := range res_list {
				ppt = append(ppt, res_item)
			}
		}

	}
	ppt = append(ppt, template.Thank) // 添加结尾页

	return ppt, nil

}
