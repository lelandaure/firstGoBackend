package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "HelloWORLD!")
}

func HandleHome(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Api...")
}
