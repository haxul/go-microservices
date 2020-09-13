package errors

import "net/http"

func NewError(status int, message error, w *http.ResponseWriter) {
	(*w).WriteHeader(status)
	(*w).Write([]byte(message.Error()))
}
