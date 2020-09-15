package model

type GithubErrorResponse struct {
	Message string `json:"message"`
	Errors  []struct {
		Resource string `json:"resource"`
		Code     string `json:"code"`
		Field    string `json:"field"`
		Message  string `json:"message"`
	}
}
