package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/098765432m/ecom/services/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()

	// Prefix to server /api/v1
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	// User Router
	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subRouter)

	log.Println("Server is running on port", s.addr)

	initStorage(s.db)

	return http.ListenAndServe(s.addr, router)
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
