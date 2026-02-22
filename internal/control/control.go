package control

import (
	"fmt"
	"os"

	"github.com/ayb-cha/taskmaster/internal/httpclient"
	"github.com/ayb-cha/taskmaster/pkg/config"
)

func Resolve(command string, config *config.Config, httpclient *httpclient.Client) {
	switch command {
	case "status":
		httpclient.GetStatus()
	default:
		fmt.Fprintf(os.Stderr, "ERROR: Unknown command `%s`\n", command)
	}
}
