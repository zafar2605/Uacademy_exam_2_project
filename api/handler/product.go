package handler

import (
	"net/http"
	"strings"

	"github.com/spf13/cast"

	"market/model"
)

func (h *Handler) Product(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.CreateProduct(w, r)
	case "GET":
		var (
			values = r.URL.Query()
			method = values.Get("method")
		)
		if method == "GET" {
			h.ProductGetById(w, r)
		} else if method == "GET_LIST" {
			h.ProductGetList(w, r)
		} else if method == "SEARCH" {
			h.ProductSearch(w, r)
		}
	case "PUT":
		h.ProductUpdate(w, r)
	case "DELETE":
		h.ProductDelete(w, r)
	}
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var userRequest = model.ProductCreate{
		Title:       r.URL.Query().Get("title"),
		CategoryId:  r.URL.Query().Get("category_id"),
		Photos:      r.URL.Query().Get("photos"),
		Description: r.URL.Query().Get("description"),
		Price:       cast.ToFloat64(r.URL.Query().Get("price")),
	}
	resp, err := h.strg.Product().Create(userRequest)
	if err != nil {
		h.handlerResponse(w, http.StatusBadRequest, err.Error(), resp)
		return
	}

	h.handlerResponse(w, http.StatusCreated, "Product Created", resp)
}

func (h *Handler) ProductGetById(w http.ResponseWriter, r *http.Request) {
	var userId = r.URL.Query().Get("id")

	resp, err := h.strg.Product().GetById(model.ProductPrimaryKey{Id: userId})
	if err != nil {
		h.handlerResponse(w, 500, "Product does not exist: "+err.Error(), resp)
		return
	}

	h.handlerResponse(w, http.StatusOK, "Product:", resp)
}

func (h *Handler) ProductGetList(w http.ResponseWriter, r *http.Request) {
	var offsetstr = r.URL.Query().Get("offset")
	var limitstr = r.URL.Query().Get("limit")

	var offset = cast.ToInt(offsetstr)
	var limit = cast.ToInt(limitstr)

	resp, err := h.strg.Product().GetList(model.GetListProductRequest{Offset: offset, Limit: limit})
	if err != nil {
		h.handlerResponse(w, 500, "Product does not exist: "+err.Error(), resp)
		return
	}

	h.handlerResponse(w, http.StatusOK, "Product:", resp)
}

func (h *Handler) ProductUpdate(w http.ResponseWriter, r *http.Request) {

	resp, err := h.strg.Product().Update(model.ProductUpdate{
		Id:          r.URL.Query().Get("id"),
		Title:       r.URL.Query().Get("title"),
		CategoryId:  r.URL.Query().Get("category_id"),
		Photos:      r.URL.Query().Get("photos"),
		Description: r.URL.Query().Get("description"),
		Price:       cast.ToFloat64(r.URL.Query().Get("price")),
	})
	if err != nil {
		h.handlerResponse(w, 500, "Product does not update: "+err.Error(), resp)
		return
	}

	h.handlerResponse(w, http.StatusAccepted, "Updated:", resp)
}

func (h *Handler) ProductDelete(w http.ResponseWriter, r *http.Request) {
	var userId = model.Product{
		Id: r.URL.Query().Get("id"),
	}

	resp, err := h.strg.Product().Delete(model.ProductPrimaryKey{Id: userId.Id})
	if err != nil {
		h.handlerResponse(w, 500, "User does not delete: "+err.Error(), resp)
		return
	}

	h.handlerResponse(w, http.StatusAccepted, "Deleted:", resp)
}

func (h *Handler) ProductSearch(w http.ResponseWriter, r *http.Request) {
	var (
		req               = r.URL.Query().Get("title")
		filter_categories = model.GetListProductResponse{}
	)

	resp, err := h.strg.Product().GetList(model.GetListProductRequest{})
	if err != nil {
		h.handlerResponse(w, 500, "Product does not exists: "+err.Error(), resp)
		return
	}
	for _, v := range resp.Products {
		if strings.Contains(v.Title, req) {
			filter_categories.Products = append(filter_categories.Products, v)
		} else if strings.Contains(v.Category.Title, req) {
			filter_categories.Products = append(filter_categories.Products, v)
		}
	}

	h.handlerResponse(w, http.StatusAccepted, "Categories:", filter_categories)
}
