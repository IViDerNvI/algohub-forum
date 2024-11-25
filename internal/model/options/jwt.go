package options

import (
	"fmt"
	"time"

	"github.com/spf13/pflag"
)

type JwtOptions struct {
	Realm      string        `json:"realm"       mapstructure:"realm"`
	Key        string        `json:"key"         mapstructure:"key"`
	Timeout    time.Duration `json:"timeout"     mapstructure:"timeout"`
	MaxRefresh time.Duration `json:"max-refresh" mapstructure:"max-refresh"`
}

// NewJwtOptions creates a JwtOptions object with default parameters.
func NewJwtOptions() *JwtOptions {

	return &JwtOptions{
		Realm:      "algohub jwt",
		Key:        "",
		Timeout:    1 * time.Hour,
		MaxRefresh: 1 * time.Hour,
	}
}

func (s *JwtOptions) Validate() []error {
	var errs []error

	if len(s.Key) < 6 || len(s.Key) > 33 {
		errs = append(errs, fmt.Errorf("jwt key %s invalid", s.Key))
	}

	return errs
}

func (s *JwtOptions) AddFlags(fs *pflag.FlagSet) {
	if fs == nil {
		return
	}

	fs.StringVar(&s.Realm, "jwt.realm", s.Realm, "Realm name to display to the user.")
	fs.StringVar(&s.Key, "jwt.key", s.Key, "Private key used to sign jwt token.")
	fs.DurationVar(&s.Timeout, "jwt.timeout", s.Timeout, "JWT token timeout.")

	fs.DurationVar(&s.MaxRefresh, "jwt.max-refresh", s.MaxRefresh, ""+
		"This field allows clients to refresh their token until MaxRefresh has passed.")
}
