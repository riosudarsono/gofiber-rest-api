package dto

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

func CreateResponseErrorData(message string, data map[string]string) Response[map[string]string] {
	return Response[map[string]string]{
		Code:    http.StatusInternalServerError,
		Message: message,
		Data:    data,
	}
}

func CreateResponseSuccess[T any](data T) Response[T] {
	return Response[T]{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    data,
	}
}
