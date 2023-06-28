package main

import (
	"context"
)

func main() {
	ctx := context.Background()

	store := dbConnect(ctx)

	go runRest(ctx)
	runGrpc(store)
}
