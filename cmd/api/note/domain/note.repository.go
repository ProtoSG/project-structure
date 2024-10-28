package domain

type NoteRepository interface {
	Create(*Note) error
	Read() ([]*Note, error)
	ReadByID(id int) (*Note, error)
	Update(*Note) error
	Delete(id int) error
}
