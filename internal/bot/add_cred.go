package bot

import (
	"context"
	"strconv"
	"strings"

	tbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (bot *Bot) handleAddCreds(ctx context.Context, b *tbot.Bot, update *models.Update) {
	userID := update.CallbackQuery.From.ID

	// Сохраняем состояние для пользователя
	bot.userStates[userID] = &UserState{
		Step: "waiting_for_cred_name",
	}

	msg, _ := b.SendMessage(ctx, &tbot.SendMessageParams{
		ChatID: update.CallbackQuery.Message.Message.Chat.ID,
		Text:   "Введите наименование данных:",
	})
	bot.addDeleteMessage(msg.ID, msg.Chat.ID, 15)
}

func (bot *Bot) handleSelectType(ctx context.Context, b *tbot.Bot, update *models.Update) {
	userID := update.CallbackQuery.From.ID
	state := bot.userStates[userID]

	typeIdStr := strings.Split(update.CallbackQuery.Data, "cred_type_")
	typeId, _ := strconv.Atoi(typeIdStr[1])

	state.TempCred.UserId = userID
	state.TempCred.TypeId = uint(typeId)

	bot.services.Credintial.Create(*state.TempCred)

	msg, _ := b.SendMessage(ctx, &tbot.SendMessageParams{
		ChatID: update.CallbackQuery.Message.Message.Chat.ID,
		Text:   "Данные успешно добавленны",
	})
	delete(bot.userStates, userID)
	bot.addDeleteMessage(msg.ID, msg.Chat.ID, 15)
}
