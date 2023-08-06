package api

import "net/http"

// System related apis

func PlaceholderApi(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
}
