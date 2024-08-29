package converter

import (
	"database/sql"
	"time"

	"week/internal/model"
	modelRepo "week/internal/repository/note/redis/model"
)

func ToNoteFromRepo(note *modelRepo.Note) *model.Note {
	var updatedAt sql.NullTime
	if note.UpdatedAtNs != nil {
		updatedAt = sql.NullTime{
			Time:  time.Unix(0, *note.UpdatedAtNs),
			Valid: true,
		}
	}

	return &model.Note{
		ID: note.ID,
		Info: model.NoteInfo{
			Title:   note.Title,
			Content: note.Content,
		},
		CreatedAt: time.Unix(0, note.CreatedAtNs),
		UpdatedAt: updatedAt,
	}
}
