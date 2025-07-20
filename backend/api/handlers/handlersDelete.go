package handlers

import (
	"encoding/json"
	"net/http"
	"private-notes/internal/db"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	userRAWID := r.Context().Value("user_id")
	userID, ok := userRAWID.(int)
	if !ok {
		http.Error(w, "user_id не найден", http.StatusUnauthorized)
		return
	}

	strNoteID := chi.URLParam(r, "id")
	noteID, err := strconv.Atoi(strNoteID)
	if err != nil {
		http.Error(w, "Хз какую ошибку вернуть посоветуй", http.StatusUnauthorized)
		return
	}

	if err := db.DeleteNote(noteID, userID); err != nil {
		http.Error(w, "Не удалось удалить заметку", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "successful",
	})

}
