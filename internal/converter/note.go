package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"week/internal/model"
	desc "week/pkg/note_v1"
)

func ToNoteFromService(note *model.Note) *desc.Note {
	var updatedAt *timestamppb.Timestamp
	if note.UpdatedAt.Valid {
		updatedAt = timestamppb.New(note.UpdatedAt.Time)
	}
	return &desc.Note{
		Id:        note.ID,
		Info:      ToNoteInfoFromService(note.Info),
		CreatedAt: timestamppb.New(note.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToNoteInfoFromService(noteInfo model.NoteInfo) *desc.NoteInfo {
	return &desc.NoteInfo{
		Title:   noteInfo.Title,
		Content: noteInfo.Content,
	}
}

func ToNoteInfoFromDesc(info *desc.NoteInfo) *model.NoteInfo {
	return &model.NoteInfo{
		Title:   info.Title,
		Content: info.Content,
	}
}
