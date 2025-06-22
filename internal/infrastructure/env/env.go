package env

import "os"

type OSEnv struct{}

func NewOSEnv() OSEnv {
	return OSEnv{}
}

func (e OSEnv) Getenv(k string) string {
	return os.Getenv(k)
}
