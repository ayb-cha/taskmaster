package cli

import "flag"

type ControlOptions struct {
	LogLevel    string
	LogFilePath string
	ConfigPath  string
}

func InitControl() ControlOptions {
	logLevel := flag.String("log-level", "INFO", "Log level (DEBUG, INFO, WARN, ERROR)")
	logFilePath := flag.String("log-file", "./logs/taskmasterctl.log", "Path to log file (optional)")
	configPath := flag.String("conf", "./config/taskmaster.yaml", "Path to config file")

	flag.Parse()

	return ControlOptions{
		LogLevel:    *logLevel,
		LogFilePath: *logFilePath,
		ConfigPath:  *configPath,
	}
}
