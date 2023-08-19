package monolith

import (
	"context"
	"database/sql"
	"eda-in-golang/internal/config"
	"eda-in-golang/internal/waiter"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type Monolith interface {
	Config() config.AppConfig
	DB() *sql.DB
	Logger() zerolog.Logger
	Mux() *chi.Mux
	RPC() *grpc.Server
	Waiter() waiter.Waiter
}

type Module interface {
	Startup(context.Context, Monolith) error
}
