package mux

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"io/ioutil"
	"fmt"
)

func TestRouter(t *testing.T) {
	router := NewRouter()
	server := httptest.NewUnstartedServer(router)
	server.Start()
	serverUrl, _ := url.Parse(server.URL)
	parts := strings.Split(serverUrl.Host, ":")
	host := parts[0]
	port := parts[1]

	router.Add(
		Path("/",
			Get("/", func(scope *RequestScope) {
				scope.Writer.Write([]byte("hello world"))
			}),
			Get("/{id:([0-9]+)}", func(scope *RequestScope) {
				scope.Writer.Write([]byte(fmt.Sprintf("id#%v", scope.PathParams["id"])))
			}),
			Get("/{user:([a-z]+)}/{id:([0-9]+)}", func(scope *RequestScope) {
				scope.Writer.Write([]byte(scope.PathParams["user"] + ":" + scope.PathParams["id"]))
			}),
			Path("/api",
				Get("/{action:([a-z]+)}", func(scope *RequestScope) {
					scope.Writer.Write([]byte(scope.PathParams["action"]))
				}),
				Controller("/user", NewTestController).
					Get("/view/{name:([a-z]+)}", (*TestController).View),
			),
		).
		Host(host).
		Port(port),
	)

	sendGet(t, server, "/", "hello world")
	sendGet(t, server, "/1", "id#1")
	sendGet(t, server, "/qwerty/12", "qwerty:12")
	sendGet(t, server, "/api/call", "call")
	sendGet(t, server, "/api/user/view/trololo", "view user trololo")
}

func sendGet(t *testing.T, server *httptest.Server, path, needle string) {
	resp, _ := http.Get(server.URL + path)
	body, err := ioutil.ReadAll(resp.Body)
	bodyStr := string(body)
	t.Log("need: " + needle)
	t.Log("receive: " + bodyStr)
	if bodyStr != needle || err != nil {
		t.Fail()
	}
}

type TestController struct {}

func NewTestController() ActionController {
	return new(TestController)
}

func (t *TestController) CallAction(action interface{}, scope *RequestScope) {
	action.(func(*TestController, *RequestScope))(t, scope)
}

func (t *TestController) View(scope *RequestScope) {
	scope.Writer.Write([]byte("view user " + scope.PathParams["name"]))
}
