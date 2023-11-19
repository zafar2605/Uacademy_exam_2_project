package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"market/config"
	"market/storage"
)

type Handler struct {
	cfg  *config.Config
	strg storage.StorageI
}

func NewHanler(cfg *config.Config, strg storage.StorageI) Handler {
	return Handler{
		cfg:  cfg,
		strg: strg,
	}
}

type response struct {
	Status      int         `json:"status"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

func (h *Handler) handlerResponse(w http.ResponseWriter, code int, message string, data interface{}) {
	resp := response{
		Status:      code,
		Description: message,
		Data:        data,
	}

	log.Printf("%+v\n", resp)

	body, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(code)
	w.Write([]byte(body))
}
