package main

import (
	"database/sql"
	"flag"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	_ "github.com/mattn/go-sqlite3"
	"goweb/internal/models"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")

	db, err := sql.Open("sqlite3", "./goweb.db")
	if err != nil {
		log.Fatal(err)

	}
	defer db.Close()
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	templateCache, err := newTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	formDecoder := form.NewDecoder()

	sessionManager := scs.New()

	sessionManager.Lifetime = 12 * time.Hour

	sessionManager.Cookie.Secure = true

	app := &application{
		errorLog:       errorLog,
		infoLog:        infoLog,
		posts:          &models.PostModel{DB: db},
		users:          &models.UserModel{DB: db},
		comments:       &models.CommentModel{DB: db},
		likes:          &models.LikesModel{DB: db},
		templateCache:  templateCache,
		formDecoder:    formDecoder,
		sessionManager: sessionManager,
	}

	srv := &http.Server{Addr: *addr,

		// Call the new app.routes() method to get the servemux containing our routes.
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
