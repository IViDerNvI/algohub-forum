package options

type SecureServingOptions struct {
	BindAddress string  `json:"bind-address" mapstructure:"bind-address"`
	BindPort    int     `json:"bind-port"    mapstructure:"bind-port"`
	ServerCert  CertKey `json:"ssl"          mapstructure:"ssl"`
}

type CertKey struct {
	CertFile string `json:"cert-file"        mapstructure:"cert-file"`
	KeyFile  string `json:"private-key-file" mapstructure:"private-key-file"`
}

func NewSecureServingOptions() *SecureServingOptions {
	return &SecureServingOptions{
		BindAddress: "127.0.0.1",
		BindPort:    8090,
		ServerCert: CertKey{
			CertFile: "./cert/cert.pem",
			KeyFile:  "./cert/key.pem",
		},
	}
}
