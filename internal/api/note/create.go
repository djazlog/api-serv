package note

import (
	"context"
	"log"
	"week/internal/converter"
	desc "week/pkg/note_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.noteService.Create(ctx, converter.ToNoteInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}
	log.Printf("Note id: %d", id)

	return &desc.CreateResponse{Id: id}, nil
}
