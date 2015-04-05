package mux

import (
	"testing"
	"fmt"
)

func TestRouter(t *testing.T) {
	router := Router(
		Get("/home", nil),
		Post("/signin", nil),
		Path("/api",
			Controller("/maneger", nil).
				Get("/", nil),
				Get("/{id:[0-9]+}", nil),
				Post("/{id:[0-9]+}", nil),
		),
	)

	fmt.Println(router)
}
