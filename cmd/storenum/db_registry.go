package main

import (
	"context"

	"github.com/matrosovm/storenum/internal/pkg/database"
)

const (
	fileStorePath = "./log/store.log"
)

func dbConnect(ctx context.Context) database.Store {
	return database.NewStore(fileStorePath)
}
