package test

import (
	"fmt"
	"github.com/byorty/hardcore/form"
	"github.com/byorty/hardcore/form/prim"
	"github.com/byorty/hardcore/mux"
	"github.com/byorty/hardcore/test/models"
	"github.com/byorty/hardcore/types"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

type FormCtrl struct{}

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

func (f FormCtrlGet) Call(ctrl interface{}, scope types.RequestScope) {
	var search string
	var page int
	strPrim := prim.String("search")
	strPrim.Required()
	strPrim.SetSource(types.GetPrimitiveSource)
	strPrim.Export(&search)
	intPrim := prim.Int("page")
	intPrim.Required()
	intPrim.SetSource(types.GetPrimitiveSource)
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
				strPrim.SetSource(types.PathPrimitiveSource)
				strPrim.Required()
				strPrim.Export(&name)
				intPrim := prim.Int("id")
				intPrim.Required()
				intPrim.SetSource(types.GetPrimitiveSource)
				intPrim.Missing("id should be defined")
				intPrim.Wrong("id is wrong")
				intPrim.Export(&id)

				testForm := form.New().
					Add(strPrim).
					Add(intPrim)

				if testForm.Check(scope) {
					scope.GetWriter().Write([]byte(fmt.Sprintf("hello %s#%d", name, id)))
				} else {
					for i := 0; i < testForm.GetErrors().Len(); i++ {
						scope.GetWriter().Write([]byte(testForm.GetErrors().Get(i).GetMessage()))
					}
				}

			}),
			mux.Get("/user/{id:([0-9]+)}", func(scope types.RequestScope) {
				var user models.User
				userPrim := prim.Int64Model("id")
				userPrim.Required()
				userPrim.SetSource(types.PathPrimitiveSource)
				userPrim.Missing("user should be defined")
				userPrim.Wrong("user is wrong")
				userPrim.Export(&user)

				var role models.UserRole
				rolePrim := prim.IntEnum("role")
				rolePrim.Required()
				rolePrim.SetSource(types.GetPrimitiveSource)
				rolePrim.Missing("role should be defined")
				rolePrim.Wrong("role is wrong")
				rolePrim.Export(&role)

				testForm := form.New().
					Add(userPrim).
					Add(rolePrim)

				if testForm.Check(scope) {
					scope.GetWriter().Write([]byte(fmt.Sprintf("user#%d - email: %s, role: %s", user.GetId(), user.GetEmail(), role.GetName())))
				} else {
					for i := 0; i < testForm.GetErrors().Len(); i++ {
						scope.GetWriter().Write([]byte(testForm.GetErrors().Get(i).GetMessage()))
					}
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
	sendGet1(t, server, "/user/1?role=1", "user#1 - email: 1451507434558119572@qwerty.com, role: Admin")
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
