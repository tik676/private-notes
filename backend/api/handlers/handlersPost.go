package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"private-notes/api/authorization"
	"private-notes/internal/db"
	"private-notes/internal/models"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var regUser models.RegisterInput

	if err := json.NewDecoder(r.Body).Decode(&regUser); err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	if regUser.Name == "" || regUser.Password == "" {
		http.Error(w, "Name and password are required", http.StatusBadRequest)
		return
	}

	if err := authorization.RegisterUser(regUser.Name, regUser.Password); err != nil {
		http.Error(w, "Failed to register user", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "ok",
	})
}

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.LoginInput

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	userID, err := authorization.LoginUser(user.Name, user.Password)
	if err != nil {
		http.Error(w, "Something is wrong with you", http.StatusBadRequest)
		return
	}

	jwtMaker := authorization.NewJWTMaker(os.Getenv("JWT_SECRET"))

	token, err := jwtMaker.CreateToken(userID, 15*time.Minute)
	if err != nil {
		http.Error(w, "Something’s not you", http.StatusBadRequest)
		return
	}

	refreshToken, refreshExp, err := authorization.GenerateRefresh()
	if err != nil {
		http.Error(w, "Failed to create refresh token", http.StatusInternalServerError)
		return
	}

	if err := db.SaveRefreshToken(userID, refreshToken, refreshExp); err != nil {
		http.Error(w, "Failed to save refresh token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"access_token":  token,
		"refresh_token": refreshToken,
		"user": map[string]interface{}{
			"id":   userID,
			"name": user.Name,
		},
	})
}

func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	userIDRaw := r.Context().Value("user_id")
	userIDint, ok := userIDRaw.(int)
	if !ok {
		http.Error(w, "user_id not found", http.StatusUnauthorized)
		return
	}

	var note models.NoteInput
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	var hashPass *string

	if note.IsPrivate {
		if note.Password == nil || strings.TrimSpace(*note.Password) == "" {
			http.Error(w, "Private note requires a password", http.StatusBadRequest)
			return
		}
		h, err := authorization.GenerateHash(*note.Password)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusBadRequest)
			return
		}
		hashPass = &h
	} else {
		if note.Password != nil {
			http.Error(w, "Public note should not have a password", http.StatusBadRequest)
			return
		}
	}

	err := db.CreateNote(userIDint, note.Content, note.ExpiresAt, note.IsPrivate, hashPass)
	if err != nil {
		http.Error(w, "Failed to add note", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "success",
	})
}

func RefreshTokenHandle(w http.ResponseWriter, r *http.Request) {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	userID, err := db.GetUserIDByRefreshToken(req.RefreshToken)
	if err != nil {
		http.Error(w, "Refresh token expired or does not exist", http.StatusUnauthorized)
		return
	}

	jwtMaker := authorization.NewJWTMaker(os.Getenv("JWT_SECRET"))

	token, err := jwtMaker.CreateToken(userID, 15*time.Minute)
	if err != nil {
		http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
		return
	}

	newRefresh, refreshExp, err := authorization.GenerateRefresh()
	if err != nil {
		http.Error(w, "Failed to generate refresh token", http.StatusInternalServerError)
		return
	}
	if err := db.SaveRefreshToken(userID, newRefresh, refreshExp); err != nil {
		http.Error(w, "Failed to save refresh token", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  token,
		"refresh_token": newRefresh,
	})
}

func UnlockPrivateNoteHandler(w http.ResponseWriter, r *http.Request) {
	type reqS struct {
		Password string `json:"password"`
	}

	var req reqS

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	noteIDRAW := chi.URLParam(r, "id")
	noteID, err := strconv.Atoi(noteIDRAW)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	note, err := db.GetNoteByID(noteID)
	if err != nil {
		http.Error(w, "Note not found or incorrect password", http.StatusBadRequest)
		return
	}

	if !note.IsPrivate {
		http.Error(w, "This note is not private", http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(*note.HashPassword), []byte(req.Password))
	if err != nil {
		http.Error(w, "Incorrect password", http.StatusUnauthorized)
		return
	}

	note.HashPassword = nil

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	if err := db.DeleteRefreshToken(req.RefreshToken); err != nil {
		http.Error(w, "Failed to delete refresh token", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "logout successful",
	})
}
