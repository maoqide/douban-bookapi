package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	MethodGet = "GET"
	//MethodHead    = "HEAD"
	MethodPost = "POST"
	MethodPut  = "PUT"
	//MethodPatch   = "PATCH" // RFC 5789
	MethodDelete = "DELETE"
	//MethodConnect = "CONNECT"
	//MethodOptions = "OPTIONS"
	//MethodTrace   = "TRACE"
)

type Resource interface {
	Get(values url.Values) (int, interface{})
	Post(values url.Values) (int, interface{})
	Put(values url.Values) (int, interface{})
	Delete(values url.Values) (int, interface{})
}

type ResourceBase struct{}

func (ResourceBase) Get(values url.Values) (int, interface{}) {
	return http.StatusMethodNotAllowed, ""
}

func (ResourceBase) Post(values url.Values) (int, interface{}) {
	return http.StatusMethodNotAllowed, ""
}

func (ResourceBase) Put(values url.Values) (int, interface{}) {
	return http.StatusMethodNotAllowed, ""
}

func (ResourceBase) Delete(values url.Values) (int, interface{}) {
	return http.StatusMethodNotAllowed, ""
}

type API struct{}

func (api *API) Abort(rw http.ResponseWriter, statusCode int) {
	rw.WriteHeader(statusCode)
}

func (api *API) requestHandler(resource Resource) http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {

		var data interface{}
		var code int

		request.ParseForm()
		method := request.Method
		values := request.Form

		switch method {
		case MethodGet:
			code, data = resource.Get(values)
		case MethodPost:
			code, data = resource.Post(values)
		case MethodPut:
			code, data = resource.Put(values)
		case MethodDelete:
			code, data = resource.Delete(values)
		default:
			api.Abort(rw, http.StatusMethodNotAllowed) //405
			return
		}

		content, err := json.Marshal(data)
		if err != nil {
			api.Abort(rw, http.StatusInternalServerError) //500
			return
		}
		rw.WriteHeader(code)
		rw.Write(content)
	}
}

func (api *API) AddResource(resource Resource, path string) {
	http.HandleFunc(path, api.requestHandler(resource))
}

func (api *API) Start(port int) {
	portString := fmt.Sprintf(":%d", port)
	http.ListenAndServe(portString, nil)
}

type Test struct {
	// Default implementation of all Resource methods
	ResourceBase
}

// Override the Get method
func (t Test) Get(values url.Values) (int, interface{}) {
	return http.StatusOK, "YAY"
}

func main() {
	var a Test
	api := API{}
	api.AddResource(a, "/")
	api.Start(4000)
}
