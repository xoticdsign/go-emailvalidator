package emailvalidator

import (
	"context"
	"errors"
	"log/slog"

	"github.com/badoux/checkmail"
)

var (
	ErrBadFormat = errors.New("email format is invalid")
	ErrBadHost   = errors.New("email host is invalid")
)

type Service struct {
	Log *slog.Logger
}

func New(log *slog.Logger) *Service {
	return &Service{
		Log: log,
	}
}

func (s *Service) Validate(ctx context.Context, emailToValidate string) (bool, error) {
	const op = "service.emailvalidator.emailvalidator.Validate()"

	s.Log.Info(
		"validating "+emailToValidate,
		slog.String("op", op),
	)

	err := checkmail.ValidateFormat(emailToValidate)
	if err != nil {
		if errors.Is(err, checkmail.ErrBadFormat) {
			return false, ErrBadFormat
		}
		s.Log.Error(
			"format validation failed: "+err.Error(),
			slog.String("op", op),
		)

		return false, err
	}

	err = checkmail.ValidateHost(emailToValidate)
	if err != nil {
		if errors.Is(err, checkmail.ErrUnresolvableHost) {
			return false, ErrBadHost
		}
		s.Log.Error(
			"host validation failed: "+err.Error(),
			slog.String("op", op),
		)

		return false, err
	}

	return true, nil
}
