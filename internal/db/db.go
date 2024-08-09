package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PsqlDatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"db"`
}

type PsqlClient struct {
	Db *pgxpool.Pool
}

func NewPsqlClient(conf *PsqlDatabaseConfig) (*PsqlClient, error) {
	dbPool, err := pgxpool.New(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
	))
	if err != nil {
		return nil, err
	}

	return &PsqlClient{Db: dbPool}, nil
}

func (client *PsqlClient) Close() {
	client.Db.Close()
}
