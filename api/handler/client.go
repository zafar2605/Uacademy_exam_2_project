package handler

import (
	"net/http"
	"strings"

	"github.com/spf13/cast"

	"market/model"
)

func (h *Handler) Client(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.CreateClient(w, r)
	case "GET":
		var (
			values = r.URL.Query()
			method = values.Get("method")
		)
		if method == "GET" {
			h.ClientGetById(w, r)
		} else if method == "GET_LIST" {
			h.ClientGetList(w, r)
		} else if method == "SEARCH" {
			h.ClientSearch(w, r)
		}
	case "PUT":
		h.ClientUpdate(w, r)
	case "DELETE":
		h.ClientDelete(w, r)
	}
}

func (h *Handler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var userRequest = model.ClientCreate{
		FirstName:   r.URL.Query().Get("firstname"),
		LastName:    r.URL.Query().Get("lastname"),
		Phone:       r.URL.Query().Get("phone"),
		Photo:       r.URL.Query().Get("photp"),
		DateOfBirth: r.URL.Query().Get("date_of_birth"),
	}
	resp, err := h.strg.Client().Create(userRequest)
	if err != nil {
		h.handlerResponse(w, http.StatusBadRequest, err.Error(), resp)
		return
	}

	h.handlerResponse(w, http.StatusCreated, "Client Created", resp)
}

func (h *Handler) ClientGetById(w http.ResponseWriter, r *http.Request) {
	var userId = r.URL.Query().Get("id")

	resp, err := h.strg.Client().GetById(model.ClientPrimaryKey{Id: userId})
	if err != nil {
		h.handlerResponse(w, 500, "Client does not exist: "+err.Error(), resp)
		return
	}

	h.handlerResponse(w, http.StatusOK, "Client:", resp)
}

func (h *Handler) ClientGetList(w http.ResponseWriter, r *http.Request) {
	var offsetstr = r.URL.Query().Get("offset")
	var limitstr = r.URL.Query().Get("limit")

	var offset = cast.ToInt(offsetstr)
	var limit = cast.ToInt(limitstr)

	resp, err := h.strg.Client().GetList(model.GetListClientRequest{Offset: offset, Limit: limit})
	if err != nil {
		h.handlerResponse(w, 500, "Client does not exist: "+err.Error(), resp)
		return
	}

	h.handlerResponse(w, http.StatusOK, "Client:", resp)
}

func (h *Handler) ClientUpdate(w http.ResponseWriter, r *http.Request) {

	resp, err := h.strg.Client().Update(model.ClientUpdate{
		Id:          r.URL.Query().Get("id"),
		FirstName:   r.URL.Query().Get("firstname"),
		LastName:    r.URL.Query().Get("lastname"),
		Phone:       r.URL.Query().Get("phone"),
		Photo:       r.URL.Query().Get("photp"),
		DateOfBirth: r.URL.Query().Get("date_of_birth"),
	})
	if err != nil {
		h.handlerResponse(w, 500, "Client does not update: "+err.Error(), resp)
		return
	}

	h.handlerResponse(w, http.StatusAccepted, "Updated:", resp)
}

func (h *Handler) ClientDelete(w http.ResponseWriter, r *http.Request) {
	var userId = model.Client{
		Id: r.URL.Query().Get("id"),
	}

	resp, err := h.strg.Client().Delete(model.ClientPrimaryKey{Id: userId.Id})
	if err != nil {
		h.handlerResponse(w, 500, "User does not delete: "+err.Error(), resp)
		return
	}

	h.handlerResponse(w, http.StatusAccepted, "Deleted:", resp)
}

func (h *Handler) ClientSearch(w http.ResponseWriter, r *http.Request) {
	var (
		req               = r.URL.Query().Get("key")
		filter_clients = model.GetListClientResponse{}
	)

	resp, err := h.strg.Client().GetList(model.GetListClientRequest{})
	if err != nil {
		h.handlerResponse(w, 500, "Client does not exists: "+err.Error(), resp)
		return
	}
	for _, v := range resp.Clients {
		if strings.Contains(v.FirstName, req) || strings.Contains(v.LastName, req) || strings.Contains(v.Phone, req) {
			filter_clients.Clients = append(filter_clients.Clients, v)
		}
	}

	h.handlerResponse(w, http.StatusAccepted, "Categories:", filter_clients)
}
