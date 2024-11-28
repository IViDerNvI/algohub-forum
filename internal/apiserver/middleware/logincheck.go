package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func LoginCheck() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if strings.HasPrefix(authHeader, "Basic ") {
			// TODO: Basic 验证

		}

		if strings.HasPrefix(authHeader, "Bearer ") {
			// TODO: Bearer 验证
		}

		// TODO: 重定向
	}
}

func BasicLoginCheck(ctx *gin.Context) {

}

func BearerLoginCheck(ctx *gin.Context) {

}
