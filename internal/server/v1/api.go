package v1

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/orlmonteverde/go-postgres-microblog/internal/data"
)

// New returns the API V1 Handler with configuration.
func New() http.Handler {
	r := chi.NewRouter()

	ur := &UserRouter{
		Repository: &data.UserRepository{
			Data: data.New(),
		},
	}

	r.Mount("/users", ur.Routes())

	pr := &PostRouter{
		Repository: &data.PostRepository{
			Data: data.New(),
		},
	}

	r.Mount("/posts", pr.Routes())

	sr := &SubjectRouter{
		Repository: &data.SubjectRepository{
			Data: data.New(),
		},
	}

	r.Mount("/subjects", sr.Routes())

	rr := &ReplyRouter{
		Repository: &data.ReplyRepository{
			Data: data.New(),
		},
	}

	r.Mount("/replies", rr.Routes())

	return r
}
