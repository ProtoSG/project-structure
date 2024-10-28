package routes

import (
	"struct-proyect/cmd/api/note/infrastructure/routes"
	"struct-proyect/internal/services"

	"github.com/gorilla/mux"
)

func NewRoutes(serviceContainer *services.ServiceContainer) *mux.Router {
	r := mux.NewRouter()

	routes.NoteRoutes(r, serviceContainer)

	return r
}
