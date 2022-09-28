package apiserver

type Config struct {
	BindAddress string `toml:"bind_address"`
	LogLevel string `toml:"log_level"`
}

func NewConfig() *Config  {
	return &Config{
		BindAddress: ":8080",
		LogLevel: "debug",
	}
}