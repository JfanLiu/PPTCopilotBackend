package models

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/beego/beego/v2/client/orm"
)

func jsonObject2string(obj JsonObject) string {
	ret, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(ret)
}

// 初始化数据表
func init() {

	// ---------------------------------
	// 			调试请使用以下代码
	// ---------------------------------

	// 获取配置文件中的数据库地址
	//mysqlUrls, err := beego.AppConfig.String("mysqlurls")
	//fmt.Println("MYSQL_Urls: ", mysqlUrls)
	//orm.RegisterDataBase("default", "mysql", "root:admin@tcp("+mysqlUrls+")/now_db?charset=utf8&loc=Local")

	// ---------------------------------
	// 		 docker部署请使用以下代码
	// ---------------------------------

	// 获取环境变量
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")
	fmt.Println("MYSQL_HOST: ", mysqlHost)
	fmt.Println("MYSQL_PORT: ", mysqlPort)
	orm.RegisterDataBase("default", "mysql", "root:admin@tcp("+mysqlHost+":"+mysqlPort+")/now_db?charset=utf8&loc=Local")

	// 注册定义的model
	orm.RegisterModel(new(Outline))
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Project))
	orm.RegisterModel(new(File))
	orm.RegisterModel(new(Favorite))
	orm.RegisterModel(new(Template))
	orm.RegisterModel(new(History))

	// 如果表不存在则创建表
	orm.RunSyncdb("default", false, true)

	// 读取scripts下jianyuezi.json文件
	jianyuezi, err := os.Open("./scripts/jianyuezi.json")
	if err != nil {
		fmt.Println(err)
	}
	var jianyueziJson []JsonObject
	// 解析json文件
	json.NewDecoder(jianyuezi).Decode(&jianyueziJson)
	template1 := Template{
		Name:       "简约紫",
		Cover:      jsonObject2string(jianyueziJson[0]),
		Transition: jsonObject2string(jianyueziJson[1]),
		Catalog_3:  jsonObject2string(jianyueziJson[2]),
		Catalog_4:  jsonObject2string(jianyueziJson[3]),
		Catalog_5:  jsonObject2string(jianyueziJson[4]),
		Content_1:  jsonObject2string(jianyueziJson[5]),
		Content_2:  jsonObject2string(jianyueziJson[6]),
		Content_3:  jsonObject2string(jianyueziJson[7]),
		Content_4:  jsonObject2string(jianyueziJson[8]),
		Thank:      jsonObject2string(jianyueziJson[9]),
	}
	CreateTemplate(template1)

	// 读取scripts下taikongren.json
	taikongren, err := os.Open("./scripts/taikongren.json")
	if err != nil {
		fmt.Println(err)
	}
	var taikongrenJson []JsonObject
	// 解析json文件
	json.NewDecoder(taikongren).Decode(&taikongrenJson)
	template2 := Template{
		Name:       "太空人",
		Cover:      jsonObject2string(taikongrenJson[0]),
		Transition: jsonObject2string(taikongrenJson[1]),
		Catalog_3:  jsonObject2string(taikongrenJson[2]),
		Catalog_4:  jsonObject2string(taikongrenJson[3]),
		Catalog_5:  jsonObject2string(taikongrenJson[4]),
		Content_1:  jsonObject2string(taikongrenJson[5]),
		Content_2:  jsonObject2string(taikongrenJson[6]),
		Content_3:  jsonObject2string(taikongrenJson[7]),
		Content_4:  jsonObject2string(taikongrenJson[8]),
		Thank:      jsonObject2string(taikongrenJson[9]),
	}
	CreateTemplate(template2)

	// 读取scripts下dangjian.json
	dangjian, err := os.Open("./scripts/dangjian.json")
	if err != nil {
		fmt.Println(err)
	}

	var dangjianJson []JsonObject
	// 解析json文件

	json.NewDecoder(dangjian).Decode(&dangjianJson)

	template3 := Template{
		Name:       "党建",
		Cover:      jsonObject2string(dangjianJson[0]),
		Transition: jsonObject2string(dangjianJson[1]),
		Catalog_3:  jsonObject2string(dangjianJson[2]),
		Catalog_4:  jsonObject2string(dangjianJson[3]),
		Catalog_5:  jsonObject2string(dangjianJson[4]),
		Content_1:  jsonObject2string(dangjianJson[5]),
		Content_2:  jsonObject2string(dangjianJson[6]),
		Content_3:  jsonObject2string(dangjianJson[7]),
		Content_4:  jsonObject2string(dangjianJson[8]),
		Thank:      jsonObject2string(dangjianJson[9]),
	}
	CreateTemplate(template3)

	// 读取scripts下jiaoyu.json
	jiaoyu, err := os.Open("./scripts/jiaoyu.json")
	if err != nil {
		fmt.Println(err)
	}

	var jiaoyuJson []JsonObject
	// 解析json文件

	json.NewDecoder(jiaoyu).Decode(&jiaoyuJson)

	template4 := Template{
		Name:       "教育",
		Cover:      jsonObject2string(jiaoyuJson[0]),
		Transition: jsonObject2string(jiaoyuJson[1]),
		Catalog_3:  jsonObject2string(jiaoyuJson[2]),
		Catalog_4:  jsonObject2string(jiaoyuJson[3]),
		Catalog_5:  jsonObject2string(jiaoyuJson[4]),
		Content_1:  jsonObject2string(jiaoyuJson[5]),
		Content_2:  jsonObject2string(jiaoyuJson[6]),
		Content_3:  jsonObject2string(jiaoyuJson[7]),
		Content_4:  jsonObject2string(jiaoyuJson[8]),
		Thank:      jsonObject2string(jiaoyuJson[9]),
	}
	CreateTemplate(template4)

	// 读取scripts下shouhuifeng.json
	shouhuifeng, err := os.Open("./scripts/shouhuifeng.json")
	if err != nil {
		fmt.Println(err)
	}

	var shouhuifengJson []JsonObject
	// 解析json文件

	json.NewDecoder(shouhuifeng).Decode(&shouhuifengJson)

	template5 := Template{
		Name:       "手绘风",
		Cover:      jsonObject2string(shouhuifengJson[0]),
		Transition: jsonObject2string(shouhuifengJson[1]),
		Catalog_3:  jsonObject2string(shouhuifengJson[2]),
		Catalog_4:  jsonObject2string(shouhuifengJson[3]),
		Catalog_5:  jsonObject2string(shouhuifengJson[4]),
		Content_1:  jsonObject2string(shouhuifengJson[5]),
		Content_2:  jsonObject2string(shouhuifengJson[6]),
		Content_3:  jsonObject2string(shouhuifengJson[7]),
		Content_4:  jsonObject2string(shouhuifengJson[8]),
		Thank:      jsonObject2string(shouhuifengJson[9]),
	}
	CreateTemplate(template5)

	// 读取scripts下keai.json
	keai, err := os.Open("./scripts/keai.json")
	if err != nil {
		fmt.Println(err)
	}

	var keaiJson []JsonObject
	// 解析json文件

	json.NewDecoder(keai).Decode(&keaiJson)

	template6 := Template{
		Name:       "可爱",
		Cover:      jsonObject2string(keaiJson[0]),
		Transition: jsonObject2string(keaiJson[1]),
		Catalog_3:  jsonObject2string(keaiJson[2]),
		Catalog_4:  jsonObject2string(keaiJson[3]),
		Catalog_5:  jsonObject2string(keaiJson[4]),
		Content_1:  jsonObject2string(keaiJson[5]),
		Content_2:  jsonObject2string(keaiJson[6]),
		Content_3:  jsonObject2string(keaiJson[7]),
		Content_4:  jsonObject2string(keaiJson[8]),
		Thank:      jsonObject2string(keaiJson[9]),
	}
	CreateTemplate(template6)

	// 读取scripts下xiantiao.json
	xiantiao, err := os.Open("./scripts/xiantiao.json")
	if err != nil {
		fmt.Println(err)
	}

	var xiantiaoJson []JsonObject
	// 解析json文件

	json.NewDecoder(xiantiao).Decode(&xiantiaoJson)

	template7 := Template{
		Name:       "线条",
		Cover:      jsonObject2string(xiantiaoJson[0]),
		Transition: jsonObject2string(xiantiaoJson[1]),
		Catalog_3:  jsonObject2string(xiantiaoJson[2]),
		Catalog_4:  jsonObject2string(xiantiaoJson[3]),
		Catalog_5:  jsonObject2string(xiantiaoJson[4]),
		Content_1:  jsonObject2string(xiantiaoJson[5]),
		Content_2:  jsonObject2string(xiantiaoJson[6]),
		Content_3:  jsonObject2string(xiantiaoJson[7]),
		Content_4:  jsonObject2string(xiantiaoJson[8]),
		Thank:      jsonObject2string(xiantiaoJson[9]),
	}
	CreateTemplate(template7)

	// 读取scripts下shuicai.json
	shuicai, err := os.Open("./scripts/shuicai.json")
	if err != nil {
		fmt.Println(err)
	}

	var shuicaiJson []JsonObject
	// 解析json文件

	json.NewDecoder(shuicai).Decode(&shuicaiJson)

	template8 := Template{
		Name:       "水彩",
		Cover:      jsonObject2string(shuicaiJson[0]),
		Transition: jsonObject2string(shuicaiJson[1]),
		Catalog_3:  jsonObject2string(shuicaiJson[2]),
		Catalog_4:  jsonObject2string(shuicaiJson[3]),
		Catalog_5:  jsonObject2string(shuicaiJson[4]),
		Content_1:  jsonObject2string(shuicaiJson[5]),
		Content_2:  jsonObject2string(shuicaiJson[6]),
		Content_3:  jsonObject2string(shuicaiJson[7]),
		Content_4:  jsonObject2string(shuicaiJson[8]),
		Thank:      jsonObject2string(shuicaiJson[9]),
	}
	CreateTemplate(template8)

	// 读取scripts下yishu.json
	yishu, err := os.Open("./scripts/yishu.json")
	if err != nil {
		fmt.Println(err)
	}

	var yishuJson []JsonObject
	// 解析json文件

	json.NewDecoder(yishu).Decode(&yishuJson)

	template9 := Template{
		Name:       "艺术",
		Cover:      jsonObject2string(yishuJson[0]),
		Transition: jsonObject2string(yishuJson[1]),
		Catalog_3:  jsonObject2string(yishuJson[2]),
		Catalog_4:  jsonObject2string(yishuJson[3]),
		Catalog_5:  jsonObject2string(yishuJson[4]),
		Content_1:  jsonObject2string(yishuJson[5]),
		Content_2:  jsonObject2string(yishuJson[6]),
		Content_3:  jsonObject2string(yishuJson[7]),
		Content_4:  jsonObject2string(yishuJson[8]),
		Thank:      jsonObject2string(yishuJson[9]),
	}
	CreateTemplate(template9)

	// 读取scripts下lengdan.json
	lengdan, err := os.Open("./scripts/lengdan.json")
	if err != nil {
		fmt.Println(err)
	}

	var lengdanJson []JsonObject
	// 解析json文件

	json.NewDecoder(lengdan).Decode(&lengdanJson)

	template10 := Template{
		Name:       "冷淡",
		Cover:      jsonObject2string(lengdanJson[0]),
		Transition: jsonObject2string(lengdanJson[1]),
		Catalog_3:  jsonObject2string(lengdanJson[2]),
		Catalog_4:  jsonObject2string(lengdanJson[3]),
		Catalog_5:  jsonObject2string(lengdanJson[4]),
		Content_1:  jsonObject2string(lengdanJson[5]),
		Content_2:  jsonObject2string(lengdanJson[6]),
		Content_3:  jsonObject2string(lengdanJson[7]),
		Content_4:  jsonObject2string(lengdanJson[8]),
		Thank:      jsonObject2string(lengdanJson[9]),
	}
	CreateTemplate(template10)

}
