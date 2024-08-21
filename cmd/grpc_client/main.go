package main

import (
	"github.com/fatih/color"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"time"
	desc "week/pkg/note_v1"
)

const (
	address = "localhost:50052"
	noteID  = 2
)

func main() {

	creds, err := credentials.NewClientTLSFromFile("serts/service.pem", "")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	c := desc.NewNoteV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Get(ctx, &desc.GetRequest{Id: noteID})
	if err != nil {
		log.Fatalf("failed to get note by id: %v", err)
	}

	log.Printf(color.RedString("Note info:\n"), color.GreenString("%+v", r.GetNote()))
}
