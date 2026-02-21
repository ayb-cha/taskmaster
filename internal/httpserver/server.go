package httpserver

import (
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"

	"github.com/ayb-cha/taskmaster/pkg/config"
)

func registerRoutes() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "Hello, Unix Socket!")
	// })
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
