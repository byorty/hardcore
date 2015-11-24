package hardcore

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
		).
		Host(host).
		Port(port),
	)

	sendGet(t, server, "/john?id=123", "hello john#123")
	sendGet(t, server, "/john", "id should be defined")
	sendGet(t, server, "/john?id=", "id should be defined")
	sendGet(t, server, "/john?id=qwert", "id is wrong")
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



