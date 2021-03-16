package v1

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/orlmonteverde/go-postgres-microblog/internal/middleware"
	"github.com/orlmonteverde/go-postgres-microblog/pkg/response"
	"github.com/orlmonteverde/go-postgres-microblog/pkg/subject"
	"net/http"
	"strconv"
)

// SubjectRouter is the router of the subjects.
type SubjectRouter struct {
	Repository subject.Repository
}

// CreateHandler Create a new subject.
func (sr *SubjectRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var s subject.Subject
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = sr.Repository.Create(ctx, &s)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), s.ID))
	response.JSON(w, r, http.StatusCreated, response.Map{"post": s})
}

// GetAllHandler response all the subjects.
func (sr *SubjectRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	subjects, err := sr.Repository.GetAll(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"posts": subjects})
}

// GetOneHandler response one subject by id.
func (sr *SubjectRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	s, err := sr.Repository.GetOne(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"post": s})
}

// UpdateHandler update a stored subject by id.
func (sr *SubjectRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	var s subject.Subject
	err = json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = sr.Repository.Update(ctx, uint(id), s)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, nil)
}

// DeleteHandler Remove a subject by ID.
func (sr *SubjectRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	err = sr.Repository.Delete(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{})
}

// GetByUserHandler response posts by user id.
func (sr *SubjectRouter) GetByYearHandler(w http.ResponseWriter, r *http.Request) {
	yearStr := chi.URLParam(r, "year")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	subjects, err := sr.Repository.GetByYear(ctx, year)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"posts": subjects})
}

// Routes returns post router with each endpoint.
func (sr *SubjectRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Authorizator)

	r.Get("/subject/{year}", sr.GetByYearHandler)

	r.Get("/", sr.GetAllHandler)

	r.Post("/", sr.CreateHandler)

	r.Get("/{id}", sr.GetOneHandler)

	r.Put("/{id}", sr.UpdateHandler)

	r.Delete("/{id}", sr.DeleteHandler)

	return r
}
