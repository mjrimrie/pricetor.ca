package datalayer

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
	"sync"
)

var _db *pgxpool.Pool

var once sync.Once

var singletonErr error

func Connect() (*pgxpool.Pool, error) {

	once.Do(func() {
		poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
		if err != nil {
			singletonErr = err
			return
		}
		poolConfig.ConnConfig.Logger =m
		_db, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
		if err != nil {
			singletonErr = err
			return
		}
	})
	if singletonErr != nil {
		return new(pgxpool.Pool), singletonErr
	}
	return _db, nil
}
