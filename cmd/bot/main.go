package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"passwordbot/internal/bot"
	"passwordbot/internal/config"
	"passwordbot/internal/services"
	"passwordbot/internal/storage/db"
	"passwordbot/internal/storage/db/sqlite"
	"passwordbot/pkg/logger"
	"passwordbot/pkg/logger/sl"
	"syscall"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("Ошибка при чтении конфига", sl.Err(err))
		os.Exit(1)
	}

	log := logger.SetupLogger(cfg.ENV)

	sqliteRepo, err := sqlite.NewSqlite(cfg.DBPath)
	if err != nil {
		log.Error("Ошибка при инициализации подключения к SQLite", sl.Err(err))
		os.Exit(1)
	}
	repo := db.New(sqliteRepo, log)
	serv := services.New(log, repo)

	ctx := context.Background()
	botCtx, botCancel := context.WithCancel(ctx)

	b, err := bot.New(botCtx, log, serv, cfg.BotToken)
	if err != nil {
		log.Error("Ошибка при создании обьекта Bot")
		os.Exit(1)
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		log.Info("Завершение работы приложения")
		botCancel()
		os.Exit(0)
	}()

	go b.DeleteMessages(botCtx)

	log.Info("Запуск бота")
	b.Run()
}
