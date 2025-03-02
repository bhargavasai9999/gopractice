package authentication

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Msg string `json:"message"`
	Err string `json:"Err"`
}

type SignUpRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
}
type SignUpResponse struct {
	Status string
	Err    error
}
