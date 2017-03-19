package testserver

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//type Response struct {
//}

type SimpleResource struct {

	// Default implementation of all Resource methods
	ResourceBase
}

/*
func (sr SimpleResource) ConvertData(data interface{}) (content []byte, err error) {

	return []byte(data)
	//return SimpleResource.ResourceBase.ConvertData(data)
}
*/

func (sr SimpleResource) Get(request *http.Request) (int, interface{}) {

	fmt.Println("HTTP METHOD GET")
	return http.StatusOK, request.Body
}

func (sr SimpleResource) Post(request *http.Request) (int, interface{}) {

	buf := make([]byte, 1024)
	fmt.Println("HTTP METHOD POST")
	i, _ := request.Body.Read(buf)

	res := map[string]interface{}{}
	err := json.Unmarshal(buf[:i], &res)
	if err != nil {
		fmt.Println("err")
	}
	//res := buf[:i]
	//fmt.Println(string(res))

	return http.StatusOK, res
}

func (sr SimpleResource) Put(request *http.Request) (int, interface{}) {

	fmt.Println("HTTP METHOD PUT")
	return http.StatusOK, request.Body
}

func (sr SimpleResource) Delete(request *http.Request) (int, interface{}) {

	fmt.Println("HTTP METHOD DELETE")
	return http.StatusOK, request.Body
}
