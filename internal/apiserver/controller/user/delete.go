package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub-forum/internal/apiserver/pkg/core"
	"github.com/ividernvi/algohub-forum/internal/model"
	"github.com/ividernvi/algohub-forum/pkg/code"
	"github.com/sirupsen/logrus"
)

func (uc UserController) Delete(ctx *gin.Context) {
	logrus.WithContext(ctx)
	username := ctx.Param("name")
	if username == "" {
		core.WriteResponse(ctx, code.ERROR_USER_NOT_EXISTS, nil)
		return
	}

	uc.Usrv.Users().DeleteByUsername(ctx, username, model.DeleteOptions{})
	core.WriteResponse(ctx, code.ERROR_OK, "OK")
	return
}
