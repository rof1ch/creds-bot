package bot

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	tbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (bot *Bot) handleTypeById(ctx context.Context, b *tbot.Bot, update *models.Update) {
	typeIdStr := strings.Split(update.CallbackQuery.Data, "user_type_")
	typeId, _ := strconv.Atoi(typeIdStr[1])
	typeCred, _ := bot.services.TypeCred.ById(uint(typeId))
	b.EditMessageText(ctx, &tbot.EditMessageTextParams{
		ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
		MessageID: update.CallbackQuery.Message.Message.ID,
		Text: fmt.Sprintf(`
Наименование: _%s_
Иконка: %s`, typeCred.Name, typeCred.Icon),
		ParseMode: models.ParseModeMarkdown,
        ReplyMarkup: typeByIdMenu(uint(typeId)),
	})
}
