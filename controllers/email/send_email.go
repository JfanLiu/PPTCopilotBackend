package email

import (
	"backend/controllers"
	"backend/models"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type SendEmailRequest struct {
	Email string `json:"email"`
}

// SendEmail 向指定邮箱发送验证码
func (this *Controller) SendEmail() {
	var request SendEmailRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&request)
	email := request.Email

	// _, err := models.GetUserByEmail(email)
	// if err != nil {
	// 	this.Data["json"] = controllers.MakeResponse(controllers.Err, "邮箱不存在", email)
	// 	this.ServeJSON()
	// 	return
	// }

	// 生成验证码
	//verificationCode := uuid.New().String()
	verificationCode := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	// 将验证码保存到缓存中
	err := models.SetCodeCache(email, verificationCode)
	if err != nil {
		// 保存验证码到缓存失败
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "验证码缓存失败", nil)
		this.ServeJSON()
		return
	}

	// 发送邮件
	err = models.SendEmail(email, verificationCode)

	if err != nil {
		// 发送邮件失败
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "邮件发送失败", email)
		this.ServeJSON()
		return
	}

	// 发送成功
	this.Data["json"] = controllers.MakeResponse(controllers.OK, "邮件发送成功", nil)
	this.ServeJSON()

}
