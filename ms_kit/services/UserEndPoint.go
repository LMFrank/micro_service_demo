package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"net/http"
)

type UserRequest struct {
	Uid    int `json:"uid"`
	Method string
}

type UserResponse struct {
	Result string `json:"result"`
}

func GenUserEndpoint(userService IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(UserRequest)
		result := "nothing"
		if r.Method == "GET" {
			result = userService.GetName(r.Uid)
		} else if r.Method == "DELETE" {
			err := userService.DeleteUser(r.Uid)
			if err != nil {
				result = err.Error()
			} else {
				result = fmt.Sprintf("userid为%d的用户删除成功", r.Uid)
			}
		}

		return UserResponse{Result: result}, nil
	}
}

func EncodeUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
