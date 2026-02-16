package main

import (
	"log/slog"

	"github.com/ayb-cha/taskmaster/pkg/cli"
	"github.com/ayb-cha/taskmaster/pkg/logger"
)

func main() {
	options := cli.InitControl()
	logger.Init(options.LogLevel, &options.LogFilePath)

	slog.Debug("starting taskmasterctl")
}
