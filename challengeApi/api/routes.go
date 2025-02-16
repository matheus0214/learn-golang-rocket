package api

import (
	"challengeApi/api/dto"
	"challengeApi/database"
	"challengeApi/domain"
	"challengeApi/utils"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(r chi.Router, db database.DB) {
	r.Get("/users", getUsers(db))
	r.Get("/users/{id}", getUserById(db))
	r.Post("/users", createUser(db))
	r.Put("/users/{id}", updateUser(db))
	r.Delete("/users/{id}", deleteUserById(db))
}

func getUsers(db database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data []domain.User

		for _, u := range db.FindAll() {
			data = append(data, u)
		}

		utils.JsonResponse(w, utils.Response{Data: data}, http.StatusOK)
	}
}

func getUserById(db database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		user, err := db.FindById(id)
		if err != nil {
			utils.JsonResponse(w, utils.Response{Error: err.Error()}, http.StatusNotFound)
			return
		}

		utils.JsonResponse(w, utils.Response{Data: user}, http.StatusOK)
	}
}

func deleteUserById(db database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if err := db.Delete(id); err != nil {
			utils.JsonResponse(w, utils.Response{Error: err.Error()}, http.StatusNotFound)
			return
		}

		utils.JsonResponse(w, utils.Response{Data: nil}, http.StatusNoContent)
	}
}

func createUser(db database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		if err := db.Insert(user); err != nil {
			utils.JsonResponse(w, utils.Response{Error: err.Error()}, http.StatusInternalServerError)
			return
		}

		utils.JsonResponse(w, utils.Response{Data: dto.CreatedUserOutputDTO{Message: "user created", ID: user.ID.String()}}, http.StatusCreated)
	}
}

func updateUser(db database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		var input dto.UpdateUserInputDTO

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			utils.JsonResponse(w, utils.Response{Error: "error to read data"}, http.StatusUnprocessableEntity)
			return
		}

		user, err := db.FindById(id)
		if err != nil {
			utils.JsonResponse(w, utils.Response{Error: err.Error()}, http.StatusNotFound)
			return
		}

		user.FirstName = input.FirstName
		user.LastName = input.LastName
		user.Biography = input.Biography

		if err := db.Update(user); err != nil {
			utils.JsonResponse(w, utils.Response{Error: err.Error()}, http.StatusInternalServerError)
			return
		}

		utils.JsonResponse(w, utils.Response{Data: dto.UpdatedUserOutputDTO{Message: "user updated", ID: user.ID.String()}}, http.StatusOK)
	}
}
