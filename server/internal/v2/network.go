package server

import (
	"context"
)

type Network interface {
	OnSetup(ctx context.Context, controller Controller) error

	OnRun() error

	OnShutdown() error
}
