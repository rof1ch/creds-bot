package bot

import (
	"context"
	"strconv"
	"strings"

	tbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (bot *Bot) handleDeleteData(ctx context.Context, b *tbot.Bot, update *models.Update) {
	dataIdStr := strings.Split(update.CallbackQuery.Data, "delete_data_")
	dataId, _ := strconv.Atoi(dataIdStr[1])

    bot.services.Credintial.Delete(uint(dataId))
    bot.handleUserData(ctx, b, update)
}
