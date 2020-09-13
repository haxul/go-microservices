package custumErrors

type AppError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
