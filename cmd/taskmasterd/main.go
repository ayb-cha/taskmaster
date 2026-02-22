package main

import (
	"github.com/ayb-cha/taskmaster/internal/httpserver"
	"github.com/ayb-cha/taskmaster/internal/process"
	"github.com/ayb-cha/taskmaster/pkg/cli"
	"github.com/ayb-cha/taskmaster/pkg/config"
	"github.com/ayb-cha/taskmaster/pkg/logger"
)

func main() {
	options := cli.InitDaemon()
	logger.Init(options.LogLevel, nil)
	config := config.ReadConfig(options.ConfigPath)
	listener, err := httpserver.Init(config)
	if err != nil {
		panic(err)
	}

	process.Start(config)

	err = httpserver.Stop(listener)
	if err != nil {
		panic(err)
	}
}
