package bot

import (
	"context"
	"log/slog"
	"passwordbot/internal/services"
	"passwordbot/internal/storage/session"
	"time"

	tbot "github.com/go-telegram/bot"

	"passwordbot/internal/domain/dto"
)

type Bot struct {
	log            *slog.Logger
	ctx            context.Context
	services       *services.Services
	b              *tbot.Bot
	userStates     map[int64]*UserState
	deleteMessages []DeleteMessage
	sessions       *session.ListSession
}

type UserState struct {
	Step     string
	TempType *dto.TypeInput
	TempCred *dto.CredintialInput
}

type DeleteMessage struct {
	MessageID int
	ChatID    int64
	ExpiredAt time.Time
}

func New(
	ctx context.Context,
	log *slog.Logger,
	services *services.Services,
	botToken string,
) (*Bot, error) {
	opts := []tbot.Option{
		tbot.WithDefaultHandler(handleDefault),
	}

	b, err := tbot.New(botToken, opts...)
	if err != nil {
		return nil, err
	}

	b.RegisterHandler(tbot.HandlerTypeMessageText, "/start", tbot.MatchTypeExact, handleDefault)

	return &Bot{
		b:          b,
		log:        log,
		services:   services,
		ctx:        ctx,
		userStates: make(map[int64]*UserState),
        sessions: session.NewList(),
	}, nil
}

func (bot *Bot) Run() {
	// Список категорий
	bot.b.RegisterHandler(
		tbot.HandlerTypeCallbackQueryData,
		"type_list",
		tbot.MatchTypeExact,
		bot.handleListType,
	)

	// Добавить категорию
	bot.b.RegisterHandler(
		tbot.HandlerTypeCallbackQueryData,
		"type_add",
		tbot.MatchTypeExact,
		bot.handleAddType,
	)

	// Назад стандартное меню
	bot.b.RegisterHandler(
		tbot.HandlerTypeCallbackQueryData,
		"default_menu",
		tbot.MatchTypeExact,
		bot.handleDefaultMenu,
	)

	// Вывод типа
	bot.b.RegisterHandler(
		tbot.HandlerTypeCallbackQueryData,
		"user_type_",
		tbot.MatchTypePrefix,
		bot.handleTypeById,
	)

	// Удаление типа
	bot.b.RegisterHandler(
		tbot.HandlerTypeCallbackQueryData,
		"type_delete_",
		tbot.MatchTypePrefix,
		bot.handleDeleteType,
	)

	// Поиск данных по категории
	bot.b.RegisterHandler(
		tbot.HandlerTypeCallbackQueryData,
		"data_type_",
		tbot.MatchTypePrefix,
		bot.handleCredByType,
	)

	// Добавление данных
	bot.b.RegisterHandler(
		tbot.HandlerTypeCallbackQueryData,
		"cred_add",
		tbot.MatchTypeExact,
		bot.handleAddCreds,
	)
	
    // Получение данных
	bot.b.RegisterHandler(
		tbot.HandlerTypeCallbackQueryData,
		"data_list",
		tbot.MatchTypeExact,
		bot.handleCredByUserId,
	)

	// Выбор типа для данных
	bot.b.RegisterHandler(
		tbot.HandlerTypeCallbackQueryData,
		"cred_type_",
		tbot.MatchTypePrefix,
		bot.handleSelectType,
	)

	// Вывод данных
	bot.b.RegisterHandler(
		tbot.HandlerTypeCallbackQueryData,
		"user_data_",
		tbot.MatchTypePrefix,
		bot.handleUserData,
	)
	
    // Удаление данных
	bot.b.RegisterHandler(
		tbot.HandlerTypeCallbackQueryData,
		"delete_data_",
		tbot.MatchTypePrefix,
		bot.handleDeleteData,
	)

	// Обработка сообщений для создания
	bot.b.RegisterHandler(
		tbot.HandlerTypeMessageText,
		"",
		tbot.MatchTypePrefix,
		bot.handleMessage,
	)

	bot.b.Start(bot.ctx)
}

func (bot *Bot) DeleteMessages(ctx context.Context) {
	const clearInterval = time.Second * 5
	for {
		select {
		case <-ctx.Done():
			return
		default:
			for _, msg := range bot.deleteMessages {
				if time.Now().After(msg.ExpiredAt) {
					bot.b.DeleteMessage(bot.ctx, &tbot.DeleteMessageParams{
						MessageID: msg.MessageID,
						ChatID:    msg.ChatID,
					})
				}
			}
		}
		time.Sleep(clearInterval)
	}

}
