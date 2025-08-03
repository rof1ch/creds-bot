package logger

import (
	"log/slog"
	"os"
	"passwordbot/pkg/logger/handlers/multi"
	"passwordbot/pkg/logger/handlers/slogpretty"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog(slog.LevelDebug)
	case envProd:
		log = setupPrettySlog(slog.LevelInfo)
	}
	return log
}

func setupPrettySlog(level slog.Level) *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: level,
		},
	}

	file, err := os.OpenFile("logs/logfile.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		slog.Error(err.Error())
	}
	fileHandler := slog.NewJSONHandler(file, &slog.HandlerOptions{})

	handler := opts.NewPrettyHandler(os.Stdout)

	multiHandler := multi.NewCopyHandler(fileHandler, handler)

	return slog.New(multiHandler)
}
