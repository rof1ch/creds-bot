package bot

import (
	"context"
	"fmt"
	"passwordbot/pkg/logger/sl"

	tbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (bot *Bot) handleCredByUserId(ctx context.Context, b *tbot.Bot, update *models.Update) {
	keyboardMarkup, err := bot.credsByUserIdMenu(update.CallbackQuery.From.ID)
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

func (bot *Bot) credsByUserIdMenu(userId int64) (*models.InlineKeyboardMarkup, error) {
	var menu models.InlineKeyboardMarkup
	creds, err := bot.services.ByUserId(userId)
	if err != nil {
		return nil, err
	}
	for _, cred := range creds {
		menu.InlineKeyboard = append(menu.InlineKeyboard, []models.InlineKeyboardButton{
			{
				Text:         fmt.Sprintf(cred.Name),
				CallbackData: fmt.Sprintf("user_data_%d", cred.Id),
			},
		})
	}
	menu.InlineKeyboard = append(menu.InlineKeyboard, []models.InlineKeyboardButton{
		{
			Text:         "Добавить данные",
			CallbackData: "cred_add",
		},
	})
	menu.InlineKeyboard = append(menu.InlineKeyboard, []models.InlineKeyboardButton{
		{
			Text:         "<- Назад",
			CallbackData: "default_menu",
		},
	})
	return &menu, nil
}
