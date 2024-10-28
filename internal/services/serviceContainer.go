package services

import (
	"database/sql"
	"struct-proyect/cmd/api/note/application"
)

type ServiceContainer struct {
	Note *application.NoteService
}

func NewServiceContainer(db *sql.DB) *ServiceContainer {
	return &ServiceContainer{
		Note: application.NewNoteService(db),
	}
}
