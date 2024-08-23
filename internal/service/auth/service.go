package note

import (
	"week/internal/client/db"
	"week/internal/repository"
	def "week/internal/service"
)

// / Валидация , что структура соответствует интерфейсу
var _ def.AuthService = (*serv)(nil)

type serv struct {
	authRepository repository.AuthRepository
	txManager      db.TxManager
}

func NewService(authRepository repository.AuthRepository, txManager db.TxManager) *serv {
	return &serv{
		authRepository: authRepository,
		txManager:      txManager,
	}
}
