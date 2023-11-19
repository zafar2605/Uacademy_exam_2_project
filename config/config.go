package config

type Config struct {
	ServerHost string
	Port       string

	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
	PostgresPort     string
	PostgresHost     string
}

func Load() *Config {
	var cfg = &Config{}

	cfg.ServerHost = "localhost"
	cfg.Port = ":8080"

	cfg.PostgresHost = "localhost"
	cfg.PostgresUser = "zafar"
	cfg.PostgresPassword = "2605"
	cfg.PostgresDatabase = "exam"
	cfg.PostgresPort = "5432"

	return cfg
}
