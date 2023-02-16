package controller

import (
	"article_app/entity"
	"article_app/helper"
	"article_app/modules/post/delivery/http/dto"
	"article_app/modules/post/usecase"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type PostController interface {
	GetPosts(w http.ResponseWriter, r *http.Request)
	GetPostByID(w http.ResponseWriter, r *http.Request)
	GetPostByUser(w http.ResponseWriter, r *http.Request)
	CreatePost(w http.ResponseWriter, r *http.Request)
	UpdatePost(w http.ResponseWriter, r *http.Request)
	DeletePost(w http.ResponseWriter, r *http.Request)
}

type postController struct{}

var (
	err error

	validate    *validator.Validate
	postUsecase usecase.PostUsecase
)

func NewPostController(useCase usecase.PostUsecase) PostController {
	postUsecase = useCase
	return &postController{}
}

func (c *postController) GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var posts []entity.Post
	if posts, err = postUsecase.GetPosts(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: err.Error()})
		return
	}

	if len(posts) < 1 {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "Post Not Found"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Success get all post",
		"data":    posts,
	})
}

func (c *postController) GetPostByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post *entity.Post
	params := mux.Vars(r)
	id := params["id"]
	userId, _ := strconv.Atoi(r.Header.Get("userId"))

	if post, err = postUsecase.GetPostByID(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: err.Error()})
		return
	}

	if post.UserID != userId {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "Unauthorized, Not your posts!"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Success get post",
		"data":    post,
	})
}

func (c *postController) GetPostByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post []entity.Post
	params := mux.Vars(r)
	id := params["id"]
	userId := r.Header.Get("userId")

	if userId != id {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "Unauthorized, Not your posts!"})
		return
	}

	if post, err = postUsecase.GetPostByUser(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Success get post",
		"data":    post,
	})
}

func (c *postController) CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post dto.PostDTO
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "Error unmarshalling the request"})
		return
	}
	defer r.Body.Close()

	message := helper.CustomeValidateError(post)
	if message != "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: message})
		return
	}

	result, err := postUsecase.CreatePost(&post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "Error add post"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Success add post",
		"data": map[string]interface{}{
			"id":    result.ID,
			"body":  result.Body,
			"title": result.Title,
		},
	})
}

func (c *postController) UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post dto.PostDTO
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "Error unmarshalling the request"})
		return
	}
	defer r.Body.Close()

	message := helper.CustomeValidateError(post)
	if message != "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: message})
		return
	}

	userId, _ := strconv.Atoi(r.Header.Get("userId"))
	checkPost, err := postUsecase.GetPostByID(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: err.Error()})
		return
	}

	if checkPost.UserID != userId {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "Unauthorized, Not your posts!"})
		return
	}

	result, err := postUsecase.UpdatePost(id, &post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "Error: " + err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Success update post",
		"data": map[string]interface{}{
			"id":    result.ID,
			"body":  result.Body,
			"title": result.Title,
		},
	})
}

func (c *postController) DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	userId, _ := strconv.Atoi(r.Header.Get("userId"))

	post, err := postUsecase.GetPostByID(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: err.Error()})
		return
	}

	if post.UserID != userId {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "Unauthorized, Not your posts!"})
		return
	}

	if err := postUsecase.DeletePost(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ServiceError{Message: "Error delete post"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Success delete post",
	})
}
