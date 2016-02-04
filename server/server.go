package server

import (
	"github.com/byorty/hardcore/types"
	"net/http"
	"fmt"
)

type Server struct {
	environment types.Environment
}

func New() types.Server {
	return new(Server)
}

func (s *Server) SetEnvironment(environment types.Environment) {
	s.environment = environment
}

func (s Server) Serve() {
	app := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", s.environment.GetHostname(), s.environment.GetPort()),
		Handler:        s.environment.GetRouter(),
		ReadTimeout:    s.environment.GetReadTimeout(),
		WriteTimeout:   s.environment.GetWriteTimeout(),
	}
	fmt.Println(app.ListenAndServe())
}