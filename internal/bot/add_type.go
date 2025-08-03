package bot

import (
	"context"

	tbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (bot *Bot) handleAddType(ctx context.Context, b *tbot.Bot, update *models.Update) {
	userID := update.CallbackQuery.From.ID

	// Сохраняем состояние для пользователя
	bot.userStates[userID] = &UserState{
		Step: "waiting_for_type_name",
	}

	msg, _ := b.SendMessage(ctx, &tbot.SendMessageParams{
		ChatID: update.CallbackQuery.Message.Message.Chat.ID,
		Text:   "Введите наименование категории:",
	})
	bot.addDeleteMessage(msg.ID, msg.Chat.ID, 15)
}
