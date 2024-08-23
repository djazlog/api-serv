package repository

import (
	"context"
	"week/internal/model"
)

type NoteRepository interface {
	Create(ctx context.Context, info *model.NoteInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.Note, error)
}

type AuthRepository interface {
	Login(ctx context.Context, user *model.UserInfo) (string, error)
}
