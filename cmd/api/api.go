package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"struct-proyect/internal/routes"
	"struct-proyect/internal/services"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (this *APIServer) Run() error {
	serviceContainer := services.NewServiceContainer(this.db)
	router := routes.NewRoutes(serviceContainer)

	router.HandleFunc("/", this.checkHandler)

	log.Println("Listening on ", this.addr)

	svr := &http.Server{
		Addr:    this.addr,
		Handler: router,
	}

	if err := svr.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (this *APIServer) checkHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World")
}
