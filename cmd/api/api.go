package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Winterson-Islary/jwt-golang.git/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(a_addr string, a_db *sql.DB) *APIServer {
	return &APIServer{
		addr: a_addr,
		db:   a_db,
	}
}

func (server *APIServer) Run() error {
	router := mux.NewRouter()
	sub_router := router.PathPrefix("/api/v1").Subrouter()
	// USER Handlers
	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(sub_router)
	log.Println("Listening On: ", server.addr)
	return http.ListenAndServe(server.addr, router)
}
