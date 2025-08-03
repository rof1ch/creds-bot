package bot

import (
	"context"
	"passwordbot/pkg/logger/sl"
	"strconv"
	"strings"

	tbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (bot *Bot) handleCredByType(ctx context.Context, b *tbot.Bot, update *models.Update) {
	typeIdStr := strings.Split(update.CallbackQuery.Data, "data_type_")
	typeId, _ := strconv.Atoi(typeIdStr[1])

	keyboardMarkup, err := bot.credsByTypeIdMenu(uint(typeId))
	if err != nil {
		bot.log.Error("Ошибка при получении списка", sl.Err(err))
		b.SendMessage(ctx, &tbot.SendMessageParams{
			ChatID: update.CallbackQuery.Message.Message.Chat.ID,
			Text:   "Произошла ошибка при поиске спика, попробуйте еще раз",
		})
	}
	b.EditMessageText(ctx, &tbot.EditMessageTextParams{
		Text:        "Список ваших данных",
		ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
		MessageID:   update.CallbackQuery.Message.Message.ID,
		ReplyMarkup: keyboardMarkup,
	})
}
