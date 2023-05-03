package admin

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/herrluk/goProjectsPractice/ginLearn/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})

}
func (con LoginController) DoLogin(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username, "+", password)

	// 1.验证验证码是否正确
	captchaId := c.PostForm("captchaId")     // hidden的 id
	verifyValue := c.PostForm("verifyValue") // 前端验证码输入框的 id

	if flag := models.VerifyCaptcha(captchaId, verifyValue); flag {
		// 2.查询数据库，判断用户名密码是否正确
		var userInfo []models.Manager
		password = models.Md5(password) // 加密
		fmt.Println(username, "+", password)
		models.DB.Where("username = ? AND  password = ?", username, password).
			Find(&userInfo)
		if len(userInfo) > 0 {
			// 3. 执行登录，保存用户信息，执行跳转

			session := sessions.Default(c)
			userinfoSlice, _ := json.Marshal(userInfo)
			session.Set("userInfo", string(userinfoSlice))
			err := session.Save()
			if err != nil {
				fmt.Println(err)
				con.error(c, "session错误", "/admin/login")

			} else {
				con.success(c, "登录成功", "/admin")

			}
		} else {
			con.error(c, "用户名或密码错误", "/admin/login")
		}
	} else {
		con.error(c, "验证码验证失败", "/admin/login")
	}
}

func (con LoginController) Captcha(c *gin.Context) {
	id, b64s, err := models.MakeCaptcha()

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"captchaId":    id,
		"captchaImage": b64s,
	})

}

func (con LoginController) LoginOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userInfo")
	err := session.Save()
	if err != nil {
		fmt.Println(err)
	}
	con.success(c, "退出登录成功", "/admin/login")

}
