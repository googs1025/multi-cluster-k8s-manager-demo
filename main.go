package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"mutli-cluster-k8s-manager/pkg/controllers"
	_ "mutli-cluster-k8s-manager/pkg/init_multi_cluster"
	_ "mutli-cluster-k8s-manager/pkg/services"

	"net/http"
)

// gin 原始中间件，支持跨域。
func cross() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if method != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization,X-Token")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

	}
}

func main() {

	// 用普通的gin对接
	//router := gin.New()
	//
	//
	//defer func() {
	//	_ = router.Run(":8080")
	//}()
	//
	//register(router)


	server := goft.Ignite(cross()).Config(

	).
		Mount("", // TODO: 支持其他更多工作负载资源
			controllers.NewDeploymentCtl(),
			controllers.NewPodCtl(),

		).
		Attach(
			//middlewares.NewCrossMiddleware(), //跨域中间件
		)

	// 前端布署静态文件。
	//server.Static("/dashboard", "./admin")
	//server.Static("/static", "./admin/static")

	server.Launch()

}



