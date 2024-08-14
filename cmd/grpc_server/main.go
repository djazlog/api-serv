package main

import (
	"context"
	"log"
	"week/internal/app"
)

const grpcPort = 50051

func main() {
	ctx := context.Background()

	/*lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	pool, err := pgxpool.Connect(ctx, os.Getenv("DATABASE_URL"))

	noteRepo := noteRepository.NewRepository(pool)
	noteSrv := noteService.NewService(noteRepo)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterNoteV1Server(s, noteApi.NewImplementation(noteSrv))

	log.Printf("server listening at %v", lis.Addr())*/
	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run()

	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
