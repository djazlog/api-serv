package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	noteApi "week/internal/api/note"
	noteRepository "week/internal/repository/note"
	noteService "week/internal/service/note"
	desc "week/pkg/note_v1"
)

const grpcPort = 50051

func main() {
	ctx := context.Background()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	pool, err := pgxpool.Connect(ctx, os.Getenv("DATABASE_URL"))

	noteRepo := noteRepository.NewRepository(pool)
	noteSrv := noteService.NewService(noteRepo)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterNoteV1Server(s, noteApi.NewImplementation(noteSrv))

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
