package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type File struct {
	Id   int
	Name string `orm:"size(100)"`
	// Description string    `orm:"size(100)"`
	Project *Project  `orm:"rel(fk)"` // 设置一对多的反向关系
	Visible bool      `orm:"default(true)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

// RefactFiles 根据project_id获取project详细信息
func RefactFiles(files []File) []File {
	for i, file := range files {
		project_temp, _ := GetProject(file.Project.Id)
		creator_temp, _ := GetUser(project_temp.Creator.Id)
		creator := User{Id: creator_temp.Id, Username: creator_temp.Username, Email: creator_temp.Email}
		project := Project{Id: project_temp.Id, Name: project_temp.Name, Creator: &creator, Star: project_temp.Star, Description: project_temp.Description, Created: project_temp.Created, Updated: project_temp.Updated}
		files[i].Name = file.Name
		files[i].Project = &project
	}
	return files
}

func GetFileOfProj(file_name string, id int) (File, error) {
	//找到该项目下的所有file
	files, err := GetAllFilesOfProj(id)

	if err == nil {
		//遍历所有file，找是否存在
		for _, file := range files {
			if file_name == file.Name {
				return file, nil
			}
		}
		return File{}, errors.New("文件不存在")
	}
	return File{}, err
}

// GetAllFilesOfProj 获取某个project下所有的file
func GetAllFilesOfProj(project_id int) ([]File, error) {
	o := orm.NewOrm()
	var files []File
	_, err := o.QueryTable("file").Filter("project_id", project_id).All(&files)
	return files, err
}

func CreateFile(name string, project_id int) (File, error) {
	o := orm.NewOrm()
	var project Project
	project.Id = project_id
	// 根据id获取项目信息
	err := o.Read(&project)
	if err != nil {
		// 项目不存在
		return File{}, err
	}

	// 文件存在，更新文件，修改更新时间
	_file, _ := GetFileOfProj(name, project_id)
	if _file.Name == name {
		//更新时间
		_file.Updated = time.Now()
		_, err = o.Update(&_file)
		return _file, err
	}

	// 文件不存在，创建文件
	file := File{Name: name, Project: &project, Visible: true}
	// 创建文件结构
	_, err = o.Insert(&file)
	return file, err
}

func CreatePptFile(name string, project_id int, visible bool) (File, error) {
	o := orm.NewOrm()
	var project Project
	project.Id = project_id
	// 根据id获取项目信息
	err := o.Read(&project)
	if err != nil {
		// 项目不存在
		return File{}, err
	}

	// 文件存在，更新文件，修改更新时间
	_file, _ := GetFileOfProj(name, project_id)
	if _file.Name == name {
		//更新时间
		_file.Updated = time.Now()
		_, err = o.Update(&_file)
		return _file, err
	}

	// 文件不存在，创建文件
	file := File{Name: name, Project: &project, Visible: visible}
	// 创建文件结构
	_, err = o.Insert(&file)

	return file, err
}

func DeleteDir(project_id int) error {
	saveDir := "static/files/" + strconv.Itoa(project_id)
	err := os.RemoveAll(saveDir)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFile(file_name string, project_id int) error {
	file, err := GetFileOfProj(file_name, project_id)
	if err != nil {
		// 文件不存在
		return err
	}
	o := orm.NewOrm()
	_, err = o.Delete(&file)
	if err != nil {
		return err //文件结构删除错误
	}

	deletePath := "static/project/" + strconv.Itoa(project_id) + "/" + file_name
	err = os.RemoveAll(deletePath)
	if err != nil {
		return err
	}

	if err != nil {
		//文件删除错误
		return err
	}
	return nil
}

func DeleteFileByPath(path string) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	err = os.Remove(absPath)
	return err
}

func GetFilePathByName(file_name string, project_id int) string {

	saveDir := "static/project/" + strconv.Itoa(project_id) + "/" + file_name

	// 创建文件夹
	_ = os.MkdirAll(saveDir, 0777)

	filePath := saveDir + "/" + file_name
	return filePath
}
func GetProjectSaveDir(project_id int) string {
	saveDir := "static/project/" + strconv.Itoa(project_id)
	return saveDir
}

func CopyFile(srcFile, dstFile string) error {
	src, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}

func SaveJsonsToFile(data interface{}, file_name string, project_id int) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	saveDir := GetProjectSaveDir(project_id)
	// 创建文件夹
	err = os.MkdirAll(saveDir, 0777)
	if err != nil {
		return err
	}

	filePath := GetFilePathByName(file_name, project_id)
	// 创建文件
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入文件
	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}
	return nil

}

func UpdateFileName(project_id int, old_file_name string, new_file_name string) (File, error) {
	o := orm.NewOrm()
	// 如果新文件名已存在
	_, err := GetFileOfProj(new_file_name, project_id)
	if err == nil {
		return File{}, errors.New("文件名已存在")
	}

	// 如果新文件名为空
	if new_file_name == "" {
		return File{}, errors.New("文件名不能为空")
	}

	// 如果新文件名以.json结尾
	if strings.HasSuffix(new_file_name, ".json") {
		//如果旧文件名不以.json结尾
		if !strings.HasSuffix(old_file_name, ".json") {
			return File{}, errors.New("文件名不能以.json结尾")
		}
	}

	if strings.HasSuffix(old_file_name, ".json") {
		//如果旧文件名以.json结尾
		if !strings.HasSuffix(new_file_name, ".json") {
			return File{}, errors.New("文件名必须以.json结尾")
		}
	}

	file, err := GetFileOfProj(old_file_name, project_id)
	if err != nil {
		return File{}, err
	}
	file.Name = new_file_name
	_, err = o.Update(&file)
	// 更新文件名
	if err != nil {
		return File{}, err
	}

	// 更新文件路径
	old_path := "static/project/" + strconv.Itoa(project_id) + "/" + old_file_name + "/" + old_file_name
	new_path := "static/project/" + strconv.Itoa(project_id) + "/" + old_file_name + "/" + new_file_name
	err = os.Rename(old_path, new_path)
	if err != nil {
		return File{}, err
	}

	// 更新文件夹路径
	old_dir := "static/project/" + strconv.Itoa(project_id) + "/" + old_file_name
	new_dir := "static/project/" + strconv.Itoa(project_id) + "/" + new_file_name
	err = os.Rename(old_dir, new_dir)
	if err != nil {
		return File{}, err
	}

	return file, err
}

// 获取所有公开的ppt
func GetAllPublicPpt() ([]File, error) {
	o := orm.NewOrm()
	var files []File
	_, err := o.QueryTable("file").Filter("name__endswith", ".json").Filter("visible", 1).All(&files)
	return files, err
}

// 搜索所有公开ppt
func SearchAllPublicPpt(keywords []string) ([]File, error) {
	o := orm.NewOrm()

	// 构造 SQL 语句
	sql := `SELECT * FROM file WHERE visible=true AND name LIKE '%.json'`
	for _, keyword := range keywords {
		// 使用 OR 连接多个关键词
		sql += fmt.Sprintf("AND (name LIKE '%%%s%%')", keyword)
	}

	// 执行 SQL 查询
	var files []File
	_, err := o.Raw(sql).QueryRows(&files)
	if err != nil {
		return nil, err
	}

	return files, nil
}

// 根据id获取文件
func GetFileById(fileId int) (File, error) {
	o := orm.NewOrm()
	var file File
	err := o.QueryTable("file").Filter("id", fileId).One(&file)
	return file, err
}
