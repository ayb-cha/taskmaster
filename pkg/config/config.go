package config

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	UnixHttpServer UnixHttpServer     `yaml:"unix_http_server,omitempty"`
	Programs       map[string]Program `yaml:"programs"`
}

type UnixHttpServer struct {
	File string `yaml:"file"`
}

type Program struct {
	Cmd          string            `yaml:"cmd,omitempty"`
	Numprocs     int               `yaml:"numprocs,omitempty"`
	Umask        int               `yaml:"umask,omitempty"`
	Workingdir   string            `yaml:"workingdir,omitempty"`
	Autostart    bool              `yaml:"autostart,omitempty"`
	Autorestart  string            `yaml:"autorestart,omitempty"`
	Exitcodes    []int             `yaml:"exitcodes,omitempty"`
	Startretries int               `yaml:"startretries,omitempty"`
	Starttime    int               `yaml:"starttime,omitempty"`
	Stopsignal   string            `yaml:"stopsignal,omitempty"`
	Stoptime     int               `yaml:"stoptime,omitempty"`
	Stdout       string            `yaml:"stdout,omitempty"`
	Stderr       string            `yaml:"stderr,omitempty"`
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
