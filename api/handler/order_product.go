package handler

import (
	"net/http"

	"market/model"

	"github.com/spf13/cast"
)

func (h *Handler) OrderProduct(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.AddOrderProducts(w, r)
	case "DELETE":
		h.RemoveOrderProducts(w, r)
	}
}

func (h *Handler) AddOrderProducts(w http.ResponseWriter, r *http.Request) {
	var orderProductReq = model.AddOrderProduct{
		OrderId:      r.URL.Query().Get("order_id"),
		ProductId:    r.URL.Query().Get("product_id"),
		DiscountType: r.URL.Query().Get("discount_type"),
		Quantity:     cast.ToInt(r.URL.Query().Get("quantity")),
	}
	resp, err := h.strg.OrderProduct().Create(orderProductReq)
	if err != nil {
		h.handlerResponse(w, http.StatusBadRequest, err.Error(), resp)
		return
	}

	h.handlerResponse(w, http.StatusCreated, "OrderProducts created", resp)
}

func (h *Handler) RemoveOrderProducts(w http.ResponseWriter, r *http.Request) {
	var userId = model.Order{Id: r.URL.Query().Get("id")}

	resp, err := h.strg.Order().Delete(model.OrderPrimaryKey{Id: userId.Id})
	if err != nil {
		h.handlerResponse(w, 500, "OrderProducts does not delete: "+err.Error(), resp)
		return
	}

	h.handlerResponse(w, http.StatusAccepted, "Deleted:", resp)
}
