package note

import (
	"week/internal/repository"
	def "week/internal/service"
)

// / Валидация , что структура соответствует интерфейсу
var _ def.NoteService = (*serv)(nil)

type serv struct {
	noteRepository repository.NoteRepository
	//txManager      db.TxManager
}

func NewService(noteRepository repository.NoteRepository) *serv {
	return &serv{
		noteRepository: noteRepository,
		//txManager:      txManager,
	}
}
