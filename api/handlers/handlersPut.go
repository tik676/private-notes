package handlers

import (
	"encoding/json"
	"net/http"
	"private-notes/internal/db"
	"private-notes/internal/models"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func UpdateNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note models.UpdateNote
	strNoteID := chi.URLParam(r, "id")
	noteID, err := strconv.Atoi(strNoteID)
	if err != nil {
		http.Error(w, "Хз какую ошибку вернуть посоветуй", http.StatusUnauthorized)
		return
	}

	userRAWID := r.Context().Value("user_id")
	userID, ok := userRAWID.(int)
	if !ok {
		http.Error(w, "user_id не найден", http.StatusUnauthorized)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "Не удалось распарсить JSON", http.StatusBadRequest)
		return
	}

	noteSource, err := db.GetNoteByIDAndUser(noteID, userID)
	if err != nil {
		http.Error(w, "Заметка не найдена", http.StatusNotFound)
		return
	}

	if note.Content != nil {
		noteSource.Content = *note.Content
	}
	if note.ExpiresAt != nil {
		noteSource.ExpiresAt = *note.ExpiresAt
	}
	if note.IsPrivate != nil {
		noteSource.IsPrivate = *note.IsPrivate
	}

	err = db.UpdateNote(noteID, userID, noteSource.Content, noteSource.ExpiresAt, noteSource.IsPrivate)
	if err != nil {
		http.Error(w, "Не удалось обновить заметку", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "successful",
	})
}
