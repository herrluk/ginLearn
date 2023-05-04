package admin

import (
	"fmt"
	"github.com/herrluk/goProjectsPractice/ginLearn/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	BaseController
}

func (con RoleController) Index(c *gin.Context) {
	var roleList []models.Role
	models.DB.Find(&roleList)

	c.HTML(http.StatusOK, "admin/role/index.html", gin.H{
		"roleList": roleList,
	})

}
func (con RoleController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/role/add.html", gin.H{})
}
func (con RoleController) DoAdd(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	if title == "" {
		con.Error(c, "角色标题不能为空！", "/admin/role/add")
		return
	}
	role := models.Role{
		Title:       title,
		Description: description,
		Status:      1,
		AddTime:     int(models.GetUnix()),
	}

	err := models.DB.Create(&role).Error
	if err != nil {
		con.Error(c, "增加角色失败，请重试！", "/admin/role/add")
	} else {
		con.Success(c, "增加角色成功！", "/admin/role")

	}
}

func (con RoleController) Edit(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误，请重试！", "/admin/role")
		return
	} else {
		role := models.Role{Id: id}
		models.DB.Find(&role)
		c.HTML(http.StatusOK, "admin/role/edit.html", gin.H{
			"role": role,
		})
	}

}
func (con RoleController) DoEdit(c *gin.Context) {

	id, err := models.Int(c.PostForm("id"))
	if err != nil {
		fmt.Println("转换 id 错误：", err)
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}
	// 去掉空格
	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")
	if title == "" {
		con.Error(c, "角色标题不能为空！", "/admin/role/edit")
	}
	role := models.Role{Id: id}

	models.DB.Find(&role)
	role.Title = title
	role.Description = description

	err = models.DB.Save(&role).Error
	if err != nil {
		fmt.Println("doedit错误：", err)
		con.Error(c, "传入数据错误", "/admin/role/edit?id="+models.String(id))
	} else {
		fmt.Println("doedit成功：", err)
		con.Success(c, "修改数据成功", "/admin/role/edit?id="+models.String(id))
	}
	// c.String(http.StatusOK, "DoEdit")
}
func (con RoleController) Delete(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		fmt.Println("转换 id 错误：", err)
		con.Error(c, "传入数据错误", "/admin/role")
	} else {
		role := models.Role{Id: id}
		err = models.DB.Delete(&role).Error
		if err != nil {
			fmt.Println("delete错误：", err)
			con.Error(c, "删除数据错误", "/admin/role/edit?id="+models.String(id))
		} else {
			con.Success(c, "删除数据成功", "/admin/role")
		}
	}

}
