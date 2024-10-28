package application

import (
	"struct-proyect/cmd/api/note/domain"
	"struct-proyect/internal/utils"
)

type NoteUpdate struct {
	repo domain.NoteRepository
}

func NewNoteUpdate(repo domain.NoteRepository) *NoteUpdate {
	return &NoteUpdate{repo}
}

func (this *NoteUpdate) Execute(id int, title, createdAt, updateAt string) error {
	if note, _ := this.repo.ReadByID(id); note == nil {
		return utils.NewEntityNotFound(id, "Nota")
	}

	note := domain.NewNote(id, title, createdAt, updateAt)
	return this.repo.Update(note)
}
