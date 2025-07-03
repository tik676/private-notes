package handlers

import (
	"encoding/json"
	"net/http"
	"private-notes/api/authorization"
	"private-notes/internal/models"
)

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var regUser models.RegisterInput

	if err := json.NewDecoder(r.Body).Decode(&regUser); err != nil {
		http.Error(w, "Не удалось распарсить JSON", http.StatusBadRequest)
		return
	}

	if regUser.Name == "" || regUser.Password == "" {
		http.Error(w, "Имя и пароль обязательны", http.StatusBadRequest)
		return
	}

	if err := authorization.RegisterUser(regUser.Name, regUser.Password); err != nil {
		http.Error(w, "Не удалось зарегестрировать пользователя", http.StatusBadRequest)
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
		http.Error(w, "Не удалось распарсить JSON", http.StatusBadRequest)
		return
	}

	res, err := authorization.LoginUser(user.Name, user.Password)
	if err != nil {
		http.Error(w, "Че то с тобой не то", http.StatusBadRequest)
		return
	}

	token, err := authorization.GenerateJWT(res)
	if err != nil {
		http.Error(w, "Че то с тобой не то", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})

}
