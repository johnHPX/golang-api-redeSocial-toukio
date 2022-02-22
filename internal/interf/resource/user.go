package resource

import (
	"API-RS-TOUKIO/internal/appl"
	"API-RS-TOUKIO/internal/domain/users"
	"API-RS-TOUKIO/internal/infra/data/response"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

/*
============================
== CREATE USER =============
============================
*/

type createUserRequest struct {
	Name     string `json:"name"`
	Nick     string `json:"nick"`
	Email    string `json:"email"`
	Password string `json:"password"`
	MID      string `json:"mid"`
}

type createUserResponse struct {
	MID string `json:"mid"`
}

// cria um usuario
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
		MID: request.MID,
	}

	response.JSON(w, http.StatusAccepted, entUserResponse)
}

/*
============================
== LISTALL USER ============
============================
*/

type listAllUsersRequest struct {
	MID string `json:"-"`
}

type listAllUsersResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Nick     string `json:"nick"`
	Email    string `json:"email"`
	Password string `json:"password"`
	MID      string `json:"mid"`
}

func ListAllUsers(w http.ResponseWriter, r *http.Request) {
	mid := strings.ToLower(r.URL.Query().Get("mid"))

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
			MID:      mid,
		})
	}

	response.JSON(w, http.StatusAccepted, result)
}

/*
============================
== LISTBYNAMEORNICK USER ===
============================
*/

type listByNameOrNickUsersRequest struct {
	NameOrNick string `json:"-"`
	MID        string `json:"-"`
}

type listByNameOrNickUsersResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Nick  string `json:"nick"`
	Email string `json:"email"`
	MID   string `json:"mid"`
}

// lista todos os usuarios pelo nome ou nick
func ListByNameOrNickUsers(w http.ResponseWriter, r *http.Request) {
	mid := strings.ToLower(r.URL.Query().Get("mid"))
	var request listByNameOrNickUsersRequest
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	request.NameOrNick = nameOrNick

	svc := appl.NewUserService()
	list, err := svc.ListByNameOrNickUsers(request.NameOrNick)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	result := make([]listByNameOrNickUsersResponse, 0)
	for _, v := range list {
		result = append(result, listByNameOrNickUsersResponse{
			ID:    v.ID,
			Name:  v.Name,
			Nick:  v.Nick,
			Email: v.Email,
			MID:   mid,
		})
	}

	response.JSON(w, http.StatusAccepted, result)
}

/*
============================
== FIND USER =============
============================
*/

type findUsersRequest struct {
	ID  int64  `json:"-"`
	MID string `json:"-"`
}

type findUsersResponse struct {
	Name  string `json:"name"`
	Nick  string `json:"nick"`
	Email string `json:"email"`
	MID   string `json:"mid"`
}

// traz um usuario atraves do id
func FindUsers(w http.ResponseWriter, r *http.Request) {
	mid := strings.ToLower(r.URL.Query().Get("mid"))
	paraments := mux.Vars(r)
	userID, err := strconv.ParseInt(paraments["userId"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	svc := appl.NewUserService()
	user, err := svc.FindUser(userID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	fmt.Println(user)

	resp := &findUsersResponse{
		Name:  user.Name,
		Nick:  user.Nick,
		Email: user.Email,
		MID:   mid,
	}

	response.JSON(w, http.StatusAccepted, resp)
}

/*
============================
== UPDATE USER =============
============================
*/

type updateUserRequest struct {
	Name  string `json:"name"`
	Nick  string `json:"nick"`
	Email string `json:"email"`
	MID   string `json:"mid"`
}

type updateUserResponse struct {
	MID string `json:"mid"`
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
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

	paraments := mux.Vars(r)
	userID, err := strconv.ParseInt(paraments["userId"], 10, 64)
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

/*
============================
== DELETE USER =============
============================
*/

type deleteUserRequest struct {
	ID  int64  `json:"-"`
	MID string `json:"-"`
}

type deleteUserResponse struct {
	MID string `json:"mid"`
}

func DeletarUser(w http.ResponseWriter, r *http.Request) {
	mid := strings.ToLower(r.URL.Query().Get("mid"))

	paraments := mux.Vars(r)
	usuarioId, err := strconv.ParseInt(paraments["userId"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	svc := appl.NewUserService()
	err = svc.DeleteUser(usuarioId)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	resp := &deleteUserResponse{
		MID: mid,
	}

	response.JSON(w, http.StatusOK, resp)
}

/*
============================
== LOGIN USER ==============
============================
*/

type loginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginUserResponse struct {
}

// faz login com um usuario cadastrado
func LoginUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user loginUserRequest
	err = json.Unmarshal(bodyRequest, &user)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	svc := appl.NewUserService()
	userSalveBase, err := svc.SearchforEmail(user.Email)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	err = appl.CheckPassword(userSalveBase.Password, user.Password)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, err := appl.CreateToken(uint64(userSalveBase.ID))
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	// resp := &loginUserResponse{
	// 	Token: []byte(token),
	// }

	w.Write([]byte(token))

}

/*
============================
== SEGUIR USER =============
============================
*/

type seguirUserRequest struct {
	MID string `json:"mid"`
}

type seguirUserResponse struct {
	MID string `json:"mid"`
}

// uma vez logado, pode seguir um outro usuario
func SeguirUser(w http.ResponseWriter, r *http.Request) {

	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user seguirUserRequest
	err = json.Unmarshal(bodyRequest, &user)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	paramentros := mux.Vars(r)
	userID, err := strconv.ParseInt(paramentros["userId"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	follewerID, err := appl.ExtractUsuarioID(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if follewerID == userID {
		response.Erro(w, http.StatusForbidden, errors.New("não é posivel seguir você mesmo"))
		return
	}

	svc := appl.NewUserService()
	err = svc.FollowUser(userID, follewerID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	resp := &seguirUserResponse{
		MID: "ok",
	}

	response.JSON(w, http.StatusOK, resp.MID)
}

/*
============================
== PARARSEGUIR USER ========
============================
*/

type pararSeguirUserRequest struct {
	MID string `json:"mid"`
}

type pararSeguirUserUserResponse struct {
	MID string `json:"mid"`
}

// uma vez logado, serve para parar de seguir um usuario
func PararSeguirUser(w http.ResponseWriter, r *http.Request) {

	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user pararSeguirUserRequest
	err = json.Unmarshal(bodyRequest, &user)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	paramentros := mux.Vars(r)
	userID, err := strconv.ParseInt(paramentros["userId"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	follewerID, err := appl.ExtractUsuarioID(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if follewerID == userID {
		response.Erro(w, http.StatusForbidden, errors.New("não é posivel parar de seguir você mesmo"))
		return
	}

	svc := appl.NewUserService()
	err = svc.StopFollowing(userID, follewerID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	resp := &pararSeguirUserUserResponse{
		MID: "ok",
	}

	response.JSON(w, http.StatusOK, resp.MID)
}

/*
============================
== LISTSEGUIDOR USER =======
============================
*/

type listSeguidoresUserRequest struct {
	MID string `json:"-"`
}

type listSeguidoresUserResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Nick      string    `json:"nick"`
	Email     string    `json:"email"`
	Create_at time.Time `json:"create_at"`
	MID       string    `json:"mid"`
}

// uma vez logado, lista todos os seguidores de um usuario
func ListSeguidoresUser(w http.ResponseWriter, r *http.Request) {
	mid := strings.ToLower(r.URL.Query().Get("mid"))
	paraments := mux.Vars(r)
	userID, err := strconv.ParseInt(paraments["userId"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	svc := appl.NewUserService()
	followers, err := svc.SearchFollowers(userID)

	result := make([]listSeguidoresUserResponse, 0)
	for _, v := range followers {
		result = append(result, listSeguidoresUserResponse{
			ID:        v.ID,
			Name:      v.Name,
			Nick:      v.Nick,
			Email:     v.Email,
			Create_at: v.Create_at,
			MID:       mid,
		})
	}

	response.JSON(w, http.StatusAccepted, result)
}

/*
============================
== LISTSEGUINDO USER =======
============================
*/

type listSeguindoUserRequest struct {
	MID string `json:"-"`
}

type listSeguindoUserResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Nick      string    `json:"nick"`
	Email     string    `json:"email"`
	Create_at time.Time `json:"create_at"`
	MID       string    `json:"mid"`
}

// uma vez logado, lista todas as usuarios que o usuario está seguindo
func ListSeguindoUser(w http.ResponseWriter, r *http.Request) {
	mid := strings.ToLower(r.URL.Query().Get("mid"))
	paraments := mux.Vars(r)
	userID, err := strconv.ParseInt(paraments["userId"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	svc := appl.NewUserService()
	followers, err := svc.SearchFollowing(userID)

	result := make([]listSeguindoUserResponse, 0)
	for _, v := range followers {
		result = append(result, listSeguindoUserResponse{
			ID:        v.ID,
			Name:      v.Name,
			Nick:      v.Nick,
			Email:     v.Email,
			Create_at: v.Create_at,
			MID:       mid,
		})
	}

	response.JSON(w, http.StatusAccepted, result)
}

/*
============================
== UPDATEPASSWORD USER =====
============================
*/

type updatepasswordUserRequest struct {
	NewPassword     string `json:"new"`
	CurrentPassword string `json:"current"`
	MID             string `json:"mid"`
}

type updatepasswordUserResponse struct {
	MID string `json:"mid"`
}

// atualiza a senha de usuario
func UpdatePasswordUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var password updatepasswordUserRequest
	err = json.Unmarshal(bodyRequest, &password)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	paraments := mux.Vars(r)
	userID, err := strconv.ParseInt(paraments["userId"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	userIDToken, err := appl.ExtractUsuarioID(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	fmt.Println(userID)
	fmt.Println(userIDToken)

	if userID != userIDToken {
		response.Erro(w, http.StatusForbidden, errors.New("Não é possivel atualizar uma senha o que não é o sua"))
		return
	}

	svc := appl.NewUserService()
	passwordSalveBD, err := svc.SearchPassword(userID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	fmt.Println("Atual: ->", password.CurrentPassword)
	fmt.Println("SenhaSalvaNoBanco ->", passwordSalveBD)

	err = appl.CheckPassword(passwordSalveBD, password.CurrentPassword)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, errors.New("a senha atual não condiz com aque está salva no banco"))
		return
	}

	passwordWithHash, err := appl.Hash(password.NewPassword)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	err = svc.UpdatePassword(userID, string(passwordWithHash))
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	resp := &updatepasswordUserResponse{
		MID: password.MID,
	}

	response.JSON(w, http.StatusOK, resp)

}
