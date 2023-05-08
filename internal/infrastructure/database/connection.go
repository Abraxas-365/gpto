package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Connection struct {
	Pool *pgxpool.Pool
}

func NewConnection(host string, port int, user string, password string, dbname string) (*Connection, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		return nil, err
	}

	return &Connection{Pool: pool}, nil
}

func (c *Connection) Close() {
	c.Pool.Close()
}
