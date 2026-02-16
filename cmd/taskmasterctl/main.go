package main

import (
	"log/slog"

	"github.com/ayb-cha/taskmaster/pkg/cli"
	"github.com/ayb-cha/taskmaster/pkg/logger"
)

func main() {
	logFile := "./logs/taskmasterctl.log"
	options := cli.InitControl()
	logger.Init(options.LogLevel, &logFile)

	slog.Debug("starting taskmasterctl")
}
