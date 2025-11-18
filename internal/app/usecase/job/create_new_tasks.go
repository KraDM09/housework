package job

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/KraDM09/housework/internal/app/usecase/job/models"
	"math/rand"
	"strings"
	"time"
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
	"протереть пыль",
	"микроволновка",
}

func (u *jobUseCase) CreateNewTasks(
	ctx context.Context,
) error {
	weekKey := u.getWeekKey()

	shuffledTasks := make([]string, 0)

	cachedData, err := u.memcached.Get(weekKey)
	if err != nil {
		return err
	}

	if cachedData != nil {
		err = json.Unmarshal(cachedData, &shuffledTasks)
		if err != nil {
			return err
		}
	}

	message := "Тебе выпали"

	if cachedData != nil {
		message = "Напоминание о задачах"
	} else {
		shuffledTasks = u.getShuffledTasks(tasks)
	}

	users := []models.User{
		{
			ChatId: u.cfg.Users.UserChatId1,
			Tasks:  shuffledTasks[:5],
		},
		{
			ChatId: u.cfg.Users.UserChatId2,
			Tasks:  shuffledTasks[5:],
		},
	}

	for _, usr := range users {
		err := u.bot.SendMessage(ctx,
			usr.ChatId,
			fmt.Sprintf("%s🫡: %s", message, strings.Join(usr.Tasks, ", ")),
		)
		if err != nil {
			panic(err)
		}

		fmt.Print("Tasks for user ", usr.ChatId, ": ", usr.Tasks)
	}

	if cachedData == nil {
		err = u.memcached.Set(weekKey, shuffledTasks, 7*24*60*60)
		if err != nil {
			return err
		}
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

func (u *jobUseCase) getWeekKey() string {
	now := time.Now().In(time.Local)
	_, week := now.ISOWeek()
	return fmt.Sprintf("week_%d", week)
}
