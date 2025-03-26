package app

import (
	"log/slog"

	eval "github.com/xoticdsign/grpcemailvalidator/emailvalidator/internal/app/emailvalidator"
)

type App struct {
	EmailValidator *eval.App
}

func New(port int, log *slog.Logger) *App {
	eval := eval.New(port, log)

	return &App{
		EmailValidator: eval,
	}
}
