package settings

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type MySQLConfig struct {
	User         string `mapstructure:"user"`
	Pass         string `mapstructure:"pass"`
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Name         string `mapstructure:"name"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxLifetime  int    `mapstructure:"max_lifetime"`
}

type SecurityConfig struct {
	JWTAccessSecret      string `mapstructure:"jwt_access_secret"`
	JWTRefreshSecret     string `mapstructure:"jwt_refresh_secret"`
	JWTAccessExpiration  string `mapstructure:"jwt_access_expiration"`
	JWTRefreshExpiration string `mapstructure:"jwt_refresh_expiration"`
}

type LogConfig struct {
	LogLevel   string `mapstructure:"log_level"`
	FilePath   string `mapstructure:"file_path"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}

type Config struct {
	Server         ServerConfig   `mapstructure:"server"`
	MySQLConfig    MySQLConfig    `mapstructure:"mysql"`
	LogConfig      LogConfig      `mapstructure:"log"`
	SecurityConfig SecurityConfig `mapstructure:"security"`
	RedisConfig    RedisConfig    `mapstructure:"redis"`
}
