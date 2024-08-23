package note

import (
	"week/internal/service"
	descAuth "week/pkg/auth_v1"
)

type Implementation struct {
	descAuth.UnimplementedAuthV1Server
	noteService service.NoteService
}

func NewImplementation(noteService service.NoteService) *Implementation {
	return &Implementation{noteService: noteService}
}
