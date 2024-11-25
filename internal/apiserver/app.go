package apiserver

import (
	"github.com/ividernvi/algohub-forum/internal/model"
	"github.com/ividernvi/algohub-forum/internal/model/options"
	"github.com/spf13/cobra"
)

func ServerRun(cmd *cobra.Command, args []string) {
	opts := options.NewOptions()
	cfg := model.NewConfigFromOptions(opts)

	app := New(cfg)
	app.Prepare().Run()
}

var ServerCmd = &cobra.Command{
	Run: ServerRun,
}
