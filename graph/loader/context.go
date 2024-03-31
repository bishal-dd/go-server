package loader

import (
	"context"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)
// For returns the dataloader for a given context
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}