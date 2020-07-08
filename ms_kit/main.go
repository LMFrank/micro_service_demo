package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	mymux "github.com/gorilla/mux"
	"ms_kit/services"
	"net/http"
)

func main() {
	user := services.UserService{}
	endpoint := services.GenUserEndpoint(user)

	serverHandler := httptransport.NewServer(endpoint, services.DecodeUserRequest, services.EncodeUserResponse)

	r := mymux.NewRouter()
	//r.Handle("/user/{uid:\\d+}", serverHandler)
	r.Methods("Get", "DELETE").Path("/user/{uid:\\d+}").Handler(serverHandler)

	http.ListenAndServe(":8000", r)
}
