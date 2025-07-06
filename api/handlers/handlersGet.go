package handlers

import (
	"encoding/json"
	"net/http"
	"private-notes/internal/db"
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)

}
