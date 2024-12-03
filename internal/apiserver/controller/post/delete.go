package post

import (
	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub-forum/internal/apiserver/pkg/core"
	"github.com/ividernvi/algohub-forum/internal/model"
	"github.com/ividernvi/algohub-forum/pkg/code"
	"github.com/sirupsen/logrus"
)

func (pc PostController) Delete(ctx *gin.Context) {
	logrus.WithContext(ctx)

	post := ctx.Param("id")

	err := pc.Psrv.Posts().DeleteById(ctx, post, model.DeleteOptions{})
	if err != nil {
		core.WriteResponse(ctx, code.ERROR_DATABASE_DELETE_FAILED, nil)
		return
	}

	core.WriteResponse(ctx, code.ERROR_OK, "OK")
	return
}
