package resource

import (
	"API-RS-TOUKIO/internal/appl"
	"API-RS-TOUKIO/internal/domain/publication"
	"API-RS-TOUKIO/internal/infra/data/response"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

/*
============================
== CREATE PUBLICATION ======
============================
*/

type createPublicationRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	MID     string `json:"mid"`
}
type createPublicationResponse struct {
	ID         int64     `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	AuthorID   int64     `json:"authorID"`
	AuthorNick string    `json:"authorNick"`
	Likes      int64     `json:"likes"`
	Create_at  time.Time `json:"create_at"`
	MID        string    `json:"mid"`
}

// cria uma publicação
func CreatePublication(w http.ResponseWriter, r *http.Request) {
	userID, err := appl.ExtractUsuarioID(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publicationRequest createPublicationRequest
	err = json.Unmarshal(bodyRequest, &publicationRequest)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	ent := &publication.Entity{
		Title:    publicationRequest.Title,
		Content:  publicationRequest.Content,
		AuthorID: userID,
	}

	svc := appl.NewPublicationService()
	ent.ID, err = svc.CreatePublication(ent)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	e := &createPublicationResponse{
		ID:         ent.ID,
		Title:      ent.Title,
		Content:    ent.Content,
		AuthorID:   ent.AuthorID,
		AuthorNick: ent.AuthorNick,
		Likes:      ent.Likes,
		Create_at:  ent.Create_at,
		MID:        publicationRequest.MID,
	}

	response.JSON(w, http.StatusCreated, e)

}

/*
============================
== LISTALL PUBLICATION =====
============================
*/

type listALLPublicationRequest struct {
	MID string `json:"-"`
}
type listALLPublicationResponse struct {
	ID         int64     `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	AuthorID   int64     `json:"authorID"`
	AuthorNick string    `json:"authorNick"`
	Likes      int64     `json:"likes"`
	Create_at  time.Time `json:"create_at"`
	MID        string    `json:"mid"`
}

func ListAllPublication(w http.ResponseWriter, r *http.Request) {
	mid := strings.ToLower(r.URL.Query().Get("mid"))
	userID, err := appl.ExtractUsuarioID(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	// bodyRequest, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	response.Erro(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }

	// var publicationRequest listALLPublicationRequest
	// err = json.Unmarshal(bodyRequest, &publicationRequest)
	// if err != nil {
	// 	response.Erro(w, http.StatusInternalServerError, err)
	// 	return
	// }

	svc := appl.NewPublicationService()
	publications, err := svc.ListAllPublication(userID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	result := make([]listALLPublicationResponse, 0)
	for _, v := range publications {
		result = append(result, listALLPublicationResponse{
			ID:         v.ID,
			Title:      v.Title,
			Content:    v.Content,
			AuthorID:   v.AuthorID,
			AuthorNick: v.AuthorNick,
			Likes:      v.Likes,
			Create_at:  v.Create_at,
			MID:        mid,
		})
	}

	response.JSON(w, http.StatusAccepted, result)

}

/*
============================
== FINDBYID PUBLICATION ====
============================
*/

type findByIDPublicationRequest struct {
	MID string `json:"-"`
}
type findByIDPublicationResponse struct {
	ID         int64     `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	AuthorID   int64     `json:"authorID"`
	AuthorNick string    `json:"authorNick"`
	Likes      int64     `json:"likes"`
	Create_at  time.Time `json:"create_at"`
	MID        string    `json:"mid"`
}

func FindByIDPublication(w http.ResponseWriter, r *http.Request) {
	mid := strings.ToLower(r.URL.Query().Get("mid"))
	paraments := mux.Vars(r)
	publicationID, err := strconv.ParseInt(paraments["publicationID"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	svc := appl.NewPublicationService()
	publication, err := svc.FindByIDPublication(publicationID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	resp := &findByIDPublicationResponse{
		ID:         publication.ID,
		Title:      publication.Title,
		Content:    publication.Content,
		AuthorID:   publication.AuthorID,
		AuthorNick: publication.AuthorNick,
		Likes:      publication.Likes,
		Create_at:  publication.Create_at,
		MID:        mid,
	}

	response.JSON(w, http.StatusAccepted, resp)
}

/*
============================
== UPDATE PUBLICATION ======
============================
*/

type updatePublicationRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	MID     string `json:"mid"`
}
type updatePublicationResponse struct {
	MID string `json:"mid"`
}

func UpdatePublication(w http.ResponseWriter, r *http.Request) {
	userID, err := appl.ExtractUsuarioID(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	paraments := mux.Vars(r)
	publicationID, err := strconv.ParseInt(paraments["publicationID"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	svc := appl.NewPublicationService()
	publictionSaveInBD, err := svc.FindByIDPublication(publicationID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if publictionSaveInBD.AuthorID != userID {
		response.Erro(w, http.StatusForbidden, errors.New("Não é possivel atualizar uma publicação que não seja sua"))
		return
	}

	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publicationRequest updatePublicationRequest
	err = json.Unmarshal(bodyRequest, &publicationRequest)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	pub := &publication.Entity{
		Title:   publicationRequest.Title,
		Content: publicationRequest.Content,
	}

	err = svc.UpdatePublication(publicationID, pub)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	pubReponse := &updatePublicationResponse{
		MID: publicationRequest.MID,
	}

	response.JSON(w, http.StatusOK, pubReponse)
}

/*
============================
== DELETE PUBLICATION ======
============================
*/

type deletePublicationRequest struct {
	MID string `json:"-"`
}
type deletePublicationResponse struct {
	MID string `json:"mid"`
}

func DeletePublication(w http.ResponseWriter, r *http.Request) {
	mid := strings.ToLower(r.URL.Query().Get("mid"))
	userID, err := appl.ExtractUsuarioID(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	paraments := mux.Vars(r)
	publicationID, err := strconv.ParseInt(paraments["publicationID"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	svc := appl.NewPublicationService()
	publictionSaveInBD, err := svc.FindByIDPublication(publicationID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if publictionSaveInBD.AuthorID != userID {
		response.Erro(w, http.StatusForbidden, errors.New("Não é possivel Deletar uma publicação que não seja sua"))
		return
	}

	err = svc.DeletePublication(publicationID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	pubReponse := &deletePublicationResponse{
		MID: mid,
	}

	response.JSON(w, http.StatusOK, pubReponse)
}

/*
============================
= LISTBYIDUSER PUBLICATION =
============================
*/

type listByIDUserPublicationRequest struct {
	MID string `json:"-"`
}
type listByIDUserPublicationResponse struct {
	ID         int64     `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	AuthorID   int64     `json:"authorID"`
	AuthorNick string    `json:"authorNick"`
	Likes      int64     `json:"likes"`
	Create_at  time.Time `json:"create_at"`
	MID        string    `json:"mid"`
}

func ListByIDUserPublication(w http.ResponseWriter, r *http.Request) {
	mid := strings.ToLower(r.URL.Query().Get("mid"))
	paraments := mux.Vars(r)
	userID, err := strconv.ParseInt(paraments["userID"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	svc := appl.NewPublicationService()
	publications, err := svc.ListByIDUserPublication(userID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	result := make([]listByIDUserPublicationResponse, 0)
	for _, v := range publications {
		result = append(result, listByIDUserPublicationResponse{
			ID:         v.ID,
			Title:      v.Title,
			Content:    v.Content,
			AuthorID:   v.AuthorID,
			AuthorNick: v.AuthorNick,
			Likes:      v.Likes,
			Create_at:  v.Create_at,
			MID:        mid,
		})
	}

	response.JSON(w, http.StatusOK, result)
}

/*
============================
== LIKE PUBLICATION ========
============================
*/

type likePublicationRequest struct {
	MID string `json:"-"`
}
type likePublicationResponse struct {
	MID string `json:"mid"`
}

func LikePublication(w http.ResponseWriter, r *http.Request) {
	mid := strings.ToLower(r.URL.Query().Get("mid"))
	paraments := mux.Vars(r)
	publicationID, err := strconv.ParseInt(paraments["publicationID"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	svc := appl.NewPublicationService()
	err = svc.LikePublication(publicationID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	e := &likePublicationResponse{
		MID: mid,
	}

	response.JSON(w, http.StatusOK, e)
}

/*
============================
== DESLIKE PUBLICATION =====
============================
*/

type deslikePublicationRequest struct {
	MID string `json:"-"`
}
type deslikePublicationResponse struct {
	MID string `json:"mid"`
}

func DeslikePublication(w http.ResponseWriter, r *http.Request) {
	mid := strings.ToLower(r.URL.Query().Get("mid"))
	paraments := mux.Vars(r)
	publicationID, err := strconv.ParseInt(paraments["publicationID"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	svc := appl.NewPublicationService()
	err = svc.DeslikePublication(publicationID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	e := &deslikePublicationResponse{
		MID: mid,
	}

	response.JSON(w, http.StatusOK, e)
}
