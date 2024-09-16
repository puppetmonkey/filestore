package middlewares

import (
	"filestore/controller"
	"filestore/pkg/jwt"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 从请求的query参数中获取token
		// 假设token放在query参数中，使用token字段
		// 例如: /api/resource?token=xxxxx.xxx.xxx
		tokenString := c.Query("token")

		if tokenString == "" {
			// 如果query中没有token，返回需要登录的错误
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}

		// 使用之前定义好的解析JWT的函数来解析token
		mc, err := jwt.ParseToken(tokenString)
		if err != nil {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set(controller.CtxUserIDKey, mc.UserID)
		c.Set(controller.CtxUserNameKey, mc.Username)
		c.Next() // 后续的处理请求的函数中，可以通过c.Get(CtxUserIDKey)来获取当前请求的用户信息
	}
}
