package app

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"week/internal/api/note"
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

	pgPool         *pgxpool.Pool
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

func (s *ServiceProvider) GetPgPool(ctx context.Context) *pgxpool.Pool {
	if s.pgPool == nil {
		pool, err := pgxpool.Connect(ctx, s.GetPGConfig().DSN())
		if err != nil {
			log.Fatal(err)
		}
		err = pool.Ping(ctx)
		if err != nil {
			log.Fatal(err)
		}
		closer.Add(func() error {
			pool.Close()
			return nil
		})
		s.pgPool = pool
	}
	return s.pgPool
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

func (s *ServiceProvider) GetNoteRepository(ctx context.Context) repository.NoteRepository {
	if s.noteRepository == nil {
		s.noteRepository = noteRepository.NewRepository(s.GetPgPool(ctx))
	}
	return s.noteRepository
}
func (s *ServiceProvider) GetNoteService(ctx context.Context) service.NoteService {
	if s.noteService == nil {
		s.noteService = noteService.NewService(s.GetNoteRepository(ctx))
	}
	return s.noteService
}

func (s *ServiceProvider) GetNoteImpl(ctx context.Context) *note.Implementation {
	if s.noteImpl == nil {
		s.noteImpl = note.NewImplementation(s.GetNoteService(ctx))
	}

	return s.noteImpl
}
