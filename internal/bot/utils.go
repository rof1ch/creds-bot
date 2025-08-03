package bot

import (
	"fmt"
	"time"

	"github.com/go-telegram/bot/models"
)

var defaultMenu = &models.InlineKeyboardMarkup{
	InlineKeyboard: [][]models.InlineKeyboardButton{
		{
			{
				Text:         "Список данных",
				CallbackData: "data_list",
			},
		},
		{
			{
				Text:         "Список категорий",
				CallbackData: "type_list",
			},
		},
	},
}

func (bot *Bot) listTypes(userID int64) (*models.InlineKeyboardMarkup, error) {
	var menu models.InlineKeyboardMarkup
	types, err := bot.services.TypeCred.List(userID)
	if err != nil {
		return nil, err
	}
	// lenTypes := len(types)
	for _, t := range types {
		menu.InlineKeyboard = append(menu.InlineKeyboard, []models.InlineKeyboardButton{
			{
				Text:         fmt.Sprintf("%s %s", t.Icon, t.Name),
				CallbackData: fmt.Sprintf("user_type_%d", t.Id),
			},
		})
	}

	menu.InlineKeyboard = append(menu.InlineKeyboard, []models.InlineKeyboardButton{
		{
			Text:         "Добавить категорию",
			CallbackData: "type_add",
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

func typeByIdMenu(typeId uint) *models.InlineKeyboardMarkup {
	var menu models.InlineKeyboardMarkup
	menu.InlineKeyboard = append(menu.InlineKeyboard, []models.InlineKeyboardButton{
		{
			Text:         "Удалить",
			CallbackData: fmt.Sprintf("type_delete_%d", typeId),
		},
		{
			Text:         "Список данных",
			CallbackData: fmt.Sprintf("data_type_%d", typeId),
		},
	})
	menu.InlineKeyboard = append(menu.InlineKeyboard, []models.InlineKeyboardButton{
		{
			Text:         "<- Назад",
			CallbackData: "type_list",
		},
	})
	return &menu
}

func (bot *Bot) addDeleteMessage(messageID int, chatID int64, secDuration time.Duration) {
	bot.deleteMessages = append(bot.deleteMessages, DeleteMessage{
		MessageID: messageID,
		ChatID:    chatID,
		ExpiredAt: time.Now().Add(time.Second * secDuration),
	})
}

func (bot *Bot) credsByTypeIdMenu(typeId uint) (*models.InlineKeyboardMarkup, error) {
	var menu models.InlineKeyboardMarkup
	creds, err := bot.services.ByTypeId(typeId)
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
			CallbackData: fmt.Sprintf("user_type_%d", typeId),
		},
	})
	return &menu, nil
}

func (bot *Bot) listTypesForSelect(userID int64) (*models.InlineKeyboardMarkup, error) {
	var menu models.InlineKeyboardMarkup
	types, err := bot.services.TypeCred.List(userID)
	if err != nil {
		return nil, err
	}
	// lenTypes := len(types)
	for _, t := range types {
		menu.InlineKeyboard = append(menu.InlineKeyboard, []models.InlineKeyboardButton{
			{
				Text:         fmt.Sprintf("%s %s", t.Icon, t.Name),
				CallbackData: fmt.Sprintf("cred_type_%d", t.Id),
			},
		})
	}

	return &menu, nil
}
