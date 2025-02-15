package apinoexternallib

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("Request %s processed in %s\n", r.URL.Path, time.Since(begin))
	})
}

func Serve() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Println(id)
		fmt.Fprintln(w, "Heatlhcheck")
	})

	srv := &http.Server{
		Addr:                         "localhost:8080",
		Handler:                      Log(mux),
		DisableGeneralOptionsHandler: false,
		ReadTimeout:                  10 * time.Second,
		WriteTimeout:                 10 * time.Second,
		IdleTimeout:                  1 * time.Minute,
		MaxHeaderBytes:               0,
	}

	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}
}
