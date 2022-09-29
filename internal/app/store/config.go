package store

type Config struct {
	// строка подключения к БД
	DataBaseURL string `toml:"database_url"`
}

// create new config
func NewConfig() *Config  {
	return &Config{}
}