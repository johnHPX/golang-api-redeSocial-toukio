package resource

import (
	"API-RS-TOUKIO/internal/appl"
	"API-RS-TOUKIO/internal/domain/users"
	"API-RS-TOUKIO/internal/infra/data/response"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type createUserRequest struct {
	Name     string `json:"name"`
	Nick     string `json:"nick"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type createUserResponse struct {
	MID string `json:"_mid"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var request createUserRequest
	err = json.Unmarshal(bodyRequest, &request)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	svc := appl.NewUserService()
	ent := &users.Entity{
		Name:     request.Name,
		Nick:     request.Nick,
		Email:    request.Email,
		Password: request.Password,
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

type listAllUsersRequest struct {
	MID string `json:"_mid"`
}

type listAllUsersResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Nick     string `json:"nick"`
	Email    string `json:"email"`
	Password string `json:"password"`
	MID      string `json:"_mid"`
}

func ListAllUsers(w http.ResponseWriter, r *http.Request) {

	request := listAllUsersRequest{
		MID: "ok",
	}

	svc := appl.NewUserService()
	list, err := svc.ListALLUser()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	result := make([]listAllUsersResponse, 0)
	for _, v := range list {
		result = append(result, listAllUsersResponse{
			ID:       v.ID,
			Name:     v.Name,
			Nick:     v.Nick,
			Email:    v.Email,
			Password: v.Password,
			MID:      request.MID,
		})
	}

	response.JSON(w, http.StatusAccepted, result)
}

type listByNameOrNickUsersRequest struct {
	Name string `json:"-"`
	Nick string `json:"-"`
	MID  string `json:"_mid"`
}

type listByNameOrNickUsersResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Nick     string `json:"nick"`
	Email    string `json:"email"`
	Password string `json:"password"`
	MID      string `json:"_mid"`
}

func ListByNameOrNickUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	request := listByNameOrNickUsersRequest{
		Name: nameOrNick,
		Nick: nameOrNick,
		MID:  "ok",
	}

	svc := appl.NewUserService()
	list, err := svc.ListByNameOrNickUsers(request.Name)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	result := make([]listByNameOrNickUsersResponse, 0)
	for _, v := range list {
		result = append(result, listByNameOrNickUsersResponse{
			ID:       v.ID,
			Name:     v.Name,
			Nick:     v.Nick,
			Email:    v.Email,
			Password: v.Password,
			MID:      request.MID,
		})
	}

	response.JSON(w, http.StatusAccepted, result)
}

type findUsersRequest struct {
	ID int64 `json:"-"`
}

type findUsersResponse struct {
	Name     string `json:"name"`
	Nick     string `json:"nick"`
	Email    string `json:"email"`
	Password string `json:"password"`
	MID      string `json:"_mid"`
}

func FindUsers(w http.ResponseWriter, r *http.Request) {
	paraments := mux.Vars(r)
	userID, err := strconv.ParseInt(paraments["userId"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	request := findUsersRequest{
		ID: userID,
	}

	svc := appl.NewUserService()
	user, err := svc.FindUser(request.ID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	u := findUsersResponse{
		Name:     user.Name,
		Nick:     user.Nick,
		Email:    user.Email,
		Password: user.Password,
		MID:      "ok",
	}

	response.JSON(w, http.StatusAccepted, u)
}

type updateUserRequest struct {
	Name  string `json:"name"`
	Nick  string `json:"nick"`
	Email string `json:"email"`
}

type updateUserResponse struct {
	MID string `json:"_mid"`
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	paraments := mux.Vars(r)
	userID, err := strconv.ParseInt(paraments["userId"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var request updateUserRequest
	err = json.Unmarshal(bodyRequest, &request)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	svc := appl.NewUserService()
	ent := &users.Entity{
		ID:    userID,
		Name:  request.Name,
		Nick:  request.Nick,
		Email: request.Email,
	}

	err = svc.UpdateUser(ent)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	resp := &createUserResponse{
		MID: "ok",
	}

	response.JSON(w, http.StatusAccepted, resp)
}
