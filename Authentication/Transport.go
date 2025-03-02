package authentication

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func NewHTTPHandler(s AuthService) http.Handler {
	mux := http.NewServeMux()

	loginHandler := httptransport.NewServer(
		MakeLoginEndpoint(s),
		decodeLoginRequest,
		encodeResponse,
	)

	mux.Handle("/login", loginHandler)
	return mux
}
