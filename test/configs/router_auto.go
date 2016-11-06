package configs

import (
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/mux"
	"github.com/byorty/hardcore/test/controllers/api"
	"github.com/byorty/hardcore/test/controllers/ws"
)

var (
	Routers = []types.Route{
		mux.Path("/api", 
			mux.Controller("/post", api.NewPost).Batch(
				mux.Get("/", api.PostListAction),
				mux.Get("/{post:([0-9]+)}", api.PostViewAction),
				mux.Post("/{post:([0-9]+)}", api.PostEditAction),
			),
			mux.Controller("/user", api.NewUser).Batch(
				mux.Get("/", api.UserListAction),
				mux.Get("/{user:([0-9]+)}", api.UserViewAction),
			),
			mux.Controller("/test", api.NewTest).Batch(
				mux.Get("/", api.TestViewAction),
			),
		),
		mux.Path("/api", 
			mux.Controller("/ws", ws.NewHandler).Batch(
				mux.Websocket("/", (*ws.Handler).OnMessage),
			
			),
		),
		
	}
)
