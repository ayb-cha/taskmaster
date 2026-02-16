package cli

import "flag"

type ControlOptions struct {
	LogLevel    string
	LogFilePath string
}

func InitControl() ControlOptions {
	logLevel := flag.String("log-level", "INFO", "Log level (DEBUG, INFO, WARN, ERROR)")
	logFilePath := flag.String("log-file", "./logs/taskmasterctl.log", "Path to log file (optional)")
	flag.Parse()

	return ControlOptions{
		LogLevel:    *logLevel,
		LogFilePath: *logFilePath,
	}
}
