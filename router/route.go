package router

import (
	"filestore/controller"
	"filestore/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter 路由
func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	//r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.RateLimitMiddleware(2*time.Second, 1))
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.LoadHTMLGlob("static/view/*")
	r.Static("/static", "./static")

	// // r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/")

	// // 注册
	// v1.POST("/signup", controller.SignUpHandler)
	// // 登录
	// v1.POST("/login", controller.LoginHandler)
	v1.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	v1.GET("/file/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	v1.POST("/file/upload", controller.UploadHandler)
	v1.GET("/file/meta", controller.GetFileMetaHandler)
	v1.POST("/file/download", controller.DownloadHandler)
	v1.POST("/file/update", controller.UpdateFileMeta)
	v1.POST("/file/delete", controller.FileDeleteHandler)
	// // 根据时间或分数获取帖子列表
	// v1.GET("/posts2", controller.GetPostListHandler2)
	// v1.GET("/posts", controller.GetPostListHandler)
	// v1.GET("/community", controller.CommunityHandler)
	// v1.GET("/community/:id", controller.CommunityDetailHandler)
	// v1.GET("/post/:id", controller.GetPostDetailHandler)

	// v1.Use(middlewares.JWTAuthMiddleware()) // 应用JWT认证中间件

	// {
	// 	v1.POST("/post", controller.CreatePostHandler)

	// 	// 投票
	// 	v1.POST("/vote", controller.PostVoteController)
	// }

	// // pprof.Register(r) // 注册pprof相关路由

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
