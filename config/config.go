package config

import (
	"code.google.com/p/gcfg"
	alog "github.com/cenkalti/log"
	"strings"
)

// Config structure that holds Config data
type Config struct {
	// FILES struct {
	// 	TEMPLATES string
	// }

	DB struct {
		HOST  string
		NAME  string
		USER  string
		PASS  string
		DEBUG bool
	}
	HTTP struct {
		HOSTNAME      string
		LISTENADDRESS string
	}
}

// LoadConfig fills Config struct with file data
func LoadConfig(log alog.Logger, config *Config) {
	if err := readConfig("./conf/app.config.ini", config); err != nil {
		log.Fatal("Cannot read config file: ", err)
	}
	// log.Info("Loaded config.ini")
	// log.Debug("%#v\n\n", &config)
}
func readConfig(file string, config *Config) error {
	return gcfg.ReadFileInto(config, file)
}
