package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckAuth() Middleware {
	return func(handler http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			var flag bool
			fmt.Println("checking Authentication")

			if !flag {
				handler(writer, request)
			} else {
				return
			}
		}
	}
}

func Login() Middleware {
	return func(handler http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(request.URL.Path, time.Since(start).Milliseconds())
			}()
			handler(writer, request)
		}
	}
}
