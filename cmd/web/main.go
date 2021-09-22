package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alex-orkuma/foodexp/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golangcollege/sessions"
)

// A Custom application struct for dependency injection
type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	session       *sessions.Session
	products      *mysql.ProductModel
	templateCache map[string]*template.Template
}

func main() {

	// Go server address flag
	addr := flag.String("addr", ":4000", "HTTP Network Address")

	// MySql Data source name
	dsn := flag.String("dsn", "foodexp:Or$kumashnd417@/foodexp?parseTime=true", "MySql Data source Name")

	// Session secret key flag
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "secret key")
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

	templateCache, err := newTemplateCache("./ui/html/")

	if err != nil {
		erroLog.Fatal(err)
	}

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour

	app := &application{
		errorLog:      erroLog,
		infoLog:       infoLog,
		session:       session,
		products:      &mysql.ProductModel{DB: *db},
		templateCache: templateCache,
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
