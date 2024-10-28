package application

import (
	"struct-proyect/cmd/api/note/domain"
	"struct-proyect/internal/utils"
)

type NoteReadByID struct {
	repo domain.NoteRepository
}

func NewNoteReadByID(repo domain.NoteRepository) *NoteReadByID {
	return &NoteReadByID{repo}
}

func (this *NoteReadByID) Execute(id int) (*domain.Note, error) {
	note, _ := this.repo.ReadByID(id)
	if note == nil {
		return nil, utils.NewEntityNotFound(id, "Nota")
	}

	return note, nil
}
