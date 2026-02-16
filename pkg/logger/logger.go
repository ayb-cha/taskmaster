package logger

import (
	"log/slog"
	"os"
	"strings"
)

var logLevelMap = map[string]slog.Level{
	"DEBUG": slog.LevelDebug,
	"INFO":  slog.LevelInfo,
	"WARN":  slog.LevelWarn,
	"ERROR": slog.LevelError,
}

func Init(level string, logPath *string) {
	slevel := parseLogLevel(level)

	logFile := os.Stdout

	if logPath != nil {
		openedLogFile, err := os.OpenFile(*logPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
		if err != nil {
			panic(err)
		}
		logFile = openedLogFile
	}

	logger := slog.New(
		slog.NewJSONHandler(
			logFile,
			&slog.HandlerOptions{
				AddSource: false,
				Level:     slevel,
			},
		),
	)

	slog.SetDefault(logger)
}

func parseLogLevel(level string) slog.Level {
	normalized := strings.ToUpper(strings.TrimSpace(level))

	if l, ok := logLevelMap[normalized]; ok {
		return l
	}

	return slog.LevelInfo
}
