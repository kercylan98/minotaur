package server

import "context"

type Network interface {
	OnSetup(ctx context.Context, core Core) error

	OnRun(ctx context.Context) error

	OnShutdown() error
}
