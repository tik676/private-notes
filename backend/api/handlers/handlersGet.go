package handlers

import (
	"encoding/json"
	"net/http"
	"private-notes/internal/db"
	"private-notes/internal/models"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func HandlerMe(w http.ResponseWriter, r *http.Request) {
	userIDRaw := r.Context().Value("user_id")
	userID, ok := userIDRaw.(int)
	if !ok {
		http.Error(w, "user_id не найден", http.StatusUnauthorized)
		return
	}

	notes, err := db.GetWithIDNotesMe(userID)
	if err != nil {
		http.Error(w, "хз какую ошибку вернуть посоветуй", http.StatusBadRequest)
		return
	}

	if notes == nil {
		notes = []models.Notes{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)

}

func GetPublicNoteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	noteID, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Некорректный ID", http.StatusBadRequest)
		return
	}

	note, err := db.GetPublicNote(noteID)
	if err != nil {
		http.Error(w, "Заметка не найдена или не публичная", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func GetNoteByIDAndUserHandler(w http.ResponseWriter, r *http.Request) {
	userIDRAW := r.Context().Value("user_id")
	userID, ok := userIDRAW.(int)
	if !ok {
		http.Error(w, "user_id не найден", http.StatusUnauthorized)
		return
	}

	idStr := chi.URLParam(r, "id")
	noteID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID заметки", http.StatusBadRequest)
		return
	}

	note, err := db.GetNoteByIDAndUser(noteID, userID)
	if err != nil {
		http.Error(w, "Не удалось получить заметку", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func CheckPrivateNoteHandler(w http.ResponseWriter, r *http.Request) {
	noteIDRAW := chi.URLParam(r, "id")
	noteID, err := strconv.Atoi(noteIDRAW)
	if err != nil {
		http.Error(w, "Некорректный ID", http.StatusBadRequest)
		return
	}

	note, err := db.GetNoteByID(noteID)
	if err != nil {
		http.Error(w, "Заметка не найдена", http.StatusNotFound)
		return
	}

	if !note.IsPrivate {
		http.Error(w, "Заметка не является приватной", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func CheckNoteHandler(w http.ResponseWriter, r *http.Request) {
	noteIDStr := chi.URLParam(r, "id")
	noteID, err := strconv.Atoi(noteIDStr)
	if err != nil {
		http.Error(w, "Некорректный ID", http.StatusBadRequest)
		return
	}

	note, err := db.GetNoteByID(noteID)
	if err != nil {
		http.Error(w, "Заметка не найдена", http.StatusNotFound)
		return
	}

	type response struct {
		IsPrivate bool `json:"is_private"`
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response{IsPrivate: note.IsPrivate})
}
