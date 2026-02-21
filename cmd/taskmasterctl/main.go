package main

import (
	"log/slog"
	"strings"

	"github.com/ayb-cha/taskmaster/internal/control"
	"github.com/ayb-cha/taskmaster/internal/httpclient"
	"github.com/ayb-cha/taskmaster/pkg/cli"
	"github.com/ayb-cha/taskmaster/pkg/config"
	"github.com/ayb-cha/taskmaster/pkg/logger"
	"github.com/chzyer/readline"
)

func main() {
	options := cli.InitControl()
	logger.Init(options.LogLevel, &options.LogFilePath)
	config := config.ReadConfig(options.ConfigPath)
	httpclient := httpclient.NewClient(config)

	httpclient.Ping()

	rl, err := readline.New("> ")
	if err != nil {
		slog.Error("failed to create readline instance", "error", err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}

		control.Resolve(strings.TrimSpace(line), config, httpclient)
		slog.Info("readline input caught", "line", line)
	}

	slog.Debug("starting taskmasterctl", "config", config)
}
