package authentication

import (
	"context"
	"errors"
	"fmt"
)

type AuthService interface {
	Login(ctx context.Context, Email string, password string) (string, error)
	SignUp(ctx context.Context, Email string, Password string, Name string, Mobile string) (string, error)
}

type authService struct {
}

var InvalidLogin = errors.New("Invalid Credentials")
var InSufficientDetails = errors.New("Insufficient Details")

func (a *authService) Login(ctx context.Context, Email string, password string) (string, error) {
	if Email == "bhargava@gmail.com" && password == "123456" {
		fmt.Println("Login Successful")
		return "Login Successful", nil
	}
	fmt.Println("Login Failed")
	return "Login Failed,", InvalidLogin

}
func (a *authService) SignUp(ctx context.Context, Email string, password string, Name string, Mobile string) (string, error) {
	if !(Email != "" && password != "" && Mobile != "" && Name != "") {
		return "Please fill all details", InSufficientDetails
	}
	return "Registration Successful", nil

}
func NewAuthService() AuthService {
	return &authService{}
}
