package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// A Custom application struct for dependency injection
type application struct	{
	errorLog *log.Logger
	infoLog *log.Logger

}

func main() {

	// Go server address flag
	addr := flag.String("addr", ":4000", "HTTP Network Address")

	// MySql Data source name
	dsn := flag.String("dsn", "", "MySql Data source Name")
	flag.Parse()

	// New Info Logger
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate | log.Ltime)
	
	// New Error Logger
	erroLog := log.New(os.Stderr, "ERROR\t", log.Ldate | log.Ltime | log.Lshortfile)

	db, err := openDB(*dsn)
	app := &application{
		errorLog: erroLog,
		infoLog: infoLog,
	}
	

	// Initialize a server struct
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: erroLog,
		Handler: app.routes(),
	}

	infoLog.Printf("starting a server on port %s", *addr)
	err := srv.ListenAndServe()
	erroLog.Fatal(err)
}
