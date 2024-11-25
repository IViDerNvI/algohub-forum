package options

import (
	"fmt"

	"github.com/spf13/pflag"
)

type GRPCOptions struct {
	BindAddress string `json:"bind-address" mapstructure:"bind-address"`
	BindPort    int    `json:"bind-port"    mapstructure:"bind-port"`
}

func NewGRPCOptions() *GRPCOptions {
	return &GRPCOptions{
		BindAddress: "0.0.0.0",
		BindPort:    8081,
	}
}

func (s *GRPCOptions) Validate() []error {
	var err []error
	if s.BindPort > 65535 || s.BindPort <= 0 {
		err = append(err, fmt.Errorf("port %s invalid", s.BindAddress))
	}

	return err
}

func (s *GRPCOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.BindAddress, "grpc.bind-address", s.BindAddress, ""+
		"The IP address on which to serve the --grpc.bind-port(set to 0.0.0.0 for all IPv4 interfaces and :: for all IPv6 interfaces).")

	fs.IntVar(&s.BindPort, "grpc.bind-port", s.BindPort, ""+
		"The port on which to serve unsecured, unauthenticated grpc access. It is assumed ")
}
