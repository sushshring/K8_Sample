package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sushshring/K8_Sample/common"

	"github.com/sushshring/K8_Sample/db"
	"github.com/sushshring/K8_Sample/model"
)

func main() {
	http.HandleFunc("/", dbConnector(func(w http.ResponseWriter, r *http.Request) {
		request := &model.Request{
			UserAgent: r.UserAgent(),
			Method:    r.Method,
			Path:      fmt.Sprint(r.URL),
			Source:    r.Host,
		}
		if connectedDB := r.Context().Value("db"); connectedDB != nil {
			if err := connectedDB.(db.DB).AddObject(request, "request"); err != nil {
				fmt.Fprintf(w, "Could not save request")
				return
			}
			fmt.Fprintf(w, "Saved request")
			return
		}
		fmt.Fprintf(w, "Could not connect to DB")
		return
	}))

	http.HandleFunc("/requests", dbConnector(func(w http.ResponseWriter, r *http.Request) {
		if connectedDB := r.Context().Value("db"); connectedDB != nil {
			results := connectedDB.(db.DB).GetAllObjects("request")
			request := &model.Request{}
			for results.Next(request) {
				w.Write([]byte(request.String()))
			}
			return
		}
		w.WriteHeader(500)
		w.Write([]byte("Could not connect to DB"))
		return
	}))

	if err := http.ListenAndServe(":"+common.GetEnv("PORT", "8080"), nil); err != nil {
		panic(err)
	}
}

func dbConnector(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		connectedDB := db.DB{}.Connect()
		ctx := context.WithValue(r.Context(), "db", connectedDB)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
