package telegram

import (
    "gopkg.in/telegram-bot-api.v4"
)

type Telegram struct {
    Token string
    Bot   *tgbotapi.BotAPI
}

func NewTelegram(token string) (*Telegram, error) {
    bot, err := tgbotapi.NewBotAPI(token)

    return &Telegram{
        Token: token,
        Bot  : bot,
    }, err
}

func (this *Telegram) SetDebug(enable bool) *Telegram {
    this.Bot.Debug = enable

    return this
}

func (this Telegram) ValidateToken() bool {
    return len(this.Bot.Self.UserName) > 0
}

func (this Telegram) GetBotUsername() string {
    return this.Bot.Self.UserName
}

func (this Telegram) GetUpdatesChannel() (<-chan tgbotapi.Update, error) {
    updateConfig := tgbotapi.NewUpdate(0)
    updateConfig.Timeout = 60

    return this.Bot.GetUpdatesChan(updateConfig)
}

func (this Telegram) CreateMessage(chatId int64, text string) tgbotapi.MessageConfig {
    return tgbotapi.NewMessage(chatId, text)
}

func (this Telegram) SendMessage(message tgbotapi.Chattable) (tgbotapi.Message, error) {
    return this.Bot.Send(message)
}
