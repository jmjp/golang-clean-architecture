package handlers

import (
	"encoding/json"
	"net/http"
	"onion/internal/domain/usecases"
)

type LoginHandler struct {
	usecase usecases.LoginUsecase
}

// NewLoginHandler creates a new instance of the LoginHandler struct.
//
// It takes a usecase of type usecases.LoginUsecase as a parameter and returns a pointer to the LoginHandler struct.
func NewLoginHandler(usecase usecases.LoginUsecase) *LoginHandler {
	return &LoginHandler{
		usecase: usecase,
	}
}

type LoginRequest struct {
	Email string `json:"email"`
}

// HandleFunc handles the HTTP request for login.
//
// It takes a http.ResponseWriter and a http.Request as parameters.
func (h *LoginHandler) HandleFunc(w http.ResponseWriter, r *http.Request) {
	body := new(LoginRequest)

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	out, err := h.usecase.Execute(r.Context(), body.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"otp": *out})

}
