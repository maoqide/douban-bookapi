package main

import (
	//"maoqide/bookapi"
	"maoqide/testserver"
)

//func main() {

//	bookapi.Test()
//}

func main() {
	var r testserver.SimpleResource
	server := testserver.Server{}
	server.AddResource(r, "/")
	server.Start(4010)
}
