package pg

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"week/internal/client/db"
)

type pgClient struct {
	masterDBC db.DB
}

func (p *pgClient) DB() db.DB {
	return p.masterDBC
}

func (p *pgClient) Close() error {
	if p.masterDBC != nil {
		p.masterDBC.Close()
	}
	return nil
}

func New(ctx context.Context, dsn string) (db.Client, error) {
	dbc, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}
	return &pgClient{
		masterDBC: &pg{dbc: dbc},
	}, nil
}
