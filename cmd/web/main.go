package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/alex-orkuma/foodexp/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
)

// A Custom application struct for dependency injection
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	products *mysql.ProductModel
}

func main() {

	// Go server address flag
	addr := flag.String("addr", ":4000", "HTTP Network Address")

	// MySql Data source name
	dsn := flag.String("dsn", "foodexp:Or$kumashnd417@/foodexp?parseTime=true", "MySql Data source Name")
	flag.Parse()

	// New Info Logger
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// New Error Logger
	erroLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		erroLog.Fatal(err)
	}

	defer db.Close()
	app := &application{
		errorLog: erroLog,
		infoLog:  infoLog,
		products: &mysql.ProductModel{DB: *db},
	}

	// Initialize a server struct
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: erroLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("starting a server on port %s", *addr)
	err = srv.ListenAndServe()
	erroLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
