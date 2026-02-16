package config

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	UnixHttpServer UnixHttpServer     `yaml:"unix_http_server"`
	Programs       map[string]Program `yaml:"programs"`
}

type UnixHttpServer struct {
	File string `yaml:"file"`
}

type Program struct {
	Cmd          string            `yaml:"cmd"`
	Numprocs     int               `yaml:"numprocs"`
	Umask        int               `yaml:"umask"`
	Workingdir   string            `yaml:"workingdir"`
	Autostart    bool              `yaml:"autostart"`
	Autorestart  string            `yaml:"autorestart"`
	Exitcodes    []int             `yaml:"exitcodes"`
	Startretries int               `yaml:"startretries"`
	Starttime    int               `yaml:"starttime"`
	Stopsignal   string            `yaml:"stopsignal"`
	Stoptime     int               `yaml:"stoptime"`
	Stdout       string            `yaml:"stdout"`
	Stderr       string            `yaml:"stderr"`
	Env          map[string]string `yaml:"env,omitempty"`
}

func ReadConfig(configPath string) *Config {
	var config Config

	data, err := os.ReadFile(configPath)
	if err != nil {
		slog.Error("failed to read config file", "error", err)
		panic(err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		slog.Error("failed to unmarshal config", "error", err)
		panic(err)
	}

	return &config
}
