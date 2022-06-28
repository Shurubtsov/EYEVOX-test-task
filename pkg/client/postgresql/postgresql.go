package postgresql

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dshurubtsov/internal/config"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Client interface {
	Exec(ctx context.Context, sql string, argumets ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, argumets ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, argumets ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, maxAttemts int, sc config.StorageConfig) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", sc.Username, sc.Password, sc.Host, sc.Port, sc.Database)

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
		log.Fatal("error with tries connect to postgre")
	}

	return pool, nil
}

func doWithTries(fn func() error, attemts int, delay time.Duration) (err error) {
	for attemts < 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attemts--

			continue
		}
		return nil
	}

	return
}
