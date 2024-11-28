package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub-forum/internal/apiserver/pkg/core"
	"github.com/ividernvi/algohub-forum/internal/model"
	"github.com/ividernvi/algohub-forum/pkg/code"
	"github.com/sirupsen/logrus"
)

func (uc *UserController) Logout(ctx *gin.Context) {
	logrus.WithContext(ctx)

	authHeader := ctx.GetHeader("Authorization")
	if authHeader != "" {
		ctx.Header("Authorization", "")
		uc.Usrv.Tokens().Delete(ctx, authHeader, model.DeleteOptions{})
	}

	core.WriteResponse(ctx, code.ERROR_OK, "OK")
}
