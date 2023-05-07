package admin

import (
	"fmt"
	"github.com/herrluk/goProjectsPractice/ginLearn/models"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AccessController struct {
	BaseController
}

func (con AccessController) Index(c *gin.Context) {
	var accessList []models.Access
	models.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList)

	fmt.Printf("%#v", accessList)
	c.HTML(http.StatusOK, "admin/access/index.html", gin.H{
		"accessList": accessList,
	})

}
func (con AccessController) Add(c *gin.Context) {
	// 获取顶级模块
	var accessList []models.Access
	models.DB.Where("module_id=?", 0).Find(&accessList)
	c.HTML(http.StatusOK, "admin/access/add.html", gin.H{
		"accessList": accessList,
	})
}
func (con AccessController) DoAdd(c *gin.Context) {
	moduleName := strings.Trim(c.PostForm("module_name"), " ")
	actionName := c.PostForm("action_name")
	accessType, err1 := models.Int(c.PostForm("type"))
	url := c.PostForm("url")
	moduleId, err2 := models.Int(c.PostForm("module_id"))
	sort, err3 := models.Int(c.PostForm("sort"))
	status, err4 := models.Int(c.PostForm("status"))
	description := c.PostForm("description")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		con.Error(c, "传入参数错误", "/admin/access/add")
		return
	}
	if moduleName == "" {
		con.Error(c, "模块名称不能为空", "/admin/access/add")
		return
	}

	access := models.Access{
		ModuleName:  moduleName,
		Type:        accessType,
		ActionName:  actionName,
		Url:         url,
		ModuleId:    moduleId,
		Sort:        sort,
		Description: description,
		Status:      status,
	}
	err5 := models.DB.Create(&access).Error
	if err5 != nil {
		con.Error(c, "增加数据失败", "/admin/access/add")
		return
	}
	con.Success(c, "增加数据成功", "/admin/access")

}
