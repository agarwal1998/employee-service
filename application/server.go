package application

import (
	"net/http"
)

func CreateServer() *http.Server {
	router := NewRouter()
	port := ":1233"
	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
	return server
}
