package types

type ProgramStatus struct {
	Name   string `json:"name"`
	State  string `json:"state"`
	PID    int    `json:"pid"`
	Uptime string `json:"uptime"`
}
