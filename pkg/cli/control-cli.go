package cli

import "flag"

type ControlOptions struct {
	LogLevel   string
	ConfigPath string
}

func InitControl() ControlOptions {
	logLevel := flag.String("log-level", "INFO", "Log level (DEBUG, INFO, WARN, ERROR)")
	flag.Parse()

	return ControlOptions{
		LogLevel: *logLevel,
	}
}
