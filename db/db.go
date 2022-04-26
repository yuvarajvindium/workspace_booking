package db

import (
	"context"
	"workspace_booking/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

var dbPool *pgxpool.Pool

func GetDbConnectionPool() *pgxpool.Pool {
	if dbPool != nil {
		return dbPool
	}

	psqlconn := config.GetDBConnectionURL()
	println(psqlconn)
	db, err := pgxpool.Connect(context.Background(), psqlconn)

	// open database
	checkError(err)
	dbPool = db
	return dbPool
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
