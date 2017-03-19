package testserver

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	Get(request *http.Request) (int, interface{})
	Post(request *http.Request) (int, interface{})
	Put(request *http.Request) (int, interface{})
	Delete(request *http.Request) (int, interface{})
	ConvertData(interface{}) ([]byte, error)
}

type ResourceBase struct{}

func (ResourceBase) Get(request *http.Request) (int, interface{}) {
	return http.StatusMethodNotAllowed, ""
}

func (ResourceBase) Post(request *http.Request) (int, interface{}) {
	return http.StatusMethodNotAllowed, ""
}

func (ResourceBase) Put(request *http.Request) (int, interface{}) {
	return http.StatusMethodNotAllowed, ""
}

func (ResourceBase) Delete(request *http.Request) (int, interface{}) {
	return http.StatusMethodNotAllowed, ""
}

func (ResourceBase) ConvertData(data interface{}) (content []byte, err error) {

	//convert data format to response
	//default json
	content, err = json.Marshal(data)
	return
}

type Server struct{}

func (server *Server) Abort(rw http.ResponseWriter, statusCode int) {
	rw.WriteHeader(statusCode)
}

func (server *Server) requestHandler(resource Resource) http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		var data interface{}
		var code int

		request.ParseForm()
		method := request.Method
		//body := request.Body
		switch method {
		case MethodGet:
			code, data = resource.Get(request)
		case MethodPost:
			code, data = resource.Post(request)
		case MethodPut:
			code, data = resource.Put(request)
		case MethodDelete:
			code, data = resource.Delete(request)
		default:
			server.Abort(rw, http.StatusMethodNotAllowed) //405
			return
		}

		content, err := resource.ConvertData(data)
		if err != nil {
			server.Abort(rw, http.StatusInternalServerError) //500
			return
		}
		rw.WriteHeader(code)
		rw.Write(content)
	}
}

func (server *Server) AddResource(resource Resource, path string) {
	http.HandleFunc(path, server.requestHandler(resource))
}

func (server *Server) Start(port int) {
	portString := fmt.Sprintf(":%d", port)
	fmt.Println("Rest Server started on port" + portString)

	http.ListenAndServe(portString, nil)
}

type Test struct {
	// Default implementation of all Resource methods
	ResourceBase
}

// Override the Get method
func (t Test) Get(request *http.Request) (int, interface{}) {

	//result, _ := json.Marshal(values)
	//return http.StatusOK, string(result)
	return http.StatusOK, request.Body
}

func main() {
	var a Test
	server := Server{}
	server.AddResource(a, "/")
	server.Start(4009)
}
