package like

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub-forum/internal/apiserver/pkg/core"
	"github.com/ividernvi/algohub-forum/internal/model"
	"github.com/ividernvi/algohub-forum/pkg/code"
	"github.com/sirupsen/logrus"
)

func (lc LikeController) Create(ctx *gin.Context) {
	logrus.WithContext(ctx)

	var l model.Like

	if err := ctx.ShouldBind(&l); err != nil {
		logrus.Errorf("%+v", err)
		core.WriteResponse(ctx, code.ERROR_USER_BIND_FAILED, nil)
		return
	}

	l.ItemId = ctx.Param("id")
	l.CreatedAt = time.Now()
	l.UpdatedAt = time.Now()

	logrus.Printf("like %+v will be insert", l)

	if err := lc.Lsrv.Likes().Create(ctx, l, model.CreateOptions{}); err != nil {
		logrus.Errorf("create user failed %+v", err)
		core.WriteResponse(ctx, code.ERROR_DATABASE_INSERT_FAILED, nil)
		return
	}

	core.WriteResponse(ctx, code.ERROR_OK, l)
}
