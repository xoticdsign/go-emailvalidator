package emailvalidator

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	evalservice "github.com/xoticdsign/go-emailvalidator/emailvalidator/internal/service/emailvalidator"
	evalproto "github.com/xoticdsign/go-emailvalidator/proto/gen/emailvalidator"
)

type App struct {
	GRPCServer *grpc.Server
	Port       int
}

type api struct {
	evalproto.UnimplementedEmailValidatorServer
	evalservice *evalservice.Service
}

func New(port int, log *slog.Logger) *App {
	grpcServer := grpc.NewServer()

	evalservice := evalservice.New(log)

	evalproto.RegisterEmailValidatorServer(grpcServer, &api{evalservice: evalservice})

	return &App{
		GRPCServer: grpcServer,
		Port:       port,
	}
}

func (a *App) MustRun() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.Port))
	if err != nil {
		panic("grpc server couldn't listen: " + err.Error())
	}

	err = a.GRPCServer.Serve(listener)
	if err != nil {
		panic("grpc server couldn't serve: " + err.Error())
	}
}

func (a *App) MustStop() {
	a.GRPCServer.GracefulStop()
}

func (a *api) Validate(ctx context.Context, req *evalproto.ValidateRequest) (*evalproto.ValidateResponse, error) {
	if req.GetEmailToValidate() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "email must be in the request")
	}

	isValid, err := a.evalservice.Validate(ctx, req.GetEmailToValidate())
	if err != nil {
		switch {
		case errors.Is(err, evalservice.ErrBadFormat):
			return nil, status.Error(codes.InvalidArgument, err.Error())

		case errors.Is(err, evalservice.ErrBadHost):
			return nil, status.Error(codes.InvalidArgument, err.Error())

		default:
			return nil, status.Error(codes.Internal, "internal error")
		}
	}

	return &evalproto.ValidateResponse{
		EmailIsValid: isValid,
	}, nil
}
