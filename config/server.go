package config

import "os"

func GetServerPort() string {
	return os.Getenv("SERVER_PORT")
}
