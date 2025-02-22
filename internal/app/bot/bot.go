package bot

import (
	"context"
	"github.com/KraDM09/housework/internal/app/config"
	tgBot "github.com/go-telegram/bot"
)

func New(
	_ context.Context,
	cfg *config.Values,
) (Bot, error) {
	b, err := tgBot.New(cfg.Bot.Token)
	if err != nil {
		return nil, err
	}

	return &bot{
		Bot: b,
	}, nil
}

type bot struct {
	Bot *tgBot.Bot
}

type Bot interface {
	Start(ctx context.Context)
	SendMessage(
		ctx context.Context,
		chatID int64,
		message string,
	) error
}

func (b *bot) Start(ctx context.Context) {
	botCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	b.Bot.Start(botCtx)
}

func (b *bot) SendMessage(
	ctx context.Context,
	chatID int64,
	message string,
) error {
	_, err := b.Bot.SendMessage(ctx, &tgBot.SendMessageParams{
		ChatID: chatID,
		Text:   message,
	})

	return err
}
