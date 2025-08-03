package bot

import (
	"context"

	tbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func handleDefault(ctx context.Context, b *tbot.Bot, update *models.Update) {
	b.SendMessage(ctx,
		&tbot.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "Добро пожаловать!",
			ReplyMarkup: defaultMenu,
		},
	)
}

func (bot *Bot) handleDefaultMenu(ctx context.Context, b *tbot.Bot, update *models.Update) {
	b.EditMessageText(ctx, &tbot.EditMessageTextParams{
		ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
		MessageID:   update.CallbackQuery.Message.Message.ID,
		ReplyMarkup: defaultMenu,
		Text:        "Добро пожаловать!",
	})
}
