package prettyprint

import (
	"fmt"
	"strconv"

	"github.com/ayb-cha/taskmaster/pkg/types"
)

func PrettyprintProgramStatus(status types.ProgramStatus) {
	fmt.Printf("%-30s", status.Name)
	fmt.Printf("%-10s", status.State)
	if status.PID != -1 {
		fmt.Printf("%-10s", "pid("+strconv.FormatInt(int64(status.PID), 10)+")")
	} else {
		fmt.Printf("%-10s", "no pid")
	}

	fmt.Printf("%-10s", status.Uptime)
	fmt.Println()
}
