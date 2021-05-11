package v1

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/orlmonteverde/go-postgres-microblog/internal/middleware"
	"github.com/orlmonteverde/go-postgres-microblog/pkg/reply"
	"github.com/orlmonteverde/go-postgres-microblog/pkg/response"
	"net/http"
	"strconv"
)

// ReplyRouter is the router of the replies.
type ReplyRouter struct {
	Repository reply.Repository
}

// CreateHandler Create a new reply.
func (rr *ReplyRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var reply reply.Reply
	err := json.NewDecoder(r.Body).Decode(&reply)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = rr.Repository.Create(ctx, &reply)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), reply.ID))
	response.JSON(w, r, http.StatusCreated, response.Map{"reply": reply})
}

//GetByPostHandler response replies by post id.
func (rr *ReplyRouter) GetByPostHandler(w http.ResponseWriter, r *http.Request) {
	postIDStr := chi.URLParam(r, "postId")

	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	replies, err := rr.Repository.GetByPost(ctx, uint(postID))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"replies": replies})
}

// UpdateHandler update a stored reply by id.
func (rr *ReplyRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	var reply reply.Reply
	err = json.NewDecoder(r.Body).Decode(&reply)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = rr.Repository.Update(ctx, uint(id), reply)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, nil)
}

func (rr *ReplyRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	err = rr.Repository.Delete(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{})
}

// Routes returns reply router with each endpoint.
func (rr *ReplyRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Authorizator)

	r.Get("/post/{postId}", rr.GetByPostHandler)

	r.Post("/", rr.CreateHandler)

	r.Put("/{id}", rr.UpdateHandler)

	r.Delete("/{id}", rr.DeleteHandler)

	return r
}
