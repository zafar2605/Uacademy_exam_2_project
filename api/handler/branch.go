package handler

import (
	"net/http"
	"strings"

	"market/model"
)

func (h *Handler) Branch(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.CreateBranch(w, r)
	case "GET":
		var (
			values = r.URL.Query()
			method = values.Get("method")
		)
		if method == "GET" {
			h.BranchGetById(w, r)
		} else if method == "GET_LIST" {
			h.BranchGetList(w, r)
		} else if method == "SEARCH" {
			h.BranchSearch(w, r)
		}
	case "PUT":
		h.BranchUpdate(w, r)
	case "DELETE":
		h.BranchDelete(w, r)
	}
}

func (h *Handler) CreateBranch(w http.ResponseWriter, r *http.Request) {
	var userRequest = model.BranchCreate{
		Name:          r.URL.Query().Get("name"),
		Phone:         r.URL.Query().Get("phone"),
		Photo:         r.URL.Query().Get("photo"),
		Adress:        r.URL.Query().Get("address"),
		Active:        r.URL.Query().Get("active"),
		WorkStartHour: r.URL.Query().Get("work_start_hour"),
		WorkEndHour:   r.URL.Query().Get("work_end_hour"),
	}
	resp, err := h.strg.Branch().Create(userRequest)
	if err != nil {
		h.handlerResponse(w, http.StatusBadRequest, err.Error(), resp)
		return
	}

	h.handlerResponse(w, http.StatusCreated, "Branch Created", resp)
}

func (h *Handler) BranchGetById(w http.ResponseWriter, r *http.Request) {
	var userId = r.URL.Query().Get("id")

	resp, err := h.strg.Branch().GetById(model.BranchPrimaryKey{Id: userId})
	if err != nil {
		h.handlerResponse(w, 500, "Branch does not exist: "+err.Error(), resp)
		return
	}

	h.handlerResponse(w, http.StatusOK, "Branch:", resp)
}

func (h *Handler) BranchGetList(w http.ResponseWriter, r *http.Request) {

	resp, err := h.strg.Branch().GetList(model.GetListBranchRequest{})

	if err != nil {
		h.handlerResponse(w, 500, "Branch does not exist: "+err.Error(), resp)
		return
	}
	h.handlerResponse(w, http.StatusOK, "Branch:", resp)
}

func (h *Handler) BranchUpdate(w http.ResponseWriter, r *http.Request) {

	resp, err := h.strg.Branch().Update(model.BranchUpdate{
		Id:            r.URL.Query().Get("id"),
		Name:          r.URL.Query().Get("name"),
		Phone:         r.URL.Query().Get("phone"),
		Photo:         r.URL.Query().Get("photo"),
		Adress:        r.URL.Query().Get("address"),
		WorkStartHour: r.URL.Query().Get("work_start_hour"),
		WorkEndHour:   r.URL.Query().Get("work_end_hour"),
	})
	if err != nil {
		h.handlerResponse(w, 500, "Branch does not update: "+err.Error(), resp)
		return
	}

	h.handlerResponse(w, http.StatusAccepted, "Updated:", resp)
}

func (h *Handler) BranchDelete(w http.ResponseWriter, r *http.Request) {
	var userId = model.Branch{
		Id: r.URL.Query().Get("id"),
	}

	resp, err := h.strg.Branch().Delete(model.BranchPrimaryKey{Id: userId.Id})
	if err != nil {
		h.handlerResponse(w, 500, "Branch does not delete: "+err.Error(), resp)
		return
	}

	h.handlerResponse(w, http.StatusAccepted, "Deleted:", resp)
}

func (h *Handler) BranchSearch(w http.ResponseWriter, r *http.Request) {
	var (
		req               = r.URL.Query().Get("name")
		filter_categories = model.GetListBranchResponse{}
	)

	resp, err := h.strg.Branch().GetList(model.GetListBranchRequest{})
	if err != nil {
		h.handlerResponse(w, 500, "Branch does not exists: "+err.Error(), resp)
		return
	}
	for _, v := range resp.Branches {
		if strings.Contains(v.Name, req) {
			filter_categories.Branches = append(filter_categories.Branches, v)
		} else if strings.Contains(v.CreatedAt, req) {
			filter_categories.Branches = append(filter_categories.Branches, v)
		}
	}

	h.handlerResponse(w, http.StatusAccepted, "Categories:", filter_categories)
}
