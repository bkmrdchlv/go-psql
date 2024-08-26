package postgresql

import (
	"context"
	"fmt"

	"github.com/go-psql/config"

	"github.com/go-psql/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db *pgxpool.Pool

	// SomethingI
	somethingInterfaceVar storage.SomethingI
}

func NewPostrgesqlConnection(ctx context.Context, cfg config.Config) (storage.Storage, error) {
	configPsql, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		return nil, err
	}

	configPsql.MaxConns = cfg.PostgresMaxConnections

	db, err := pgxpool.ConnectConfig(ctx, configPsql)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
}

//////////////////////////////////////////////////

func (s *Store) SomethingI() storage.SomethingI {

	if s.somethingInterfaceVar == nil {
		s.somethingInterfaceVar = NewSomethingI(s.db)
	}
	return s.somethingInterfaceVar
}

func (s *Store) CloseDB() {
	s.db.Close()
}
