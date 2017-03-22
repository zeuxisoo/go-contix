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
