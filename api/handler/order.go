package handler

import (
	"net/http"
	"strings"

	"market/model"
)

func (h *Handler) Order(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.CreateOrder(w, r)
	case "GET":
		var (
			values = r.URL.Query()
			method = values.Get("method")
		)
		if method == "GET" {
			h.OrderGetById(w, r)
		} else if method == "GET_LIST" {
			h.OrderGetList(w, r)
		} else if method == "SEARCH" {
			h.OrderSearch(w, r)
		}
	case "PUT":
		var (
			values = r.URL.Query()
			method = values.Get("method")
		)
		if method == "STATUS" {
			h.OrderStatusUpdate(w, r)
		} else if method == "" {
			h.OrderUpdate(w, r)
		}

	case "DELETE":
		h.OrderDelete(w, r)
	}
}

func (h *Handler) OrderStatusUpdate(w http.ResponseWriter, r *http.Request) {

	resp, err := h.strg.Order().StatusUpdate(model.OrderStatusUpdate{
		Id:     r.URL.Query().Get("id"),
		Status: r.URL.Query().Get("status"),
	})

	if err != nil {
		h.handlerResponse(w, http.StatusBadRequest, "Order status do not update: "+err.Error(), resp)
		return
	}

	h.handlerResponse(w, http.StatusAccepted, "Status updated:", resp)
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var orderRequest = model.OrderCreate{
		Client:  r.URL.Query().Get("client_id"),
		Branch:  r.URL.Query().Get("branch_id"),
		Address: r.URL.Query().Get("address"),
	}
	resp, err := h.strg.Order().Create(orderRequest)
	if err != nil {
		h.handlerResponse(w, http.StatusBadRequest, err.Error(), resp)
		return
	}

	h.handlerResponse(w, http.StatusCreated, "Order Created", resp)
}

func (h *Handler) OrderGetById(w http.ResponseWriter, r *http.Request) {
	var userId = r.URL.Query().Get("id")

	resp, err := h.strg.Order().GetById(model.OrderPrimaryKey{Id: userId})
	if err != nil {
		h.handlerResponse(w, http.StatusBadRequest, "Order does not exist: "+err.Error(), resp)
		return
	}

	h.handlerResponse(w, http.StatusOK, "Order:", resp)
}

func (h *Handler) OrderGetList(w http.ResponseWriter, r *http.Request) {

	resp, err := h.strg.Order().GetList(model.GetListOrderRequest{})
	if err != nil {
		h.handlerResponse(w, 500, "Order does not exist: "+err.Error(), resp)
		return
	}
	for i, v := range resp.Orders {
		res, err := h.strg.OrderProduct().GetByOrderId(model.GetOrderProductByOrderId{Id: v.Id})
		if err != nil {
			h.handlerResponse(w, 500, "OrderProduct does not exist: "+err.Error(), resp)
		}
		resp.Orders[i].OrderProducts = *res
	}

	h.handlerResponse(w, http.StatusOK, "Order:", resp)
}

func (h *Handler) OrderUpdate(w http.ResponseWriter, r *http.Request) {

	resp, err := h.strg.Order().Update(model.OrderUpdate{})
	if err != nil {
		h.handlerResponse(w, http.StatusBadRequest, "Order does not update: "+err.Error(), resp)
		return
	}

	h.handlerResponse(w, http.StatusAccepted, "Updated:", resp)
}

func (h *Handler) OrderDelete(w http.ResponseWriter, r *http.Request) {
	var userId = model.Order{
		Id: r.URL.Query().Get("id"),
	}

	resp, err := h.strg.Order().Delete(model.OrderPrimaryKey{Id: userId.Id})
	if err != nil {
		h.handlerResponse(w, 500, "Order does not delete: "+err.Error(), resp)
		return
	}

	h.handlerResponse(w, http.StatusAccepted, "Deleted:", resp)
}

func (h *Handler) OrderSearch(w http.ResponseWriter, r *http.Request) {
	var (
		req              = r.URL.Query().Get("req")
		filtered_order_p = model.GetListOrderResponse{}
	)
	resp, err := h.strg.Order().GetList(model.GetListOrderRequest{})
	if err != nil {
		h.handlerResponse(w, 500, "Order do not find: "+err.Error(), resp)
		return
	}
	for _, v := range resp.Orders {

		if strings.Contains(v.Id, req) {
			filtered_order_p.Orders = append(filtered_order_p.Orders, v)
		} else if strings.Contains(v.ClientId, req) {
			filtered_order_p.Orders = append(filtered_order_p.Orders, v)
		} else if strings.Contains(v.BranchId, req) {
			filtered_order_p.Orders = append(filtered_order_p.Orders, v)
		}
	}

	h.handlerResponse(w, http.StatusAccepted, "Orders:", filtered_order_p)
}
