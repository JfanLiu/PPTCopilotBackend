package ppt

import (
	"backend/controllers"
	"backend/models"
	"fmt"
	"os"
)

func (this *Controller) ClonePpt() {
	// 获取用户
	token := this.Ctx.Request.Header.Get("token")
	err := models.CheckToken(token)

	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "未登录", nil)
		this.ServeJSON()
		return
	}
	userId := models.GetUserId(token)

	// 获取文件id
	fileId, err := this.GetInt("file_id")
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	// 获取ppt信息
	source_file, err := models.GetFileById(fileId)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "获取文件信息失败", nil)
		this.ServeJSON()
		return
	}

	// 获取用户的默认项目
	defaultProject, err := models.GetDefaultProjectByUser(userId)
	defaultProjectDir := models.GetProjectSaveDir(defaultProject.Id)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "无法获取用户默认项目", nil)
		this.ServeJSON()
		return
	}

	// 如果文件已存在
	_, err = models.GetFileOfProj(source_file.Name, defaultProject.Id)
	if err == nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "同名文件已经存在", nil)
		this.ServeJSON()
		return
	}

	// 创建项目文件夹
	err = os.MkdirAll(defaultProjectDir, 0777)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "创建项目文件夹失败", nil)
		this.ServeJSON()
		return
	}

	// 若为用户自己的文件，复制失败

	// 创建文件
	_, err = models.CreateFile(source_file.Name, defaultProject.Id)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "创建文件失败", nil)
		this.ServeJSON()
		return
	}

	// 创建文件夹
	saveDir := defaultProjectDir + "/" + source_file.Name
	_ = os.MkdirAll(saveDir, 0777)

	//将原文件复制入新地址
	file_path := defaultProjectDir + "/" + source_file.Name + "/" + source_file.Name
	old_file_path := models.GetProjectSaveDir(source_file.Project.Id) + "/" + source_file.Name + "/" + source_file.Name
	fmt.Println(file_path)
	fmt.Println(old_file_path)
	err = models.CopyFile(old_file_path, file_path)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "文件写入失败", nil)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", nil)
	this.ServeJSON()
}
