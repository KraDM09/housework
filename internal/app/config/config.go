package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"

	"github.com/KraDM09/housework/internal/app/constants"
)

type Values struct {
	Bot       *BotConfig   `envconfig:"BOT" required:"true"`
	Users     *UsersConfig `envconfig:"USERS" required:"true"`
	Memcached Memcached    `envconfig:"MEMCACHED" required:"true"`
}

type Memcached struct {
	Server string `envconfig:"SERVER" required:"true"`
}

type BotConfig struct {
	Token string `envconfig:"TOKEN" required:"true"`
}

type UsersConfig struct {
	UserChatId1 int64 `envconfig:"CHAT_ID_1" required:"true"`
	UserChatId2 int64 `envconfig:"CHAT_ID_2" required:"true"`
}

func New() (*Values, error) {
	err := LoadEnvFile()
	if err != nil {
		return nil, err
	}

	cfg := &Values{}
	err = envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func LoadEnvFile() error {
	if needUseLocalEnvFile() {
		err := godotenv.Load(constants.DefaultEnvFile)
		if err != nil {
			return err
		}
	}
	return nil
}

func needUseLocalEnvFile() bool {
	for _, arg := range os.Args {
		if arg == constants.UseLocalEnvFileArg {
			return true
		}
	}
	return false
}
