package bot

import (
	"context"
	"passwordbot/internal/domain/dto"

	tbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (bot *Bot) handleMessage(ctx context.Context, b *tbot.Bot, update *models.Update) {
	userID := update.Message.From.ID
	state, ok := bot.userStates[userID]
	if !ok {
		// Нет активного состояния — обычный обработчик
		b.SendMessage(ctx, &tbot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Команда не распознана. Используйте /start",
		})
		return
	}

	var isDelete bool

	switch state.Step {

	// Добавление категории. 1 этап. Ожидание названия
	case "waiting_for_type_name":
		tempType := dto.TypeInput{
			Name: update.Message.Text,
		}
		state.TempType = &tempType
		state.Step = "waiting_for_type_icon"
		msg, _ := b.SendMessage(ctx, &tbot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Теперь отправьте иконку:",
		})
		bot.addDeleteMessage(msg.ID, msg.Chat.ID, 20)

	// Добавление категории. 2 этап. Ожидание иконки
	case "waiting_for_type_icon":
		err := bot.services.TypeCred.Create(state.TempType.Name, update.Message.Text, userID)
		if err != nil {
			msg, _ := b.SendMessage(ctx, &tbot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "❌ Произошла ошибка при добавлении, попробуйте еще раз",
			})
			bot.addDeleteMessage(msg.ID, msg.Chat.ID, 10)
		} else {
			msg, _ := b.SendMessage(ctx, &tbot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "✅ Категория успешно добавлена",
			})
			bot.addDeleteMessage(msg.ID, msg.Chat.ID, 5)
			state.Step = ""
			delete(bot.userStates, userID)
			bot.handleListType(ctx, b, update)
		}

	// Добавление данных. 1 этап. Считывание названия
	case "waiting_for_cred_name":
		tempCred := dto.CredintialInput{
			Name: update.Message.Text,
		}
		state.TempCred = &tempCred
		state.Step = "waiting_for_cred_desc"
		msg, _ := b.SendMessage(ctx, &tbot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Теперь отправьте описание:",
		})
		bot.addDeleteMessage(msg.ID, msg.Chat.ID, 20)

		// 2 этап. Считывание описания
	case "waiting_for_cred_desc":
		state.TempCred.Description = update.Message.Text
		state.Step = "waiting_for_cred_login"
		msg, _ := b.SendMessage(ctx, &tbot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Теперь отправьте логин от аккаунта:",
		})
		bot.addDeleteMessage(msg.ID, msg.Chat.ID, 20)

	case "waiting_for_cred_login":
		state.TempCred.Login = update.Message.Text
		state.Step = "waiting_for_cred_pass"
		msg, _ := b.SendMessage(ctx, &tbot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Теперь отправьте пароль от аккаунта:",
		})
		bot.addDeleteMessage(msg.ID, msg.Chat.ID, 20)

	case "waiting_for_cred_pass":
		state.TempCred.Password = update.Message.Text
		state.Step = "waiting_for_cred_key"
		msg, _ := b.SendMessage(ctx, &tbot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Теперь отправьте ключ для шифрования (длина должна быть 16 символов):",
		})
		bot.addDeleteMessage(msg.ID, msg.Chat.ID, 20)

	case "waiting_for_cred_key":
		inputMessage := update.Message.Text
		if len(inputMessage) != 16 {
			b.SendMessage(ctx, &tbot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "Неверная длина ключа шифрования, попробуйте еще раз",
			})
			return
		}
		bot.sessions.NewSession(userID, inputMessage)
		state.TempCred.Key = inputMessage
		state.Step = "waiting_for_cred_type"

		typeSelect, _ := bot.listTypesForSelect(userID)
		msg, _ := b.SendMessage(ctx, &tbot.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "Выберите категорию",
			ReplyMarkup: typeSelect,
		})
		bot.addDeleteMessage(msg.ID, msg.Chat.ID, 20)

	case "wait_key":
		inputMessage := update.Message.Text
		if len(inputMessage) != 16 {
			b.SendMessage(ctx, &tbot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "Неверная длина ключа шифрования, попробуйте еще раз",
			})
			return
		}
		delete(bot.userStates, userID)
		bot.sessions.NewSession(userID, inputMessage)
		msg, _ := b.SendMessage(ctx, &tbot.SendMessageParams{
			Text:   "Ключ успешно добавлен",
			ChatID: update.Message.Chat.ID,
		})
		bot.addDeleteMessage(msg.ID, msg.Chat.ID, 10)
		bot.addDeleteMessage(update.Message.ID, update.Message.Chat.ID, 5)
		isDelete = true
	}

	if !isDelete {
		bot.addDeleteMessage(update.Message.ID, update.Message.Chat.ID, 20)
	}
}
