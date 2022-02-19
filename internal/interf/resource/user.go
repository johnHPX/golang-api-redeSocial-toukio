package resource

import (
	"API-RS-TOUKIO/internal/appl"
	"API-RS-TOUKIO/internal/domain/users"
	"API-RS-TOUKIO/internal/infra/data/response"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type createUserRequest struct {
	Name     string `json:"name"`
	Nick     string `json:"nick"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type createUserResponse struct {
	MID string `json:"mid"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var entUserRequest createUserRequest
	err = json.Unmarshal(bodyRequest, &entUserRequest)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	svc := appl.NewUserService()
	ent := &users.Entity{
		Name:     entUserRequest.Name,
		Nick:     entUserRequest.Nick,
		Email:    entUserRequest.Email,
		Password: entUserRequest.Password,
	}

	err = svc.CreateUser(ent)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	entUserResponse := &createUserResponse{
		MID: "ok",
	}

	response.JSON(w, http.StatusAccepted, entUserResponse)
}
