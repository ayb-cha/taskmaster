package main

import (
	"log/slog"

	"github.com/ayb-cha/taskmaster/pkg/cli"
	"github.com/ayb-cha/taskmaster/pkg/config"
	"github.com/ayb-cha/taskmaster/pkg/logger"
)

func main() {
	options := cli.InitControl()
	logger.Init(options.LogLevel, &options.LogFilePath)
	config := config.ReadConfig(options.ConfigPath)

	slog.Debug("starting taskmasterctl", "config", config)
}
