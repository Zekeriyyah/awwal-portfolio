package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/Zekeriyyah/my-portfolio/internal/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
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

	userName := os.Getenv("MYSQL_USERNAME")
	dbName := os.Getenv("MYSQL_DBNAME")
	password := os.Getenv("MYSQL_PASSWORD")

	addr := flag.String("addr", ":5005", "HTTP Network address")
	dsn := flag.String("dsn", userName+":"+password+"@/"+dbName+"?parseTime=true", "Database data source name")

	flag.Parse()

	db, err = openDB(*dsn) // Connecting to mysql database
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

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
		Addr:    *addr,
		Handler: app.Route(),
	}

	infoLog.Printf("Starting a server on port %s\n", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
