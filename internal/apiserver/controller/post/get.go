package post

import (
	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub-forum/internal/apiserver/pkg/core"
	"github.com/ividernvi/algohub-forum/internal/model"
	"github.com/ividernvi/algohub-forum/pkg/code"
	"github.com/sirupsen/logrus"
)

func (pc PostController) Get(ctx *gin.Context) {
	logrus.WithContext(ctx)

	postId := ctx.Param("id")

	post, err := pc.Psrv.Posts().GetById(ctx, postId, model.GetOptions{})

	if err != nil {
		core.WriteResponse(ctx, code.ERROR_POST_NOT_EXISTS, nil)
		return
	}

	core.WriteResponse(ctx, code.ERROR_OK, post)
	return
}
