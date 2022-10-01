package apiserver

// import "go_restapi/internal/app/store/sqlstore"

type Config struct {
	BindAddress string `toml:"bind_address"`
	LogLevel    string `toml:"log_level"`
	databaseURL	string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		BindAddress: ":8080",
		LogLevel:    "debug",
	}
}