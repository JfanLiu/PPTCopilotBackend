package agent

import (
	"backend/controllers"
	"backend/models"
	"encoding/json"
	"fmt"

	"github.com/alibabacloud-go/tea/tea"
	ha3engine "github.com/aliyun/alibabacloud-ha3-go-sdk/client"
)

type GetImageRequest struct {
	Keyword string `json:"keyword"`
}

func (this *Controller) GetImage() {
	var request GetImageRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&request)

	fmt.Println(request)
	images := QueryImage(request.Keyword)

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", images)
	this.ServeJSON()
}

type OpensearchResult struct {
	Id     int     `json:"id"`
	Score  float32 `json:"score"`
	Source float32 `json:"__source__"`
}

type OpensearchResponse struct {
	TotalCount int                `json:"totalCount"`
	Result     []OpensearchResult `json:"result"`
	TotalTime  float32            `json:"totalTime"`
}

func QueryImage(context string) []string {

	// 创建请求客户端配置
	config := &ha3engine.Config{
		// 私网域名或者公网域名
		Endpoint: tea.String("ha-cn-jte3ona1w01.public.ha.aliyuncs.com"),
		//  实例名称，可在实例详情页左上角查看，例:ha-cn-i7*****605
		InstanceId: tea.String("ha-cn-jte3ona1w01"),
		// 用户名，可在实例详情页>网络信息 查看
		AccessUserName: tea.String("tongji"),
		// 密码，可在实例详情页>网络信息 修改
		AccessPassWord: tea.String("tongji"),
	}

	// 初始化一个client, 用以发送请求.
	client, _clientErr := ha3engine.NewClient(config)

	// 如果 NewClient 过程中出现异常. 则 返回 _clientErr 且输出 错误信息.
	if _clientErr != nil {
		fmt.Println(_clientErr)
		return nil
	}

	return inferenceQuery(client, context)
}

/**
 *	预测查询
 */
func inferenceQuery(client *ha3engine.Client, context string) []string {

	searchRequestModel := &ha3engine.QueryRequest{}
	searchRequestModel.SetTableName("image")
	searchRequestModel.SetModal("text")
	searchRequestModel.SetContent(context)
	searchRequestModel.SetTopK(3)
	searchRequestModel.SetSearchParams("{\"qc.searcher.scan_ratio\":0.01}")
	searchRequestModel.SetIncludeVector(false)

	response, _requestErr := client.InferenceQuery(searchRequestModel)

	// 如果 发送请求 过程中出现异常. 则 返回 _requestErr 且输出 错误信息.
	if _requestErr != nil {
		fmt.Println(_requestErr)
		return nil
	}

	//testResponse := `{"totalCount":1,"result":[{"id":0,"score":1.139796,"__source__":1}],"totalTime":1.548}`

	// 将response body转为结构体
	var opensearchResponse OpensearchResponse
	var _ = json.Unmarshal([]byte(*response.Body), &opensearchResponse)

	// 输出正常返回的 response 内容
	fmt.Println(*response.Body)
	fmt.Println(opensearchResponse)

	// 去数据库中取图片base64
	var images []string
	for i := 0; i < len(opensearchResponse.Result); i++ {
		imageBase64, _ := models.GetBase64ById(opensearchResponse.Result[i].Id)
		images = append(images, imageBase64)
	}

	return images
}
