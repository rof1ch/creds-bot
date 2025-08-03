package bot

import (
	"context"
	"passwordbot/pkg/logger/sl"

	tbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (bot *Bot) handleListType(ctx context.Context, b *tbot.Bot, update *models.Update) {
	keyboardMarkup, err := bot.listTypes(update.CallbackQuery.From.ID)
	if err != nil {
		bot.log.Error("Ошибка при получении списка", sl.Err(err))
		b.SendMessage(ctx, &tbot.SendMessageParams{
			ChatID: update.CallbackQuery.Message.Message.Chat.ID,
			Text:   "Произошла ощибка при поиске спика, попробуйте еще раз",
		})
	}
	b.EditMessageText(ctx, &tbot.EditMessageTextParams{
		Text:        "Список ваших категорий",
		ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
		MessageID:   update.CallbackQuery.Message.Message.ID,
		ReplyMarkup: keyboardMarkup,
	})
}
