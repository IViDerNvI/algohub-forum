package post

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub-forum/internal/apiserver/pkg/core"
	"github.com/ividernvi/algohub-forum/internal/model"
	"github.com/ividernvi/algohub-forum/pkg/code"
	"github.com/ividernvi/algohub-forum/pkg/util/idutil"
	"github.com/sirupsen/logrus"
)

func (pc *PostController) Create(ctx *gin.Context) {
	logrus.WithContext(ctx)

	var p model.Post

	if err := ctx.ShouldBind(&p); err != nil {
		logrus.Errorf("%+v", err)
		core.WriteResponse(ctx, code.ERROR_USER_BIND_FAILED, nil)
		return
	}

	p.ObjectMeta.InstanceID = idutil.GenerateUUID(p.GetDatabaseName())
	p.ObjectMeta.Name = p.Title
	p.ObjectMeta.CreatedAt = time.Now()
	p.ObjectMeta.UpdatedAt = time.Now()

	logrus.Printf("post %+v will be created", p)

	if err := pc.Psrv.Posts().Create(ctx, p, model.CreateOptions{}); err != nil {
		logrus.Errorf("create post failed: %+v", err)
		core.WriteResponse(ctx, code.ERROR_DATABASE_INSERT_FAILED, nil)
		return
	}

	core.WriteResponse(ctx, code.ERROR_OK, p)
}
