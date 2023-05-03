package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MainController struct{}

func (con MainController) Index(c *gin.Context) {
	// 获取 userInfo 对应的 session
	// session := sessions.Default(c)
	// userInfo := session.Get("userInfo")
	// c.JSON(http.StatusOK, gin.H{
	// 	"userInfo": userInfo,
	// })
	c.HTML(http.StatusOK, "admin/main/index.html", gin.H{})
}

func (con MainController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}
