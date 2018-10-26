// +build wireinject

package main

import (
	"context"
	"github.com/google/go-cloud/wire"
	"github.com/alextuan1024/wire-test/foobarbaz"
)

func initializeBaz(ctx context.Context) (foobarbaz.Baz, error) {
	wire.Build(foobarbaz.SuperSet)
	return foobarbaz.Baz{}, nil
}
