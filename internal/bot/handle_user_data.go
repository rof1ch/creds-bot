package bot

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	tbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (bot *Bot) handleUserData(ctx context.Context, b *tbot.Bot, update *models.Update) {
	userID := update.CallbackQuery.From.ID
	chatID := update.CallbackQuery.Message.Message.Chat.ID
	session, isOk := bot.sessions.GetSession(userID)
	if !isOk {
		msg, _ := b.SendMessage(ctx, &tbot.SendMessageParams{
			ChatID: chatID,
			Text:   "Введите ключ для разблокировки данных",
		})
		bot.addDeleteMessage(msg.ID, msg.Chat.ID, 15)
		bot.userStates[userID] = &UserState{
			Step: "wait_key",
		}
		return
	}
	if bot.sessions.NeedsReauth(userID) {
		msg, _ := b.SendMessage(ctx, &tbot.SendMessageParams{
			ChatID: chatID,
			Text:   "Введите ключ для разблокировки данных",
		})
		bot.addDeleteMessage(msg.ID, msg.Chat.ID, 15)
		bot.userStates[userID] = &UserState{
			Step: "wait_key",
		}
		return
	}
	dataIdStr := strings.Split(update.CallbackQuery.Data, "user_data_")
	dataId, _ := strconv.Atoi(dataIdStr[1])

	cred, err := bot.services.Credintial.ById(uint(dataId), session.DecryptKey)
	if err != nil {
		msg, _ := b.SendMessage(ctx, &tbot.SendMessageParams{
			ChatID: chatID,
			Text:   "Произошла ошибка, попробуйте позже",
		})
		bot.addDeleteMessage(msg.ID, msg.Chat.ID, 10)
	}

	b.EditMessageText(ctx, &tbot.EditMessageTextParams{
		ChatID:      chatID,
		MessageID:   update.CallbackQuery.Message.Message.ID,
		ReplyMarkup: bot.userDataMenu(userID),
		ParseMode:   models.ParseModeMarkdown,
		Text: fmt.Sprintf(`
Название: %s
Описание: %s
Категория: %s
Логин: %s
Пароль: %s
`, cred.Name, cred.Description, cred.Type.Name, fmt.Sprintf("`%s`", cred.Login), fmt.Sprintf("`%s`", cred.Password)),
	})
}

func (bot *Bot) userDataMenu(userId int64) *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{
					Text:         "Удалить",
					CallbackData: fmt.Sprintf("delete_data_%d", userId),
				},
			},
            {
                {
                    Text: "<- Назад",
                    CallbackData: "data_list",
                },
            },
		},
	}
}