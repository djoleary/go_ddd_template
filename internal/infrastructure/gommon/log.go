package gommon

import (
	"github.com/labstack/gommon/log"
)

// GetLevel converts an lowercase english string info a log level
func GetLevel(l string) log.Lvl {
	switch l {
	case "debug":
		return log.DEBUG
	case "info":
		return log.INFO
	case "warn":
		return log.WARN
	case "error":
		return log.ERROR
	}

	return log.INFO
}
