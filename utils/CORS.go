// 跨域处理
package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func CORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 合法域名
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*") 
		// 缓存预检请求时间
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")  
		// 合法的方法（get、post、put等）
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*") 
		// 允许请求带的header
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*") 
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true") 

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}