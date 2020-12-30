package main

import (
	"net/http"
	"time"

	"github.com/yakushou730/blog-service/internal/routers"
)

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:              ":8080",
		Handler:           router,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	s.ListenAndServe()
}
