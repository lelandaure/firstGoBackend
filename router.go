package main

import (
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, exist := r.rules[path]
	handler, methodExist := r.rules[path][method]
	return handler, methodExist, exist
}

func (r *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	handler, methodExist, exist := r.FindHandler(request.URL.Path, request.Method)

	if !exist {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	if !methodExist {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(writer, request)
}

func NewRouter() *Router {
	return &Router{rules: make(map[string]map[string]http.HandlerFunc)}
}
