package authentication

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

func MakeLoginEndpoint(s AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		response, err := s.Login(ctx, req.Email, req.Password)
		fmt.Println(response, err)
		if err != nil {
			return LoginResponse{Msg: response, Err: err.Error()}, nil
		}
		return LoginResponse{Msg: response, Err: ""}, nil
	}
}
func MakeSignUpEndpoint(s AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SignUpRequest)
		response, err := s.SignUp(ctx, req.Email, req.Password, req.Name, req.Mobile)

		if err != nil {
			return SignUpResponse{Status: response, Err: err}, nil

		}
		return SignUpResponse{Status: response, Err: err}, nil
	}
}
