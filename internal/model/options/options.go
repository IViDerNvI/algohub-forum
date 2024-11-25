package options

type Options struct {
	RESTfulServerRunOptions *ServerRunOptions       `json:"server"   mapstructure:"server"`
	GRPCOptions             *GRPCOptions            `json:"grpc"     mapstructure:"grpc"`
	InsecureServing         *InsecureServingOptions `json:"insecure" mapstructure:"insecure"`
	SecureServing           *SecureServingOptions   `json:"secure"   mapstructure:"secure"`
	MySQLOptions            *MYSQLOptions           `json:"mysql"    mapstructure:"mysql"`
	RedisOptions            *RedisOptions           `json:"redis"    mapstructure:"redis"`
	JwtOptions              *JwtOptions             `json:"jwt"      mapstructure:"jwt"`
}

func NewOptions() *Options {
	opt := &Options{
		RESTfulServerRunOptions: NewServerRunOptions(),
		GRPCOptions:             NewGRPCOptions(),
		InsecureServing:         NewInsecureServingOptions(),
		SecureServing:           NewSecureServingOptions(),
		MySQLOptions:            NewDBOptions(),
		RedisOptions:            NewRedisOptions(),
		JwtOptions:              NewJwtOptions(),
	}

	return opt
}
