package note

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"week/internal/model"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.Note, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "note.Get")
	defer span.Finish()
	span.SetTag("id", id)
	note, err := s.noteRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return note, nil
}
