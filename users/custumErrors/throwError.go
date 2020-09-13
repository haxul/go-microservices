package custumErrors

import (
	"encoding/json"
	"net/http"
)

func NewError(status int, err error, w *http.ResponseWriter) {
	(*w).WriteHeader(status)
	jsonValue, _ := json.Marshal(AppError{
		Status:  status,
		Message: err.Error(),
	})
	(*w).Write(jsonValue)
}
