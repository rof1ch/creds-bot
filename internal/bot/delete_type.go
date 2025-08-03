package bot

import (
	"context"
	"strconv"
	"strings"

	tbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (bot *Bot) handleDeleteType(ctx context.Context, b *tbot.Bot, update *models.Update) {
	typeIdStr := strings.Split(update.CallbackQuery.Data, "type_delete_")
	typeId, _ := strconv.Atoi(typeIdStr[1])

	err := bot.services.TypeCred.Delete(uint(typeId))

    var msg *models.Message
	if err == nil {
		msg, _ = b.SendMessage(ctx, &tbot.SendMessageParams{
            ChatID: update.CallbackQuery.Message.Message.Chat.ID,
			Text: "Категория успешно удалена",
		})
		bot.handleListType(ctx, b, update)
	} else {
		msg, _ = b.SendMessage(ctx, &tbot.SendMessageParams{
			Text: "Произошла ошибка, попробуйте позже",
		})
	}
    
    bot.addDeleteMessage(msg.ID, msg.Chat.ID, 10)
}
