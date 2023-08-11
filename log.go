package socks5

import (
	"os"
)

var doInfo bool

func init() {
	switch os.Getenv("LOG_LEVEL") {
	case "INFO", "DEBUG", "TRACE":
		doInfo = true
	}
}
