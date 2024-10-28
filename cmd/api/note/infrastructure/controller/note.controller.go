package controller

import (
	"net/http"
	"struct-proyect/cmd/api/note/domain"
	"struct-proyect/internal/services"
	"struct-proyect/internal/utils"
)

type NoteController struct {
	serviceContainer *services.ServiceContainer
}

func NewNoteController(serviceContainer *services.ServiceContainer) *NoteController {
	return &NoteController{serviceContainer}
}

func (this *NoteController) Create(w http.ResponseWriter, r *http.Request) {
	note := domain.Note{}
	if err := utils.ParseJSON(r, &note); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := this.serviceContainer.Note.Create.Execute(note.ID, note.Title, note.CreatedAt, note.UpdatedAt)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "Nota creada"})
}

func (this *NoteController) Read(w http.ResponseWriter, r *http.Request) {
	notes, err := this.serviceContainer.Note.Read.Execute()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, nil)
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, notes); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

func (this *NoteController) ReadByID(w http.ResponseWriter, r *http.Request) {}
func (this *NoteController) Update(w http.ResponseWriter, r *http.Request)   {}
func (this *NoteController) Delete(w http.ResponseWriter, r *http.Request)   {}
