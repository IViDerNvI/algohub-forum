package post

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub-forum/internal/apiserver/pkg/core"
	"github.com/ividernvi/algohub-forum/internal/model"
	"github.com/ividernvi/algohub-forum/pkg/code"
	"github.com/sirupsen/logrus"
)

func (pc PostController) Update(ctx *gin.Context) {
	logrus.WithContext(ctx)

	var p model.Post

	if err := ctx.ShouldBind(&p); err != nil {
		logrus.Errorf("%+v", err)
		core.WriteResponse(ctx, code.ERROR_USER_BIND_FAILED, nil)
		return
	}

	old_post, err := pc.Psrv.Posts().GetById(ctx, ctx.Param("id"), model.GetOptions{})
	if err != nil {
		logrus.Errorf("%+v", err)
		core.WriteResponse(ctx, code.ERROR_USER_NOT_EXISTS, nil)
		return
	}

	old_post.Title = p.Title
	old_post.Content = p.Content

	old_post.ObjectMeta.UpdatedAt = time.Now()

	if err := pc.Psrv.Posts().Update(ctx, *old_post, model.UpdateOptions{}); err != nil {
		logrus.Errorf("save user failed %+v", err)
		core.WriteResponse(ctx, code.ERROR_DATABASE_SAVE_FAILED, nil)
		return
	}

	core.WriteResponse(ctx, 0, old_post)
}
