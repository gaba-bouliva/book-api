package api

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type APIServer struct {
	Addr 			string
	DB 				*sql.DB
	InfoLog 	*log.Logger
	ErrorLog	*log.Logger
}

func NewServer(db *sql.DB, Addr string) *APIServer{
	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime|log.Lshortfile)

	return &APIServer{
		InfoLog: infoLog,
		ErrorLog: errorLog,
		Addr: Addr,
		DB: db,
	}
}

func (api *APIServer) Run() error {
	server := http.Server{
		Addr: api.Addr,
		Handler: api.routes(),	
	}
	return server.ListenAndServe()
}

func (api *APIServer) routes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"https://*", "http://*"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: true,
    MaxAge:           300, 
  }))

	apiRouter 	:= chi.NewRouter()

	r.Mount("/api", apiRouter)

	return apiRouter	
}