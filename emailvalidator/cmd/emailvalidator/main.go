package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/xoticdsign/go-emailvalidator/emailvalidator/internal/app"
	"github.com/xoticdsign/go-emailvalidator/emailvalidator/internal/config"
)

func mustRunLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "local":
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))

	case "dev":
		var file *os.File
		var err error

		file, err = os.Open("dev.log")
		if err != nil {
			file, err = os.Create("dev.log")
			if err != nil {
				panic("couldn't create a log file for dev mode: " + err.Error())
			}
		}

		log = slog.New(slog.NewJSONHandler(file, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))

	case "prod":
		var file *os.File
		var err error

		file, err = os.Open("prod.log")
		if err != nil {
			file, err = os.Create("prod.log")
			if err != nil {
				panic("couldn't create a log file for prod mode: " + err.Error())
			}
		}

		log = slog.New(slog.NewJSONHandler(file, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	}

	return log
}

func main() {
	cfg := config.MustLoad()

	log := mustRunLogger(cfg.Env)

	app := app.New(cfg.EmailValidatorConfig.Port, log)

	log.Info(
		"starting application",
		slog.Any("config", cfg),
	)

	go app.EmailValidator.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	log.Info(
		"stopping application",
		slog.Any("signal", sign),
	)

	app.EmailValidator.MustStop()
}
