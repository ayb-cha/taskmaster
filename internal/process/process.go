package process

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/ayb-cha/taskmaster/pkg/config"
	"golang.org/x/sys/unix"
)

type ProcessState int

const (
	StateStopped ProcessState = iota
	StateRunning
	StateFailed
	StateFatal
)

type Process struct {
	State ProcessState
}

type ProgramAutoRestartValue int

const (
	AutoRestartTrue = iota
	AutoRestartUnexpected
	AutostartFalse
)

type Program struct {
	Name         string
	Numprocs     int
	Umask        int
	Workdir      string
	Autostart    bool
	ExitCodes    []int
	StartRetries int
	StartTime    int
	StopSignal   int
	StopTime     int
	Stdout       string
	Stderr       string
	Env          map[string]string
	Processes    []Process
}

func newProgram(name string, prog config.Program) *Program {
	errors := []string{}

	if prog.Numprocs < 1 {
		slog.Error("numprocs is leass than 1", "numpproc", prog.Numprocs)
		errors = append(errors, "ERROR: `numprocs` must be at least 1")
	}

	if prog.Umask < 0 {
		slog.Error("umak cannot be negative", "umask", prog.Umask)
		errors = append(errors, "ERROR: `umask` cannot be negative")
	}
	if prog.Umask&^0o777 != 0 {
		slog.Error("umask contains invalid bits", "umask", prog.Umask)
		errors = append(errors, "ERROR: umask must only contain permission bits (000-777)")
	}

	if prog.Autorestart != "true" && prog.Autorestart != "false" && prog.Autorestart != "unexpected" {
		slog.Error("invalid value for autorestart", "autorestart", prog.Autorestart)
		errors = append(errors, "ERROR: autorestart can only be `true`, `false` or `unexpected`")
	}

	if unix.SignalNum(prog.Stopsignal) == 0 {
		slog.Error("invalid stop signal", "stopsignal", prog.Stopsignal)
		errors = append(errors, "ERROR: invalid stop signal")
	}

	if prog.Startretries < 0 {
		slog.Error("startretries cannot be negative", "startretries", prog.Startretries)
		errors = append(errors, "ERROR: `startretries` cannot be negative")
	}

	if prog.Starttime < 0 {
		slog.Error("starttime cannot be negative", "starttime", prog.Starttime)
		errors = append(errors, "ERROR: `starttime` cannot be negative")
	}

	if prog.Stoptime < 0 {
		slog.Error("stoptime cannot be negative", "stoptime", prog.Stoptime)
		errors = append(errors, "ERROR: `stoptime` cannot be negative")
	}

	if len(errors) > 0 {
		for _, error := range errors {
			fmt.Fprintln(os.Stderr, error)
		}
		os.Exit(1)
	}

	return &Program{
		Name:         name,
		Workdir:      prog.Workingdir,
		Umask:        prog.Umask,
		Autostart:    prog.Autostart,
		ExitCodes:    prog.Exitcodes,
		StartRetries: prog.Startretries,
		StartTime:    prog.Starttime,
		StopSignal:   int(unix.SignalNum(prog.Stopsignal)),
		StopTime:     prog.Stoptime,
		Stdout:       prog.Stdout,
		Stderr:       prog.Stderr,
		Env:          prog.Env,
	}
}

func Start(config *config.Config) {
	var programs []Program

	for name, prog := range config.Programs {
		programs = append(programs, *newProgram(name, prog))
	}

	fmt.Print(programs)
}
