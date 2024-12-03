package like

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// TODO: add get method
func (lc LikeController) Get(ctx *gin.Context) {
	logrus.WithContext(ctx)
}
