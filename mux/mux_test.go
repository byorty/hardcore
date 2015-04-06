package mux

import (
	"testing"
	"net/http"
)

type Test struct {
	Request *http.Request
	Route   *Route
}

func TestRouter(t *testing.T) {
//	router := NewRouter(
//		Get("/home", nil),
//		Post("/{foo|bar}", nil),
//		Path("/api",
//			Controller("/maneger", nil).
//				Get("/", nil),
//				Get("/{id:[0-9]+}", nil),
//				Post("/{id:[0-9]+}", nil),
//		),
//	)
//
//	fmt.Println(router)
}

func newRequest(method, url string) *http.Request {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	return req
}


func TestActions(t *testing.T) {
	tests := []Test{
		Test{
			newRequest("GET", "http://google.com/search"),
			Get("/search", nil),
		},
//		Test{
//			Url: url.Parse(""),
//			Get("", nil),
//		},
//		Test{
//			Url: url.Parse(""),
//			Get("", nil),
//		},

	}

	for _, test := range tests {
		test.Route.prepare()
		test.Route.build()
		if !test.Route.Match(test.Request) {
//			t.Fail()
		}
	}
}
