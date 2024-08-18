package note

import (
	"week/internal/client/db"
	"week/internal/repository"
	def "week/internal/service"
)

// / Валидация , что структура соответствует интерфейсу
var _ def.NoteService = (*serv)(nil)

type serv struct {
	noteRepository repository.NoteRepository
	txManager      db.TxManager
}

func NewService(noteRepository repository.NoteRepository, txManager db.TxManager) *serv {
	return &serv{
		noteRepository: noteRepository,
		txManager:      txManager,
	}
}

func NewMockService(deps ...interface{}) *serv {
	srv := serv{}

	for _, v := range deps {
		switch s := v.(type) {
		case repository.NoteRepository:
			srv.noteRepository = s
		}
	}

	return &srv
}
