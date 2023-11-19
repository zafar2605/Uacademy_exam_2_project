package handler

import (
	"net/http"
	"strings"

	"github.com/spf13/cast"

	"market/model"
)

func (h *Handler) Category(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.CreateCategory(w, r)
	case "GET":
		var (
			values = r.URL.Query()
			method = values.Get("method")
		)
		if method == "GET" {
			h.CategoryGetById(w, r)
		} else if method == "GET_LIST" {
			h.CategoryGetList(w, r)
		} else if method == "SEARCH" {
			h.CategorySearch(w, r)
		}
	case "PUT":
		h.CategoryUpdate(w, r)
	case "DELETE":
		h.CategoryDelete(w, r)
	}
}

func (h *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var userRequest = model.CategoryCreate{
		Title:    r.URL.Query().Get("title"),
		Image:    r.URL.Query().Get("image"),
		ParentID: r.URL.Query().Get("parent_id"),
	}
	resp, err := h.strg.Category().Create(userRequest)
	if err != nil {
		h.handlerResponse(w, http.StatusBadRequest, err.Error(), resp)
		return
	}

	h.handlerResponse(w, http.StatusCreated, "Category Created", resp)
}

func (h *Handler) CategoryGetById(w http.ResponseWriter, r *http.Request) {
	var userId = r.URL.Query().Get("id")

	resp, err := h.strg.Category().GetById(model.CategoryPrimaryKey{Id: userId})
	if err != nil {
		h.handlerResponse(w, 500, "Category does not exist: "+err.Error(), resp)
		return
	}

	h.handlerResponse(w, http.StatusOK, "Category:", resp)
}

func (h *Handler) CategoryGetList(w http.ResponseWriter, r *http.Request) {
	var offsetstr = r.URL.Query().Get("offset")
	var limitstr = r.URL.Query().Get("limit")

	var offset = cast.ToInt(offsetstr)
	var limit = cast.ToInt(limitstr)

	resp, err := h.strg.Category().GetList(model.GetListCategoryRequest{Offset: offset, Limit: limit})
	if err != nil {
		h.handlerResponse(w, 500, "Category does not exist: "+err.Error(), resp)
		return
	}

	h.handlerResponse(w, http.StatusOK, "Category:", resp)
}

func (h *Handler) CategoryUpdate(w http.ResponseWriter, r *http.Request) {

	resp, err := h.strg.Category().Update(model.CategoryUpdate{
		Id:       r.URL.Query().Get("id"),
		Title:    r.URL.Query().Get("title"),
		Image:    r.URL.Query().Get("image"),
		ParentID: r.URL.Query().Get("parent_id"),
	})
	if err != nil {
		h.handlerResponse(w, 500, "Category does not update: "+err.Error(), resp)
		return
	}

	h.handlerResponse(w, http.StatusAccepted, "Updated:", resp)
}

func (h *Handler) CategoryDelete(w http.ResponseWriter, r *http.Request) {
	var userId = model.Category{
		Id: r.URL.Query().Get("id"),
	}

	resp, err := h.strg.Category().Delete(model.CategoryPrimaryKey{Id: userId.Id})
	if err != nil {
		h.handlerResponse(w, 500, "User does not delete: "+err.Error(), resp)
		return
	}

	h.handlerResponse(w, http.StatusAccepted, "Deleted:", resp)
}

func (h *Handler) CategorySearch(w http.ResponseWriter, r *http.Request) {
	var (
		req               = r.URL.Query().Get("title")
		filter_categories = model.GetListCategoryResponse{}
	)

	resp, err := h.strg.Category().GetList(model.GetListCategoryRequest{})
	if err != nil {
		h.handlerResponse(w, 500, "Category does not exists: "+err.Error(), resp)
		return
	}
	for _, v := range resp.Categories {
		if strings.Contains(v.Title, req) {
			filter_categories.Categories = append(filter_categories.Categories, v)
		}
	}

	h.handlerResponse(w, http.StatusAccepted, "Categories:", filter_categories)
}
