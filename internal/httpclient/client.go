package httpclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"os"

	"github.com/ayb-cha/taskmaster/internal/prettyprint"
	"github.com/ayb-cha/taskmaster/pkg/config"
	"github.com/ayb-cha/taskmaster/pkg/types"
)

type Client struct {
	client *http.Client
}

func NewClient(config *config.Config) *Client {
	_, err := os.Stat(config.UnixHttpServer.File)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			slog.Error("socket file dose not exists", "error", err)
			fmt.Fprintln(os.Stderr, "ERROR: unix://"+config.UnixHttpServer.File+" does not exist")
			os.Exit(1)
		}
		fmt.Fprintln(os.Stderr, "ERROR: unix://"+config.UnixHttpServer.File+" "+err.Error())
		os.Exit(1)
	}
	return &Client{
		client: &http.Client{
			Transport: &http.Transport{
				DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
					return net.Dial("unix", config.UnixHttpServer.File)
				},
			},
		},
	}
}

func (c *Client) Ping() {
	resp, err := c.client.Get("http://unix/ping")
	if err != nil {
		slog.Error("failed to send request to unix socket", "error", err)
		fmt.Fprintln(os.Stderr, "ERROR: failed to send request to: "+err.Error())
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("failed to read response body", "error", err)
		return
	}

	slog.Debug("received response from unix socket", "status", resp.Status, "body", string(body))
}

func (c *Client) GetStatus() {
	resp, err := c.client.Get("http://unix/status")
	if err != nil {
		slog.Error("failed to send request to unix socket", "error", err)
		fmt.Fprintln(os.Stderr, "ERROR: failed to send request to unix socket: "+err.Error())
		return
	}

	defer resp.Body.Close()

	slog.Debug("received response from unix socket", "status", resp.Status)

	var statuses []types.ProgramStatus
	err = json.NewDecoder(resp.Body).Decode(&statuses)

	if err != nil {
		slog.Error("error decoding status response json", "error", err)
	}

	for _, progStatus := range statuses {
		prettyprint.PrettyprintProgramStatus(progStatus)
	}
}
