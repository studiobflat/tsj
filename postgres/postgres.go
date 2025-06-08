package postgres

import (
	"context"
	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
	"github.com/jackc/pgx-zap"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/studiobflat/tsj/logger"
)

type Postgres struct {
	*pgxpool.Pool
}

func NewPostgres(config *Config) (*Postgres, error) {
	pgConfig, err := pgxpool.ParseConfig(config.Url)
	if err != nil {
		return nil, err
	}

	tracer := &tracelog.TraceLog{
		Logger:   zap.NewLogger(logger.GetLogger("db").Desugar()),
		LogLevel: tracelog.LogLevel(config.LogLevel),
	}

	pgConfig.MaxConns = config.MaxConnection

	pgConfig.MinConns = config.MinConnection

	pgConfig.MaxConnIdleTime = config.MaxConnectionIdleTime

	pgConfig.ConnConfig.Tracer = tracer

	pgConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		// Register the decimal type
		pgxdecimal.Register(conn.TypeMap())

		return nil
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), pgConfig)
	if err != nil {
		return nil, err
	}

	return &Postgres{
		Pool: pool,
	}, nil
}
