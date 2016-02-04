package configs

import (
	"github.com/byorty/hardcore/mux"
	"github.com/byorty/hardcore/test/controllers/api"
)

func Router() *mux.Router {
	return router
}

var (
	router = mux.NewRouter().Add(
		mux.Path("/api", 
			mux.Controller("/post", api.NewPost).Batch(
				mux.Get("/", api.PostListAction),
				mux.Get("/{post:([0-9]+)}", api.PostViewAction),
				mux.Post("/{post:([0-9]+)}", api.PostEditAction),
			),
			mux.Controller("/user", api.NewUser).Batch(
				mux.Get("/", api.UserListAction),
			),
		),
	).Add(makeRouter())
)
