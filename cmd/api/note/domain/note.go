package domain

type Note struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Completed bool   `json:"completed"`
}

func NewNote(id int, title, createdAt, updatedAt string) *Note {
	return &Note{
		ID:        id,
		Title:     title,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Completed: false,
	}
}
