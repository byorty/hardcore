package test

import (
	"testing"
	"github.com/byorty/hardcore/mux"
	"net/http/httptest"
	"net/url"
	"strings"
	"github.com/byorty/hardcore/types"
	"net/http"
	"io/ioutil"
	"github.com/byorty/hardcore/form"
	"github.com/byorty/hardcore/form/prim"
	"fmt"
)

type FormCtrl struct {}

func NewFormCtrl() types.ActionController {
	return new(FormCtrl)
}

func (f *FormCtrl) CallAction(action interface{}, scope types.RequestScope) {
	if callable, ok := action.(types.CallableAction); ok {
		callable.Call(f, scope)
	} else {
		action.(func(*FormCtrl, types.RequestScope))(f, scope)
	}
}

func (f *FormCtrl) Get(search string, page int) string {
	return fmt.Sprintf("search - %s, page - %d", search, page)
}

type FormCtrlGet func(*FormCtrl, string, int) string

func (f FormCtrlGet) Call(ctrl interface{}, scope types.RequestScope)  {
	var search string
	var page int
	strPrim := prim.String("search")
	strPrim.Required()
	strPrim.Export(&search)
	intPrim := prim.Int("page")
	intPrim.Required()
	intPrim.Export(&page)
	form.New().
		Add(strPrim).
		Add(intPrim).
		Check(scope)

	result := f(ctrl.(*FormCtrl), search, page)
	scope.GetWriter().Write([]byte(result))
}

var (
	formCtrlGet FormCtrlGet = (*FormCtrl).Get
)

func TestForm(t *testing.T) {
	router := mux.NewRouter()
	server := httptest.NewUnstartedServer(router)
	server.Start()
	serverUrl, _ := url.Parse(server.URL)
	parts := strings.Split(serverUrl.Host, ":")
	host := parts[0]
	port := parts[1]

	router.Add(
		mux.Path("/",
			mux.Get("/{name:([a-z]+)}", func(scope types.RequestScope) {
				var name string
				var id int
				strPrim := prim.String("name")
				strPrim.Required()
				strPrim.Export(&name)
				intPrim := prim.Int("id")
				intPrim.Required()
				intPrim.Missing("id should be defined")
				intPrim.Wrong("id is wrong")
				intPrim.Export(&id)

				hasErrors, errors := form.New().
					Add(strPrim).
					Add(intPrim).
					Check(scope)

				if hasErrors {
					for _, error := range errors {
						scope.GetWriter().Write([]byte(error))
					}
				} else {
					scope.GetWriter().Write([]byte(fmt.Sprintf("hello %s#%d", name, id)))
				}

			}),
			mux.Controller("/api/form", NewFormCtrl).
				Get("/", formCtrlGet),
		).
		Host(host).
		Port(port),
	)

	sendGet1(t, server, "/john?id=123", "hello john#123")
	sendGet1(t, server, "/john", "id should be defined")
	sendGet1(t, server, "/john?id=", "id should be defined")
	sendGet1(t, server, "/john?id=qwert", "id is wrong")
	sendGet1(t, server, "/api/form?search=book&page=1", "search - book, page - 1")
}

func sendGet1(t *testing.T, server *httptest.Server, path, needle string) {
	resp, _ := http.Get(server.URL + path)
	checkResponse1(t, resp, needle)
}

func checkResponse1(t *testing.T, resp *http.Response, needle string) {
	body, err := ioutil.ReadAll(resp.Body)
	bodyStr := string(body)
	t.Log("need: " + needle)
	t.Log("receive: " + bodyStr)
	if bodyStr != needle || err != nil {
		t.Fail()
	}
}



