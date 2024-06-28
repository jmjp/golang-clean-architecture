package database

import (
	"context"
	"onion/config"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type PostgresDB struct {
	Conn pgx.Conn
}

func NewPostgresDB() *PostgresDB {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, config.Get.DB_URL)
	if err != nil {
		panic(err)
	}
	return &PostgresDB{
		Conn: *conn,
	}
}

func (db *PostgresDB) Close() {
	db.Conn.Close(context.Background())
}

func (db *PostgresDB) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return db.Conn.Exec(ctx, query, args...)
}

func (db *PostgresDB) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return db.Conn.Query(ctx, query, args...)
}

func (db *PostgresDB) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return db.Conn.QueryRow(ctx, query, args...)
}
