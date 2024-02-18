package server

type Config struct {
	Address string
	Port    int
}

func GetDefaultConfig() Config {
	return Config{
		Address: "localhost",
		Port:    8080,
	}
}
