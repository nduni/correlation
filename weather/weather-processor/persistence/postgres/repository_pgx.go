package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nduni/correlation/weather/weather-processor/app"
)

type PGXRepository struct {
	DbPool pgxpool.Pool
}

func NewPGXRepository() (*PGXRepository, error) {
	db := PGXRepository{}
	err := db.Connect()

	return &db, err
}

func (db PGXRepository) Connect() error {
	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslomode=%s pool_max_conns=6",
		app.Config.Db.DB_HOST, app.Config.Db.DB_PORT, app.Config.Db.DB_USER, app.Config.Db.DB_PASSWORD, app.Config.Db.DB_NAME, app.Config.Db.DB_SSL)

	dbPool, err := pgxpool.Connect(context.Background(), config)
	if err != nil {
		return fmt.Errorf("unable to connect to database %v", err)
	}

	db.DbPool = *dbPool
	return nil
}
