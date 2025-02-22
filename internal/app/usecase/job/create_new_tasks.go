package job

import (
	"context"
	"fmt"
	"github.com/KraDM09/housework/internal/app/usecase/job/models"
	"math/rand"
	"strings"
)

var tasks = []string{
	"полы",
	"пылесос",
	"зеркала",
	"стирка",
	"санузел",
	"ванная",
	"плита",
	"развесить бельё",
}

func (u *jobUseCase) CreateNewTasks(
	ctx context.Context,
) error {
	shuffledTasks := u.getShuffledTasks(tasks)

	users := []models.User{
		{
			ChatId: u.cfg.Users.UserChatId1,
			Tasks:  shuffledTasks[:4],
		},
		{
			ChatId: u.cfg.Users.UserChatId2,
			Tasks:  shuffledTasks[4:],
		},
	}

	for _, usr := range users {
		err := u.bot.SendMessage(ctx,
			usr.ChatId,
			fmt.Sprintf("Тебе выпали🫡: %s", strings.Join(usr.Tasks, ", ")),
		)

		if err != nil {
			panic(err)
		}

		fmt.Print("Tasks for user ", usr.ChatId, ": ", usr.Tasks)
	}

	return nil
}

func (u *jobUseCase) getShuffledTasks(tasks []string) []string {
	taskIds := make([]int, 0, len(tasks))

	for i, _ := range tasks {
		taskIds = append(taskIds, i)
	}

	shuffledTasks := make([]string, 0, len(tasks))

	for _, taskId := range u.shuffle(taskIds) {
		shuffledTasks = append(shuffledTasks, tasks[taskId])
	}

	return shuffledTasks
}

func (u *jobUseCase) shuffle(numbers []int) []int {
	for i := range numbers {
		j := rand.Intn(i + 1)
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}

	return numbers
}
