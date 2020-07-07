package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"ms_kit/services"
	"net/http"
)

func main() {
	user := services.UserService{}
	endpoint := services.GenUserEndpoint(user)

	serverHandler := httptransport.NewServer(endpoint, services.DecodeUserRequest, services.EncodeUserResponse)

	http.ListenAndServe(":8000", serverHandler)
}
