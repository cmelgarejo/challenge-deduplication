package database

import (
	"context"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	ErrMigrationNoChange = "no change"
)

type Repository interface {
	RunMigrations()
	Querier
}

type RepositoryImpl struct{ Querier }

func NewRepository(ctx context.Context) Repository {
	pgxCfg, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	if pool, err := pgxpool.ConnectConfig(ctx, pgxCfg); err != nil {
		fmt.Fprint(os.Stderr, err)
	} else {
		return &RepositoryImpl{
			New(pool),
		}
	}

	return nil
}

func (s *RepositoryImpl) RunMigrations() {
	fmt.Println("Running migrations...")
	defer fmt.Println("Finished")
	m, err := migrate.New(os.Getenv("DATABASE_MIGRATION_FILES"), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err)
	}
	if err := m.Up(); err != nil {
		if err.Error() == ErrMigrationNoChange {
			fmt.Fprint(os.Stdout, ErrMigrationNoChange)
		} else {
			fmt.Fprint(os.Stderr, err)
		}
	}
}
