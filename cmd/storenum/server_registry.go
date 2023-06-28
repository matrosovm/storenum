package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/matrosovm/storenum/internal/app/storenum"
	"github.com/matrosovm/storenum/internal/pkg/database"
	pbStorenum "github.com/matrosovm/storenum/pkg/api/storenum"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func runRest(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	rmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pbStorenum.RegisterStorenumHandlerFromEndpoint(ctx, rmux, "localhost:12201", opts)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	prefix := "/swagger/"
	fs := http.FileServer(http.Dir("./config/swagger/swaggerui"))
	mux.Handle(prefix, http.StripPrefix(prefix, fs))
	mux.Handle("/", rmux)

	server := &http.Server{
		Addr:         "localhost:8080",
		Handler:      mux,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}

	log.Printf("server listening at 8080")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func runGrpc(store database.Store) {
	lis, err := net.Listen("tcp", "localhost:12201")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pbStorenum.RegisterStorenumServer(server, storenum.NewService(store))
	log.Printf("server listening at %v", lis.Addr())
	if err = server.Serve(lis); err != nil {
		panic(err)
	}
}
