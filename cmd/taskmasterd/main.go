package main

import (
	"log/slog"

	"github.com/ayb-cha/taskmaster/pkg/cli"
	"github.com/ayb-cha/taskmaster/pkg/config"
	"github.com/ayb-cha/taskmaster/pkg/logger"
)

func main() {
	options := cli.InitDaemon()
	logger.Init(options.LogLevel, nil)
	config := config.ReadConfig(options.ConfigPath)

	slog.Debug("starting taskmasterd", "config", config)
}
