package app

import (
	"context"
	"log"
	"week/internal/api/note"
	"week/internal/client/db"
	"week/internal/client/db/pg"
	"week/internal/client/db/transactions"
	"week/internal/closer"
	"week/internal/config"
	"week/internal/repository"
	noteRepository "week/internal/repository/note"
	"week/internal/service"
	noteService "week/internal/service/note"
)

type ServiceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient       db.Client
	txManager      db.TxManager
	noteRepository repository.NoteRepository
	noteService    service.NoteService

	noteImpl *note.Implementation
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

func (s *ServiceProvider) GetPGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatal(err)
		}
		s.pgConfig = cfg
	}
	return s.pgConfig
}

func (s *ServiceProvider) GetPgPool(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.GetPGConfig().DSN())
		if err != nil {
			log.Fatal(err)
		}
		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatal(err)
		}
		closer.Add(cl.Close)
		s.dbClient = cl
	}
	return s.dbClient
}

func (s *ServiceProvider) GetGRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatal(err)
		}
		s.grpcConfig = cfg
	}
	return s.grpcConfig
}

func (s *ServiceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.GetPGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *ServiceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transactions.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *ServiceProvider) GetNoteRepository(ctx context.Context) repository.NoteRepository {
	if s.noteRepository == nil {
		s.noteRepository = noteRepository.NewRepository(s.GetPgPool(ctx))
	}
	return s.noteRepository
}
func (s *ServiceProvider) GetNoteService(ctx context.Context) service.NoteService {
	if s.noteService == nil {
		s.noteService = noteService.NewService(
			s.GetNoteRepository(ctx),
			s.TxManager(ctx),
		)
	}
	return s.noteService
}

func (s *ServiceProvider) GetNoteImpl(ctx context.Context) *note.Implementation {
	if s.noteImpl == nil {
		s.noteImpl = note.NewImplementation(s.GetNoteService(ctx))
	}

	return s.noteImpl
}
