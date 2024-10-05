package lib

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/YarKhan02/Smart-Hospital-System/projectPath"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
)

var db dbConnection

type dbConnection struct {
	pool *pgx.Conn
}

func (db *dbConnection) ConnectToDB() {
	connString := "postgresql://postgres:password@db:5432/smart_hospital_system"

	// Establish a connection to the database
	pool, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	fmt.Println("Connected to PostgreSQL!")

	db.pool = pool
}

func dbInit() {
	db.ConnectToDB()
}

func Template(args string) string {
	pathing := []string{projectpath.Root, "/", args, ".sql"}
	template_path := strings.Join(pathing, "")

	content, err := os.ReadFile(template_path)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func (db *dbConnection) query_object(sql string, result interface{}) error {
	dbInit()
	defer db.pool.Close(context.Background())

	err := pgxscan.Select(context.Background(), db.pool, result, sql)
	if err != nil {
		return fmt.Errorf("error executing query: %v", err)
	}

	return nil
}

func (db *dbConnection) query_commit(sql string, )

func QueryJson(sql string, result interface{}) error {
	return db.query_object(sql, result)
}