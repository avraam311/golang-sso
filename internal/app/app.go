package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/avraam311/golang-sso.git/internal/app/grpc"
	"github.com/avraam311/golang-sso.git/internal/services/auth"
	"github.com/avraam311/golang-sso.git/internal/storage/sqlite"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	storage, err := sqlite.New(storagePath)
	if err!= nil {
        panic(err)
    }

	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}