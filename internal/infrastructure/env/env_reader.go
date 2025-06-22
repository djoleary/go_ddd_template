package env

type EnvReaderInterface interface {
	Getenv(k string) string
}
