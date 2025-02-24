package api

import (
	"errors"
	"gobid/internal/jsonutils"
	"gobid/internal/services"
	"gobid/internal/usecases/user"
	"net/http"
)

func (a *Api) handleSignupUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJSON[user.CreateUserReq](r)
	if err != nil {
		_ = jsonutils.EncodeJSON(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	id, err := a.UserService.Create(r.Context(), data.UserName, data.Email, data.Password, data.Bio)
	if err != nil {
		if errors.Is(err, services.ErrDuplicatedUsernameOrEmail) {
			_ = jsonutils.EncodeJSON(w, r, http.StatusConflict, map[string]string{"error": err.Error()})
			return
		}
		return
	}

	_ = jsonutils.EncodeJSON(w, r, http.StatusConflict, map[string]string{"id": id.String()})
}

func (a *Api) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJSON[user.LoginUserReq](r)
	if err != nil {
		_ = jsonutils.EncodeJSON(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	id, err := a.UserService.AuthenticateUser(r.Context(), data.Email, data.Password)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			_ = jsonutils.EncodeJSON(w, r, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		_ = jsonutils.EncodeJSON(w, r, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		return
	}

	err = a.Sessions.RenewToken(r.Context())
	if err != nil {
		_ = jsonutils.EncodeJSON(w, r, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		return
	}

	a.Sessions.Put(r.Context(), "AuthenticatedUserId", id)

	_ = jsonutils.EncodeJSON(w, r, http.StatusOK, map[string]string{"message": "logged in successfully"})
}

func (a *Api) handleLogoutUser(w http.ResponseWriter, r *http.Request) {
	err := a.Sessions.RenewToken(r.Context())
	if err != nil {
		_ = jsonutils.EncodeJSON(w, r, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		return
	}

	a.Sessions.Remove(r.Context(), "AuthenticatedUserId")

	_ = jsonutils.EncodeJSON(w, r, http.StatusOK, map[string]string{"message": "logged out successfully"})
}
