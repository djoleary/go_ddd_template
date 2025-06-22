package environ

// Getenver is a narrow interface exposing only Getenv
type Getenver interface {
	Getenv(k string) string
}
