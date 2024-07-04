package dto

import "net/http"

type CreateCustomerDto struct {
	ParentID string `json:"parent_id"`
}

func (mr *CreateCustomerDto) Bind(r *http.Request) error {
	return nil
}
