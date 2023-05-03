package middlewares

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/herrluk/goProjectsPractice/ginLearn/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(c *gin.Context) {
	// 判断用户是否登录

	fmt.Println(time.Now())

	fmt.Println(c.Request.URL)

	c.Set("username", "张三")

	// 定义一个goroutine统计日志  当在中间件或 handler 中启动新的 goroutine 时，不能使用原始的上下文（c *gin.Context）， 必须使用其只读副本（c.Copy()）
	cCp := c.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Done! in path " + cCp.Request.URL.Path)
	}()
}

func InitAdminAuthMiddleware(c *gin.Context) {
	fmt.Println("InitAdminAuthMiddleware")
	// 进行权限判断 没有登录的用户 不能进入后台管理中心

	// 1、获取Url访问的地址  /admin/captcha

	// 2、获取Session里面保存的用户信息

	// 3、判断Session中的用户信息是否存在，如果不存在跳转到登录页面（注意需要判断） 如果存在继续向下执行

	// 4、如果Session不存在，判断当前访问的URl是否是login doLogin captcha，如果不是跳转到登录页面，如果是不行任何操作

	//  1、获取Url访问的地址   /admin/captcha?t=0.8706946438889653
	// 获取 URL 路径去掉 Get 传值
	pathName := strings.Split(c.Request.URL.String(), "?")[0]
	fmt.Println("pathName:", pathName)
	// 2.获取 session 里面保存的用户信息
	session := sessions.Default(c)
	userInfo := session.Get("userInfo")
	fmt.Println("userInfo:", userInfo)
	// 类型断言 判断 userInfo 是不是一个 string
	userInfoStr, ok := userInfo.(string)
	fmt.Println("Ok:", ok)
	if ok {
		var userinfoStruct []models.Manager
		err := json.Unmarshal([]byte(userInfoStr), &userinfoStruct)

		if err != nil || !(len(userinfoStruct) > 0 && userinfoStruct[0].Username != "") {
			if pathName != "/admin/login" && pathName != "/admin/doLogin" && pathName != "/admin/captcha" {
				c.Redirect(302, "/admin/login")
			}
		}
	} else { // 没有登录
		if pathName != "/admin/login" && pathName != "/admin/doLogin" && pathName != "/admin/captcha" {
			c.Redirect(302, "/admin/login")
		}
	}
}
