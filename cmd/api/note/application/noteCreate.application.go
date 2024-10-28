package application

import "struct-proyect/cmd/api/note/domain"

type NoteCreate struct {
	repo domain.NoteRepository
}

func NewNoteCreate(repo domain.NoteRepository) *NoteCreate {
	return &NoteCreate{repo: repo}
}

func (this *NoteCreate) Execute(id int, title, createdAt, updatedAt string) error {
	note := domain.NewNote(id, title, createdAt, updatedAt)

	return this.repo.Create(note)
}
