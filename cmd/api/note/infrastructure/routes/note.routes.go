package routes

import (
	"struct-proyect/cmd/api/note/infrastructure/controller"
	"struct-proyect/internal/services"

	"github.com/gorilla/mux"
)

func NoteRoutes(r *mux.Router, serviceContainer *services.ServiceContainer) {
	controller := controller.NewNoteController(serviceContainer)

	r.HandleFunc("/note", controller.Create).Methods("POST")
	r.HandleFunc("/note", controller.Read).Methods("GET")
	r.HandleFunc("/note/{id}", controller.ReadByID).Methods("GET")
	r.HandleFunc("/note/{id}", controller.Update).Methods("PUT")
	r.HandleFunc("/note/{id}", controller.Delete).Methods("DELETE")
}
