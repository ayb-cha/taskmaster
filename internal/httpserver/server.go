package httpserver

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"

	"github.com/ayb-cha/taskmaster/pkg/config"
	"github.com/ayb-cha/taskmaster/pkg/types"
)

func registerRoutes() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("new client connected")
		fmt.Fprintln(w, "pong")
	})

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("status requested")

		nginxDummyStatus := types.ProgramStatus{
			Name:   "nginx",
			State:  "RUNNING",
			PID:    553,
			Uptime: "1:05:21",
		}

		myAppDummyStatus := types.ProgramStatus{
			Name:   "my-app",
			State:  "EXITED",
			PID:    -1,
			Uptime: "Exited too quickly (process log may have details)",
		}

		data := []types.ProgramStatus{
			nginxDummyStatus,
			myAppDummyStatus,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func Init(config *config.Config) (net.Listener, error) {
	error := os.Remove(config.UnixHttpServer.File)
	if error != nil && !os.IsNotExist(error) {
		slog.Error("failed to remove existing unix socket file", "error", error)
		return nil, error
	}

	listener, err := net.Listen("unix", config.UnixHttpServer.File)
	if err != nil {
		slog.Error("failed to create unix socket", "error", err)
		return nil, err
	}

	registerRoutes()

	// go func() {
	if err := http.Serve(listener, nil); err != nil {
		slog.Error("failed to Listen on unix socket", "error", err)
	}

	slog.Debug("start listening for http requests")
	// }()

	return listener, nil
}

func Stop(listener net.Listener) error {
	if err := listener.Close(); err != nil {
		slog.Error("failed to close unix socket listener", "error", err)
		return err
	}

	slog.Debug("stopped listening for http requests")
	return nil
}
