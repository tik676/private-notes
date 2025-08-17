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
		http.Error(w, "user_id not found", http.StatusUnauthorized)
		return
	}

	notes, err := db.GetWithIDNotesMe(userID)
	if err != nil {
		http.Error(w, "Failed to get notes", http.StatusBadRequest)
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
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	note, err := db.GetPublicNote(noteID)
	if err != nil {
		http.Error(w, "Note not found or not public", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func GetNoteByIDAndUserHandler(w http.ResponseWriter, r *http.Request) {
	userIDRaw := r.Context().Value("user_id")
	userID, ok := userIDRaw.(int)
	if !ok {
		http.Error(w, "user_id not found", http.StatusUnauthorized)
		return
	}

	idStr := chi.URLParam(r, "id")
	noteID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	note, err := db.GetNoteByIDAndUser(noteID, userID)
	if err != nil {
		http.Error(w, "Failed to get note", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func CheckPrivateNoteHandler(w http.ResponseWriter, r *http.Request) {
	noteIDRaw := chi.URLParam(r, "id")
	noteID, err := strconv.Atoi(noteIDRaw)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	note, err := db.GetNoteByID(noteID)
	if err != nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	if !note.IsPrivate {
		http.Error(w, "Note is not private", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func CheckNoteHandler(w http.ResponseWriter, r *http.Request) {
	noteIDStr := chi.URLParam(r, "id")
	noteID, err := strconv.Atoi(noteIDStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	note, err := db.GetNoteByID(noteID)
	if err != nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	type response struct {
		IsPrivate bool `json:"is_private"`
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response{IsPrivate: note.IsPrivate})
}
