package main

import (
	"fmt"
	"net/http"

	authentication "github.com/bhargavasai9999/go-auth/Authentication"
)

func main() {
	authService := authentication.NewAuthService()
	httpHandler := authentication.NewHTTPHandler(authService)

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", httpHandler)
}
