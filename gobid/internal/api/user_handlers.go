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
	panic("Need implementation")
}

func (a *Api) handleLogoutUser(w http.ResponseWriter, r *http.Request) {
	panic("Need implementation")
}
