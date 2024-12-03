package user

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub-forum/internal/apiserver/pkg/core"
	"github.com/ividernvi/algohub-forum/internal/model"
	"github.com/ividernvi/algohub-forum/pkg/code"
	"github.com/sirupsen/logrus"
)

func (uc UserController) CheckLogin(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")

	if strings.HasPrefix(authHeader, "Basic ") {
		uc.Login(ctx)
		ctx.Next()
	}

	if strings.HasPrefix(authHeader, "Bearer ") {
		tokenString := authHeader[len("Bearer "):]
		logrus.Printf("bearer token: %s", tokenString)
		token, err := uc.Usrv.Tokens().Get(ctx, tokenString, model.GetOptions{})
		if err != nil {
			core.WriteResponse(ctx, code.ERROR_DATABASE_GET_FAILED, nil)
			ctx.Abort()
			return
		}

		if token == nil {
			ctx.Redirect(304, "/login")
			ctx.Abort()
			return
		}

		ctx.Next()
		return
	}

	ctx.Redirect(304, "/login")
	ctx.Abort()
}
