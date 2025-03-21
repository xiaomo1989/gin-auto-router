package router

import (
	_ "demo2/app/controller/v1" //一定要导入这个Controller包，用来注册需要访问的方法
	"demo2/router/middleware/jwt"
	"github.com/gin-gonic/gin"
	ginAutoRouter "github.com/xiaomo1989/gin-auto-router"
)

func InitRouter() *gin.Engine {
	//初始化路由
	r := gin.Default()
	//开启v1分组
	v1Route := r.Group("/v1")
	//加载并使用登录验证中间件
	v1Route.Use(jwt.JWT())
	{
		//绑定Group路由，访问路径：/v1/article/list
		ginAutoRouter.BindGroup(v1Route)
	}

	return r
}
