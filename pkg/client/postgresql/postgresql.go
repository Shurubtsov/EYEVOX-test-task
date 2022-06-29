package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/dshurubtsov/internal/config"
	"github.com/dshurubtsov/pkg/logging"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Interface to implementing methods from pure driver PGX
type Client interface {
	Exec(ctx context.Context, sql string, argumets ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, argumets ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, argumets ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

// Create new client for work with driver of database
func NewClient(ctx context.Context, maxAttemts int, sc config.StorageConfig, logger *logging.Logger) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", sc.Username, sc.Password, sc.Host, sc.Port, sc.Database)
	logger.Info("dsn -> ", dsn)

	err = doWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			return err
		}

		return nil

	}, maxAttemts, 5*time.Second)
	if err != nil {
		logger.Fatal("Error with tries connect to postgres")
	}

	logger.Info("[OK] Connected to postgres")
	return pool, nil
}

// Utility func for many attemts to connecting to database
func doWithTries(fn func() error, attemts int, delay time.Duration) (err error) {
	for attemts > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attemts--

			continue
		}
		return nil
	}

	return
}
