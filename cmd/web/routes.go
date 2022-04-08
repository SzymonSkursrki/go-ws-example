package main

import (
	"net/http"

	"github.com/SzymonSkursrki/go-ws-example/internal/handlers"

	"github.com/bmizerany/pat"
)

func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))

	return mux
}