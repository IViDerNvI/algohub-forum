package options

type RedisOptions struct {
	Host     string `json:"host"                     mapstructure:"host"`
	Port     int    `json:"port"                     mapstructure:"port"`
	Username string `json:"username"                 mapstructure:"username"`
	Password string `json:"password"                 mapstructure:"password"`
	Database int    `json:"database"                 mapstructure:"database"`
	Timeout  int    `json:"timeout"                  mapstructure:"timeout"`
}

func NewRedisOptions() *RedisOptions {
	return &RedisOptions{
		Host:     "127.0.0.1",
		Port:     6379,
		Username: "",
		Password: "",
		Database: 0,
		Timeout:  0,
	}
}
