package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Zekeriyyah/my-portfolio/internal/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

type App struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file!")
	}

	addr := flag.String("addr", ":5005", "HTTP Network address")
	flag.Parse()

	host := os.Getenv("PQSQL_HOST")
	port := os.Getenv("PQSQL_PORT")
	userName := os.Getenv("PQSQL_USERNAME")
	dbName := os.Getenv("PQSQL_DBNAME")
	password := os.Getenv("PQSQL_PASSWORD")

	postgresInfo := fmt.Sprintf("host=%s port=%s username=%s passwordo=%s dbname=%s sslmode=disable",
		host, port, userName, password, dbName)

	db, err = openDB(postgresInfo) // Connecting to mysql database
	if err != nil {
		log.Printf("Database connection failed!\n%v", err)
		return
	}

	defer db.Close()

	err = models.Initialize(db)
	if err != nil {
		log.Fatalf("database initialization failed! -- %v", err)
	}

	infoLog := log.New(os.Stderr, "INFO\t", log.Ldate|log.Ltime)                  // initializing an info logger
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile) // initializing an error logger

	app := &App{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.Route(),
	}

	infoLog.Printf("Starting a server on port %s\n", *addr)
	// err = srv.ListenAndServe()
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
