package handler

import (
	dto "customer/internal/dto/customer"
	"customer/internal/response"
	"customer/internal/service"
	"net/http"

	"github.com/go-chi/render"
)

type CustomerHandler struct {
	customerService service.CustomerService
}

func NewCustomerHandler(customerService service.CustomerService) CustomerHandler {
	return CustomerHandler{
		customerService: customerService,
	}
}

func (h CustomerHandler) Create(w http.ResponseWriter, r *http.Request) {
	data := &dto.CreateCustomerDto{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, response.Response[any]{
			Status:     "failed",
			StatusCode: 400,
			Message:    "Bad request",
		})
		return
	}
	customer, err := h.customerService.Store(data.ParentID)
	if err != nil {
		render.Render(w, r, response.Response[any]{
			Status:     "failed",
			StatusCode: 500,
			Message:    err.Error(),
		})
		return
	}
	res := response.Response[any]{
		Status:     "success",
		StatusCode: 200,
		Message:    "Berhasil",
		Data:       customer,
	}
	render.Render(w, r, res)
}
