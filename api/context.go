package main

import (
	"github.com/mauleyzaola/challenge/operations"
	"github.com/mauleyzaola/challenge/storage"
)

type context struct {
	storage operations.Storage
}

func newContext() *context {
	ctx := &context{}
	ctx.storage = &storage.Memory{}
	ctx.storage.Init()
	return ctx
}
