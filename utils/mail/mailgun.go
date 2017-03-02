package mail

import (
    "gopkg.in/mailgun/mailgun-go.v1"
)

type Mailgun struct {
    Domain          string
    ApiKey          string
    PublicApiKey    string

    Sender          string
    Subject         string
    Recipient       string
    Content         string
}

func NewMailgun(domain string, apiKey string, publicApiKey string) *Mailgun {
    return &Mailgun{
        Domain      : domain,
        ApiKey      : apiKey,
        PublicApiKey: publicApiKey,
    }
}

func (this *Mailgun) SetSender(sender string) *Mailgun {
    this.Sender = sender
    return this
}

func (this *Mailgun) SetSubject(subject string) *Mailgun {
    this.Subject = subject
    return this
}

func (this *Mailgun) SetRecipient(recipient string) *Mailgun {
    this.Recipient = recipient
    return this
}

func (this *Mailgun) SetContent(content string) *Mailgun {
    this.Content = content
    return this
}

func (this *Mailgun) Send() (response string, id string, err error) {
    mailer  := mailgun.NewMailgun(this.Domain, this.ApiKey, this.PublicApiKey)
    message := mailgun.NewMessage(
        this.Sender,
        this.Subject,
        this.Content,
        this.Recipient,
    )

    return mailer.Send(message)
}
