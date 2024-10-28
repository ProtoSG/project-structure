package application

import (
	"database/sql"
	"struct-proyect/cmd/api/note/domain"
	"struct-proyect/cmd/api/note/infrastructure"
)

type NoteService struct {
	Create   NoteCreate
	Read     NoteRead
	ReadByID NoteReadByID
	Update   NoteUpdate
	Delete   NoteDelete
}

func NewNoteService(db *sql.DB) *NoteService {
	var repo domain.NoteRepository = infrastructure.NewInMemoryNoteRepository()

	return &NoteService{
		Create:   *NewNoteCreate(repo),
		Read:     *NewNoteRead(repo),
		ReadByID: *NewNoteReadByID(repo),
		Update:   *NewNoteUpdate(repo),
		Delete:   *NewNoteDelete(repo),
	}
}
