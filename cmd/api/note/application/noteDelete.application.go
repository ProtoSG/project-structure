package application

import (
	"struct-proyect/cmd/api/note/domain"
	"struct-proyect/internal/utils"
)

type NoteDelete struct {
	repo domain.NoteRepository
}

func NewNoteDelete(repo domain.NoteRepository) *NoteDelete {
	return &NoteDelete{repo}
}

func (this *NoteDelete) Execute(id int) error {
	if note, _ := this.repo.ReadByID(id); note == nil {
		return utils.NewEntityNotFound(id, "Nota")
	}

	return this.repo.Delete(id)
}
