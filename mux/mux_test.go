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
			Get("/", func(scope RequestScope) {
				scope.GetWriter().Write([]byte("hello world"))
			}),
			Get("/{id:([0-9]+)}", func(scope RequestScope) {
				scope.GetWriter().Write([]byte(fmt.Sprintf("id#%v", scope.GetPathParams()["id"])))
			}),
			Get("/{user:([a-z]+)}/{id:([0-9]+)}", func(scope RequestScope) {
				scope.GetWriter().Write([]byte(scope.GetPathParams()["user"] + ":" + scope.GetPathParams()["id"]))
			}),
			Path("/api",
				Get("/{action:([a-z]+)}", func(scope RequestScope) {
					scope.GetWriter().Write([]byte(scope.GetPathParams()["action"]))
				}),
				Controller("/user", NewTestController).
					Get("/view/{name:([a-z]+)}", (*TestController).View),
			),
			Path("/with/middleware",
				Get("/func", func(scope RequestScope) {
					scope.GetWriter().Write([]byte("call func, "))
				}).Before(func (scope RequestScope) {
					scope.GetWriter().Write([]byte("before#1, "))
				}).Before(func (scope RequestScope) {
					scope.GetWriter().Write([]byte("before#2, "))
				}).After(func (scope RequestScope) {
					scope.GetWriter().Write([]byte("after#1"))
				}),
				Controller("/ctrl", NewTestController).
					Before(func (scope RequestScope) {
						scope.GetWriter().Write([]byte("before#0, "))
					}).
					Add(
						Get("/view/{name:([a-z]+)}", (*TestController).View).
							Before(func (scope RequestScope) {
								scope.GetWriter().Write([]byte("before#1, "))
							}).Before(func (scope RequestScope) {
								scope.GetWriter().Write([]byte("before#2, "))
							}).After(func (scope RequestScope) {
								scope.GetWriter().Write([]byte(", after#1"))
							}),
					).
					Get("/view2/{name:([a-z0-9]+)}", testControllerActionView).
					After(func (scope RequestScope) {
						scope.GetWriter().Write([]byte(", after#0"))
					}),
			),
			Path("/with/headers",
				Get("/func", func (scope RequestScope) {
					scope.GetWriter().Write([]byte("foo"))
				}).
				Header("X-SOME-KEY-2", "some#{someId:([0-9]+)}"),
				Get("/func", func (scope RequestScope) {
					someId := scope.GetHeaderParams().GetInt("someId")
					otherId := scope.GetHeaderParams().GetInt("otherId")
					scope.GetWriter().Write([]byte(fmt.Sprintf("someId#%d, otherId#%d", someId, otherId)))
				}).
				Header("X-SOME-KEY", "some#{someId:([0-9]+)}").
				Header("X-OTHER-KEY", "other#{otherId:([0-9]+)}"),
			),
			Path("/with/custom/scope",
				Get("/func", func(scope RequestScope) {
					scope.GetWriter().Write([]byte(scope.(*TestScope).Message))
				}).
				Scope(NewTestScope),
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
	sendGet(t, server, "/with/middleware/func", "before#1, before#2, call func, after#1")
	sendGet(t, server, "/with/middleware/ctrl/view/trololo", "before#0, before#1, before#2, view user trololo, after#0, after#1")
	sendGet(t, server, "/with/middleware/ctrl/view2/trololo12", "before#0, view user trololo12, after#0")

	client := &http.Client{}
	req, _ := http.NewRequest("GET", server.URL + "/with/headers/func", nil)
	req.Header.Add("X-SOME-KEY", "some#123")
	req.Header.Add("X-OTHER-KEY", "other#456")
	resp, _ := client.Do(req)
	checkResponse(t, resp, "someId#123, otherId#456")

	sendGet(t, server, "/with/custom/scope/func", "hello test scope")
}

func sendGet(t *testing.T, server *httptest.Server, path, needle string) {
	resp, _ := http.Get(server.URL + path)
	checkResponse(t, resp, needle)
}

func checkResponse(t *testing.T, resp *http.Response, needle string) {
	body, err := ioutil.ReadAll(resp.Body)
	bodyStr := string(body)
	t.Log("need: " + needle)
	t.Log("receive: " + bodyStr)
	if bodyStr != needle || err != nil {
		t.Fail()
	}
}

type TestScope struct {
	BaseRequestScope
	Message string
}

func NewTestScope() RequestScope {
	return &TestScope{
		Message: "hello test scope",
	}
}

type TestController struct {}

func NewTestController() ActionController {
	return new(TestController)
}

func (t *TestController) CallAction(action interface{}, scope RequestScope) {
	typeAction, ok := action.(TestControllerActionView)
	if ok {
		typeAction(t, scope)
	} else {
		action.(func(*TestController, RequestScope))(t, scope)
	}
}

func (t *TestController) View(scope RequestScope) {
	scope.GetWriter().Write([]byte("view user " + scope.GetPathParams()["name"]))
}

var (
	testControllerActionView TestControllerActionView = (*TestController).View
)

type TestControllerActionView func(*TestController, RequestScope)

func (a *TestControllerActionView) Form() {

}
