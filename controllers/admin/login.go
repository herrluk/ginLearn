package admin

import (
	"fmt"
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
	captchaId := c.PostForm("captchaId") // hidden的 id
	verifyValue := c.PostForm("verify")  // 前端验证码输入框的 id

	if flag := models.VerifyCaptcha(captchaId, verifyValue); flag {
		// c.String(http.StatusOK, "验证码验证成功")
		con.success(c, "验证码验证成功", "/admin")
	} else {
		c.String(http.StatusOK, "眼曾妈验证失败")
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
