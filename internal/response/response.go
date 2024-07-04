package response

import "net/http"

type Response[T any] struct {
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       T      `json:"data"`
}

func (Response[T]) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
