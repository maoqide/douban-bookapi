package testserver

import (
	"fmt"
	"net/http"
	"net/url"
)

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

func (sr SimpleResource) Get(values url.Values) (int, interface{}) {

	fmt.Println("METHOD GET")
	return http.StatusOK, values
}

func (sr SimpleResource) Post(values url.Values) (int, interface{}) {

	fmt.Println("METHOD POST")
	return http.StatusOK, values
}

func (sr SimpleResource) Put(values url.Values) (int, interface{}) {

	fmt.Println("METHOD PUT")
	return http.StatusOK, values
}

func (sr SimpleResource) Delete(values url.Values) (int, interface{}) {

	fmt.Println("METHOD DELETE")
	return http.StatusOK, values
}
