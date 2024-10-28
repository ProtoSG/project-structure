package infrastructure

import "struct-proyect/cmd/api/note/domain"

type InMemoryNoteRepository struct {
	notes []*domain.Note
}

func NewInMemoryNoteRepository() *InMemoryNoteRepository {
	return &InMemoryNoteRepository{notes: make([]*domain.Note, 0)}
}

func (this *InMemoryNoteRepository) Create(note *domain.Note) error {
	this.notes = append(this.notes, note)
	return nil
}

func (this *InMemoryNoteRepository) Read() ([]*domain.Note, error) {
	return this.notes, nil
}

func (this *InMemoryNoteRepository) ReadByID(id int) (*domain.Note, error) {
	for _, note := range this.notes {
		if note.ID == id {
			return note, nil
		}
	}

	return nil, nil
}

func (this *InMemoryNoteRepository) Update(note *domain.Note) error {
	for i, n := range this.notes {
		if n.ID == note.ID {
			this.notes[i] = note
			return nil
		}
	}

	return nil
}

func (this *InMemoryNoteRepository) Delete(id int) error {
	var newNotes []*domain.Note

	for _, note := range this.notes {
		if note.ID != id {
			newNotes = append(newNotes, note)
		}
	}

	this.notes = newNotes
	return nil
}
