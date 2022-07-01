package core

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Token           string `yaml:"token" env:"TOKEN"`
	ConfirmationKey string `yaml:"confirmation_key" env:"CONFIRMATION_KEY"`
	Host            string `yaml:"host" env:"HOST"`
}

func (s *Config) Load(filename string) error {
	if err := cleanenv.ReadConfig(filename, s); err != nil {
		return err
	}

	return nil
}