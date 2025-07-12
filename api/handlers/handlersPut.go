package handlers

import (
	"encoding/json"
	"net/http"
	"private-notes/api/authorization"
	"private-notes/internal/db"
	"private-notes/internal/models"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func UpdateNoteHandler(w http.ResponseWriter, r *http.Request) {
	var input models.UpdateNote

	noteID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Некорректный ID", http.StatusBadRequest)
		return
	}

	userRawID := r.Context().Value("user_id")
	userID, ok := userRawID.(int)
	if !ok {
		http.Error(w, "user_id не найден", http.StatusUnauthorized)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Ошибка парсинга JSON", http.StatusBadRequest)
		return
	}

	note, err := db.GetNoteByIDAndUser(noteID, userID)
	if err != nil {
		http.Error(w, "Заметка не найдена", http.StatusNotFound)
		return
	}

	if input.Content != nil {
		note.Content = *input.Content
	}
	if input.ExpiresAt != nil {
		note.ExpiresAt = *input.ExpiresAt
	}
	if input.IsPrivate != nil {
		note.IsPrivate = *input.IsPrivate
	}

	var hashPass *string

	if note.IsPrivate {
		if input.Password == nil {
			// если приватная, но пароль не передан — сохраняем старый (если есть)
			hashPass = note.HashPassword
		} else {
			if *input.Password == "" {
				http.Error(w, "Пароль не может быть пустым", http.StatusBadRequest)
				return
			}
			h, err := authorization.GenerateHash(*input.Password)
			if err != nil {
				http.Error(w, "Ошибка хеширования пароля", http.StatusInternalServerError)
				return
			}
			hashPass = &h
		}
	} else {
		// если заметка становится публичной — убираем пароль
		hashPass = nil
	}

	err = db.UpdateNote(noteID, userID, note.Content, note.ExpiresAt, note.IsPrivate, hashPass)
	if err != nil {
		http.Error(w, "Не удалось обновить заметку", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "successful",
	})
}
