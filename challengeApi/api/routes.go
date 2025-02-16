package api

import (
	"challengeApi/api/dto"
	"challengeApi/domain"
	"challengeApi/utils"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(r chi.Router) {
	r.Get("/users", getUsers)
	r.Post("/users", createUser)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I'm alive"))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateUserInputDTO

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.JsonResponse(w, utils.Response{Error: "error to read data"}, http.StatusUnprocessableEntity)
		return
	}

	user, err := domain.NewUser(input.FirstName, input.LastName, input.Biography)
	if err != nil {
		utils.JsonResponse(w, utils.Response{Error: err.Error()}, http.StatusCreated)
		return
	}

	utils.JsonResponse(w, utils.Response{Data: dto.CreatedUserOutputDTO{Message: "user created", ID: user.ID.String()}}, http.StatusCreated)
}
