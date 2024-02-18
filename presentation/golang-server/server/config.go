package server

type Config struct {
	Address, Port string
}

func GetDefaultConfig() Config {
	return Config{
		Address: "localhost",
		Port:    "8080",
	}
}
