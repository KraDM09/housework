package job

import (
	"context"
	"github.com/KraDM09/housework/internal/app/bot"
	"github.com/KraDM09/housework/internal/app/config"
)

type jobUseCase struct {
	bot bot.Bot
	cfg *config.Values
}

func NewUseCase(
	bot bot.Bot,
	cfg *config.Values,
) UseCase {
	return &jobUseCase{
		bot: bot,
		cfg: cfg,
	}
}

type UseCase interface {
	CreateNewTasks(
		ctx context.Context,
	) error
}
