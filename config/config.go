package config

var SecretKey = []byte("super_secret_key")

type ServerConfig struct {
	Addr       string
	DbUser     string
	DbPassword string
	DbHost     string
	DbName     string
}

func LoadServerConfig() (*ServerConfig, error) {
	return &ServerConfig{
		Addr:       ":3000",
		DbUser:     "matcha_latte",
		DbPassword: "matcha@latte",
		DbHost:     "localhost",
		DbName:     "matcha_latte_db",
	}, nil
}
