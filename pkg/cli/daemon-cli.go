package cli

import (
	"flag"
	"os"
)

type DaemonOptions struct {
	LogLevel   string
	ConfigPath string
}

func InitDaemon() DaemonOptions {
	fs := flag.NewFlagSet("daemon", flag.ExitOnError)

	logLevel := fs.String("log-level", "INFO", "Log level (DEBUG, INFO, WARN, ERROR)")
	configPath := fs.String("conf", "./config/config.yaml", "Path to config file")

	error := fs.Parse(os.Args[1:])

	if error != nil {
		panic(error)
	}

	return DaemonOptions{
		LogLevel:   *logLevel,
		ConfigPath: *configPath,
	}
}
