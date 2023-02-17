package controller

import (
	"article_app/entity"
	"article_app/helper"
	"article_app/modules/tag/delivery/http/dto"
	"article_app/modules/tag/usecase"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/mashingan/smapping"
	"log"
	"net/http"
)

type TagController interface {
	GetTags(w http.ResponseWriter, r *http.Request)
	GetTagByID(w http.ResponseWriter, r *http.Request)
	CreateTag(w http.ResponseWriter, r *http.Request)
	UpdateTag(w http.ResponseWriter, r *http.Request)
	DeleteTag(w http.ResponseWriter, r *http.Request)
}

type tagController struct{}

var (
	err error

	validate   *validator.Validate
	tagUsecase usecase.TagUsecase
)

func NewTagController(useCase usecase.TagUsecase) TagController {
	tagUsecase = useCase
	return &tagController{}
}

func (c *tagController) GetTags(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var tags []entity.Tag
	if tags, err = tagUsecase.GetTags(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: err.Error()})
		return
	}

	if len(tags) < 1 {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "Tag Not Found"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Success get all tag",
		"data":    tags,
	})
}

func (c *tagController) GetTagByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var tag *entity.Tag
	params := mux.Vars(r)
	id := params["id"]

	if tag, err = tagUsecase.GetTagByID(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Success get tag",
		"data":    tag,
	})
}

func (c *tagController) CreateTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var tag dto.TagDTO
	if err := json.NewDecoder(r.Body).Decode(&tag); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "Error unmarshalling the request"})
		return
	}
	defer r.Body.Close()

	message := helper.CustomeValidateError(tag)
	if message != "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: message})
		return
	}

	var u = entity.Tag{}
	err := smapping.FillStruct(&u, smapping.MapFields(&tag))
	if err != nil {
		log.Fatalf("failed map: %v", err)
	}

	result, err := tagUsecase.CreateTag(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "Error add tag"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Success add tag",
		"data": map[string]interface{}{
			"id":   result.ID,
			"name": result.Name,
		},
	})
}

func (c *tagController) UpdateTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var tag dto.TagDTO
	params := mux.Vars(r)
	id := params["id"]

	if err := json.NewDecoder(r.Body).Decode(&tag); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "Error unmarshalling the request"})
		return
	}
	defer r.Body.Close()

	message := helper.CustomeValidateError(tag)
	if message != "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: message})
		return
	}

	var u = entity.Tag{}
	err := smapping.FillStruct(&u, smapping.MapFields(&tag))
	if err != nil {
		log.Fatalf("failed map: %v", err)
	}

	result, err := tagUsecase.UpdateTag(id, &u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "Error: " + err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Success update tag",
		"data": map[string]interface{}{
			"id":   result.ID,
			"name": result.Name,
		},
	})
}

func (c *tagController) DeleteTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	if err := tagUsecase.DeleteTag(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "Error delete tag"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Success delete tag",
	})
}
