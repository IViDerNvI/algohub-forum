package server

type SecureServingInfo struct {
	BindAddress string
	BindPort    int
	ServerCert  CertKey
}

type CertKey struct {
	CertFile string
	KeyFile  string
}
