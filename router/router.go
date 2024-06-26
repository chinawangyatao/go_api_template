package router

import (
	"github.com/gin-gonic/gin"
	"go_admin/middleware"
	"go_admin/server/captcha"
	"go_admin/server/miniLogin"
	"go_admin/server/miniUpload"
	"go_admin/server/miniUserBanner"
	"go_admin/server/userInfo"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//router.Use(gin.Recovery())
	router.Use(middleware.Cors())
	router.Static("/static", "./static")
	register(router)
	return router
}

// 路由注册
func register(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/api/captcha", captcha.ResCaptcha)
		v1.POST("/api/login", userInfo.Login)
		v1.POST("/api/register", userInfo.Register)
		v1.POST("/api/verifyUsername", userInfo.VerifyUsername)
	}
	v1.Use(middleware.AuthJWT())
	// 以下是权限路由
	{
		v1.GET("/api/test", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": map[string]interface{}{},
			})
		})

	}
	//	todo
	//	小程序接口
	mini := router.Group("/mini")
	{
		mini.POST("/api/login", miniLogin.Login)                 // 登录
		mini.POST("/api/miniUpload", miniUpload.Uploads)         // 文件上传
		mini.POST("/api/getUserInfo", miniLogin.GetMiniUserInfo) // 获取用户信息

		mini.GET("/api/getUserBanner", miniUserBanner.GetBanner)          // curd banner
		mini.POST("/api/setUpDateBanner", miniUserBanner.SetUpDateBanner) // curd banner
		mini.DELETE("/api/deleteBanner", miniUserBanner.DeleteBanner)     // curd banner
	}

}
