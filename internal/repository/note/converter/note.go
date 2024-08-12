package converter

import (
	"week/internal/model"
	modelRepo "week/internal/repository/note/model"
)

func ToNoteFromRepo(note *modelRepo.Note) *model.Note {
	return &model.Note{
		ID:        note.ID,
		Info:      ToNoteInfoFromRepo(note.Info),
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}
}

func ToNoteInfoFromRepo(noteInfo modelRepo.NoteInfo) model.NoteInfo {
	return model.NoteInfo{
		Title:   noteInfo.Title,
		Content: noteInfo.Content,
	}
}
