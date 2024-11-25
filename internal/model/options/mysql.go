package options

type MYSQLOptions struct {
	Host     string `json:"host"       mapstructure:"host"`
	Username string `json:"username"   mapstructure:"username"`
	Password string `json:"password"   mapstructure:"password"`
	Database string `json:"database"   mapstructure:"database"`
}

func NewDBOptions() *MYSQLOptions {
	return &MYSQLOptions{
		Host:     "127.0.0.1",
		Username: "root",
		Password: "root",
		Database: "algohub",
	}
}
