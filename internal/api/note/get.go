package note

import (
	"context"
	"week/internal/converter"
	desc "week/pkg/note_v1"
)

// Get ...
func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	noteObj, err := i.noteService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		Note: converter.ToNoteFromService(noteObj),
	}, nil
}
