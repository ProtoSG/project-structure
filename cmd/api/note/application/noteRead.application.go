package application

import "struct-proyect/cmd/api/note/domain"

type NoteRead struct {
	repo domain.NoteRepository
}

func NewNoteRead(repo domain.NoteRepository) *NoteRead {
	return &NoteRead{repo}
}

func (this *NoteRead) Execute() ([]*domain.Note, error) {
	return this.repo.Read()
}
