package apiserver

import "github.com/ividernvi/algohub-forum/internal/apiserver"

func Run() {
	apiserver.ServerCmd.Run(apiserver.ServerCmd, []string{})
}
