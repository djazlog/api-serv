package transaction

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"week/internal/client/db"
	"week/internal/client/db/pg"
)

type manager struct {
	db db.Transactor
}

func NewTransactionManager(db db.Transactor) db.TxManager {
	return &manager{db: db}
}

func (m *manager) transaction(ctx context.Context, opts pgx.TxOptions, fn db.Handler) (err error) {
	tx, ok := ctx.Value(pg.TxKey).(pgx.Tx)
	if ok {
		return fn(ctx)
	}

	tx, err = m.db.BeginTx(ctx, opts)
	if err != nil {
		return err
	}

	ctx = pg.MakeContextTx(ctx, tx)

	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("panic: %v", r)
		}
		if err != nil {
			if errRollBack := tx.Rollback(ctx); errRollBack != nil {
				err = errors.Wrapf(err, "rollback: %v", errRollBack)
			}

			return
		}
		if nil == err {
			err = tx.Commit(ctx)
			if err != nil {
				err = errors.Wrapf(err, "commit: %v", tx)
			}
		}
	}()

	if err = fn(ctx); err != nil {
		err = errors.Wrap(err, "failed executing code inside transaction")
	}

	return err
}

func (m *manager) ReadCommited(ctx context.Context, f db.Handler) error {
	txOpts := pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
	return m.transaction(ctx, txOpts, f)
}
