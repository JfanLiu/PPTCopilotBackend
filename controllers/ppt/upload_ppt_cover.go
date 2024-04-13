package ppt

import (
	"backend/controllers"
	"strconv"
)

// 已弃用
func (this *Controller) UploadPptCover() {

	projectId, err := this.GetInt("project_id")
	filename := this.GetString("filename")
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	//file, err := models.GetFileOfProj(filename, projectId)

	//save_dir := "static/ppt_cover/" + strconv.Itoa(file.Id)

	// 递归创建目录
	//err = os.MkdirAll(save_dir, 0777)
	//if err != nil {
	//	this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
	//	this.ServeJSON()
	//	return
	//}

	// TODO:上传文件格式的校验

	file_path := "static/project/" + strconv.Itoa(projectId) + "/" + filename + "/" + "cover.png"

	// 将form-data中的文件保存到本地
	err = this.SaveToFile("uploadname", file_path)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", nil)
	this.ServeJSON()
}
