package service

import (
	"context"
	"week/internal/model"
)

type NoteService interface {
	Create(ctx context.Context, indo *model.NoteInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.Note, error)
}
