package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq" // Third party package
	"levijames.net/test/pkg/models/postgresql"
)

func setUpDB(dsn string) (*sql.DB, error) {
	// Provide the credentials for our database
	const (
		host     = "localhost"
		port     = 5432
		user     = "music"
		password = "$swordfish$"
		dbname   = "music"
	)

	// Establish a connection to the database
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	// Test our connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Dependencies (things/variables)
// Dependency Injection (passing)
type application struct {
	music *postgresql.MusicModel
}

func main() {
	//create a command line flag
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", os.Getenv("MUSIC_DB_DSN"),
		"PostgreSQL DSN (Data Source Name)")

	flag.Parse()

	var db, err = setUpDB(*dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // Always do this before exiting
	app := &application{
		music: &postgresql.MusicModel{
			DB: db,
		},
	}

	//create custom web server
	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}
	// start server
	log.Printf("Starting server on port %s", *addr)
	err = srv.ListenAndServe()
	log.Fatal(err)
}
