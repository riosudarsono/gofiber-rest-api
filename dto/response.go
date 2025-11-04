package api

import "net/http"

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func CreateResponseError(message string) Response[string] {
	return Response[string]{
		Code:    http.StatusInternalServerError,
		Message: message,
		Data:    "",
	}
}
