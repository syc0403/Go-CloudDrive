package route

/**
* @Description
* @Author YuCheng
* @Date 2023/1/18 23:06
**/
import (
	"Go-CloudDrive/handler"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	// gin framework 包括Logger，Recovery
	router := gin.Default()

	// 处理静态资源
	router.Static("/static/", "./static")

	// 不需要经过验证就能访问的接口
	router.GET("/user/signup", handler.SignupHandler)
	router.POST("/user/signup", handler.DoSignupHandler)
	router.GET("/user/signin", handler.SignInHandler)
	router.POST("/user/signin", handler.DoSigninHandler)

	//加入中间件，用于校验token的拦截器
	router.Use(handler.DoHTTPInterceptor())

	//User之后的所有handler都会经过拦截器
	return router
}
