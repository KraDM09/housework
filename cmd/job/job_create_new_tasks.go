package main

import (
	"context"
	"github.com/KraDM09/housework/internal/app/bot"
	"github.com/KraDM09/housework/internal/app/config"
	"github.com/KraDM09/housework/internal/app/usecase/job"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	telegramBot, err := bot.New(ctx, cfg)

	if err != nil {
		panic(err)
	}

	newTasksUseCase := job.NewUseCase(
		telegramBot,
		cfg,
	)

	err = newTasksUseCase.CreateNewTasks(ctx)

	if err != nil {
		panic(err)
	}
}
