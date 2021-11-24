package api

import (
	"net/http"

	"github.com/mark-api/route"
)

var (
	Router = route.New()
)

func Get(pattern string, handler http.HandlerFunc) {
	Router.Get(pattern, handler)
}

func Post(pattern string, handler http.HandlerFunc) {
	Router.Post(pattern, handler)
}

func Put(pattern string, handler http.HandlerFunc) {
	Router.Del(pattern, handler)
}

func Del(pattern string, handler http.HandlerFunc) {
	Router.Del(pattern, handler)
}
