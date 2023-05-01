package admin

import "github.com/gin-gonic/gin"

type BaseController struct{}

func (con BaseController) success(c *gin.Context,message string,redirectUrl string) {
	c.
}

func (con BaseController) error(c *gin.Context) {
	c.String(200, "失败")
}
