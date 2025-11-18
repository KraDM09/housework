package job

import (
	"context"
	"github.com/KraDM09/housework/internal/app/bot"
	"github.com/KraDM09/housework/internal/app/client/memcached"
	"github.com/KraDM09/housework/internal/app/config"
)

type jobUseCase struct {
	bot       bot.Bot
	cfg       *config.Values
	memcached memcached.Provider
}

func NewUseCase(
	bot bot.Bot,
	cfg *config.Values,
	memcached memcached.Provider,
) UseCase {
	return &jobUseCase{
		bot:       bot,
		cfg:       cfg,
		memcached: memcached,
	}
}

type UseCase interface {
	CreateNewTasks(
		ctx context.Context,
	) error
}
