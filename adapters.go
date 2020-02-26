package main

import (
	"context"
	"fmt"
	"net/http"
)

const RequestHeader = "identifier"

// Adapter works as a middleware between request and response
type Adapter func(http.Handler) http.Handler

// Adapt handles all adapters
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

func WithID() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id := r.Header.Get(RequestHeader)
			ctx := context.WithValue(r.Context(), "ID", id)

			fmt.Println("middleware: ", id)

			// pass execution to the original handler
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}