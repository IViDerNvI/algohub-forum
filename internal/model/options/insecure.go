package options

import (
	"fmt"

	"github.com/spf13/pflag"
)

type InsecureServingOptions struct {
	BindAddress string `json:"bind-address" mapstructure:"bind-address"`
	BindPort    int    `json:"bind-port"    mapstructure:"bind-port"`
}

func NewInsecureServingOptions() *InsecureServingOptions {
	return &InsecureServingOptions{
		BindAddress: "127.0.0.1",
		BindPort:    8089,
	}
}

func (s *InsecureServingOptions) Validate() []error {
	var err []error
	if s.BindPort > 65535 || s.BindPort <= 0 {
		err = append(err, fmt.Errorf("port %s invalid", s.BindAddress))
	}

	return err
}

func (s *InsecureServingOptions) AddFlags(fs *pflag.FlagSet) {
	pflag.StringVar(&s.BindAddress, "insecure.bind-addr", s.BindAddress, ""+
		"The IP address on which to serve the --insecure.bind-port(set to 0.0.0.0 for all IPv4 interfaces and :: for all IPv6 interfaces).")
	pflag.IntVar(&s.BindPort, "insecure.bind-port", s.BindPort, ""+
		"The port on which to serve unsecured, unauthenticated insecure access.")
	return
}
