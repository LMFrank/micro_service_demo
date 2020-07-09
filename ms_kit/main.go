package main

import (
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	mymux "github.com/gorilla/mux"
	"log"
	"ms_kit/services"
	"ms_kit/util"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	user := services.UserService{}
	endpoint := services.GenUserEndpoint(user)

	serverHandler := httptransport.NewServer(endpoint, services.DecodeUserRequest, services.EncodeUserResponse)

	router := mymux.NewRouter()
	//r.Handle("/user/{uid:\\d+}", serverHandler)
	router.Methods("Get", "DELETE").Path("/user/{uid:\\d+}").Handler(serverHandler)
	router.Methods("GET").Path("/health").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-type", "application/json")
		writer.Write([]byte(`{"status": "ok"}`))
	})

	errChan := make(chan error)

	go func() {
		// 注册服务
		util.RegService()
		err := http.ListenAndServe(":8000", router)
		if err != nil {
			log.Println(err)
			errChan <- err
		}
	}()

	go func() {
		sigChan := make(chan os.Signal)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-sigChan)
	}()

	getErr := <-errChan
	util.UnregService()
	log.Println(getErr)
}
