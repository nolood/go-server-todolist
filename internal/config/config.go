package config

import (
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Env string `yaml:"env" env-default:"dev"`
}

var (
	Logger   *slog.Logger
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func MustLoad() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	viper.AutomaticEnv()
}

func InitLogger() {
	switch viper.GetString("ENV") {
	case envLocal:
		Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	Logger.With(slog.String("env", viper.GetString("ENV")))
}
