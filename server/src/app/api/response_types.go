package api

type ErrorResponse struct {
	Error string `json:"error"`
}

type ValidateTokenResponse struct {
	Valid bool `json:"valid"`
}
