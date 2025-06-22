// Package environ provides facades and interfaces for working with os env
package environ

import "os"

// OSEnv is a facade for os env
type OSEnv struct{}

func NewOSEnv() OSEnv {
	return OSEnv{}
}

func (e OSEnv) Getenv(k string) string {
	return os.Getenv(k)
}
