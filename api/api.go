package api

import (
	v1 "app/api/handlers/v1"
	"fmt"
	"net/http"
)

func NewServer(adr string , api *http.ServeMux) http.Server{

	return http.Server{
		Addr: adr,
		Handler: api,
	}
}

func Api(){
	
	h:=v1.NewHandler()

	api:=http.NewServeMux()

	api.HandleFunc("GET /ping",h.Pong )
	api.HandleFunc("GET /getusers",h.GetUsers)

	server:=NewServer(":8080",api)

	fmt.Println("http server is running on port: 8080")

	server.ListenAndServe()
}