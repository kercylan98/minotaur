package server

import (
	"context"
)

type Network interface {
	OnSetup(ctx context.Context, event NetworkCore) error

	OnRun() error

	OnShutdown() error
}
