package utils

import "net/http"

type HandlerFuncWithError func(w http.ResponseWriter, r *http.Request) error

func Handler(next HandlerFuncWithError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := next(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}
