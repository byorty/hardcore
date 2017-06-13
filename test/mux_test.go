package test

import (
	"fmt"
	"github.com/byorty/hardcore/mux"
	"github.com/byorty/hardcore/scope"
	"github.com/byorty/hardcore/types"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestRouter(t *testing.T) {
	router := mux.NewHttpRouter()
	server := httptest.NewUnstartedServer(router)
	server.Start()
	serverUrl, _ := url.Parse(server.URL)
	parts := strings.Split(serverUrl.Host, ":")
	host := parts[0]
	port := parts[1]

	router.Add(
		mux.Path("/",
			mux.Get("/hello", func(scope types.RequestScope) {
				scope.GetWriter().Write([]byte("hello world"))
			}),
			mux.Get("/hello/:world?", func(scope types.RequestScope) {
				scope.GetWriter().Write([]byte(fmt.Sprintf("hello %s2", scope.GetPathParams().GetString("world"))))
			}),
			mux.Get("/static/:filename>", func(scope types.RequestScope) {
				scope.GetWriter().Write([]byte(fmt.Sprintf("/static/%s", scope.GetPathParams().GetString("filename"))))
			}),
			//mux.Get("/:id(int)", func(scope types.RequestScope) {
			//	scope.GetWriter().Write([]byte(fmt.Sprintf("id#%v", scope.GetPathParams().GetInt("id"))))
			//}),
			//mux.Get("/:user/:id", func(scope types.RequestScope) {
			//	scope.GetWriter().Write([]byte(scope.GetPathParams().GetString("user") + ":" + scope.GetPathParams().GetString("id")))
			//}),
			mux.Path("/api",
				mux.Get("/:action/extra", func(scope types.RequestScope) {
					scope.GetWriter().Write([]byte(scope.GetPathParams().GetString("action")))
				}),
				mux.Controller("/user", NewTestController).
					Get("/view/:name", (*TestController).View),
			),
			mux.Path("/with/middleware",
				mux.Get("/func", func(scope types.RequestScope) {
					scope.GetWriter().Write([]byte("call func, "))
				}).Before(func(scope types.RequestScope) {
					scope.GetWriter().Write([]byte("before#1, "))
				}).Before(func(scope types.RequestScope) {
					scope.GetWriter().Write([]byte("before#2, "))
				}).After(func(scope types.RequestScope) {
					scope.GetWriter().Write([]byte("after#1"))
				}),
				mux.Controller("/ctrl", NewTestController).
					Before(func(scope types.RequestScope) {
						scope.GetWriter().Write([]byte("before#0, "))
					}).
					Add(
						mux.Get("/view/:name", (*TestController).View).
							Before(func(scope types.RequestScope) {
								scope.GetWriter().Write([]byte("before#1, "))
							}).Before(func(scope types.RequestScope) {
							scope.GetWriter().Write([]byte("before#2, "))
						}).After(func(scope types.RequestScope) {
							scope.GetWriter().Write([]byte(", after#1"))
						}),
					).
					Get("/view2/:name", (*TestController).View).
					After(func(scope types.RequestScope) {
						scope.GetWriter().Write([]byte(", after#0"))
					}),
			),
			//mux.Path("/with/headers",
			//	mux.Get("/func", func(scope types.RequestScope) {
			//		scope.GetWriter().Write([]byte("foo"))
			//	}).
			//		Header("X-SOME-KEY-2", "some#{someId:([0-9]+)}"),
			//	mux.Get("/func", func(scope types.RequestScope) {
			//		someId := scope.GetHeaderParams().GetInt("someId")
			//		otherId := scope.GetHeaderParams().GetInt("otherId")
			//		scope.GetWriter().Write([]byte(fmt.Sprintf("someId#%d, otherId#%d", someId, otherId)))
			//	}).
			//		Header("X-SOME-KEY", "some#{someId:([0-9]+)}").
			//		Header("X-OTHER-KEY", "other#{otherId:([0-9]+)}"),
			//),
			//mux.Path("/with/custom/scope",
			//	mux.Get("/func", func(scope types.RequestScope) {
			//		scope.GetWriter().Write([]byte(scope.(*TestScope).Message))
			//	}).
			//		Scope(NewTestScope),
			//),
		).
			Host(host).
			Port(port),
	)

	sendGet(t, server, "/hello", "hello world")
	sendGet(t, server, "/hello/world", "hello world2")
	sendGet(t, server, "/hello/", "hello 2")
	sendGet(t, server, "/static/some/dir/path.ext", "/static/some/dir/path.ext")

	sendGet(t, server, "/api/call/extra", "call")
	sendGet(t, server, "/api/user/view/trololo", "view user trololo")
	sendGet(t, server, "/with/middleware/func", "before#1, before#2, call func, after#1")
	sendGet(t, server, "/with/middleware/ctrl/view/trololo", "before#0, before#1, before#2, view user trololo, after#0, after#1")
	sendGet(t, server, "/with/middleware/ctrl/view2/trololo12", "before#0, view user trololo12, after#0")

	//client := &http.Client{}
	//req, _ := http.NewRequest("GET", server.URL+"/with/headers/func", nil)
	//req.Header.Add("X-SOME-KEY", "some#123")
	//req.Header.Add("X-OTHER-KEY", "other#456")
	//resp, _ := client.Do(req)
	//checkResponse(t, resp, "someId#123, otherId#456")
	//
	//sendGet(t, server, "/with/custom/scope/func", "hello test scope")
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
	scope.RequestImpl
	Message string
}

func NewTestScope() types.RequestScope {
	return &TestScope{
		Message: "hello test scope",
	}
}

type TestController struct{}

func NewTestController() types.ActionController {
	return new(TestController)
}

func (t *TestController) CallAction(action interface{}, scope types.RequestScope) {
	action.(func(*TestController, types.RequestScope))(t, scope)
}

func (t *TestController) View(scope types.RequestScope) {
	scope.GetWriter().Write([]byte("view user " + scope.GetPathParams().GetString("name")))
}
