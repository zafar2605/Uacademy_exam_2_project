package api

import (
	"net/http"

	"market/api/handler"
	"market/config"
	"market/storage"
)

func NewApi(cfg *config.Config, strg storage.StorageI) {
	handler := handler.NewHanler(cfg, strg)

	http.HandleFunc("/category", handler.Category)
	http.HandleFunc("/product", handler.Product)
	http.HandleFunc("/branch", handler.Branch)
	http.HandleFunc("/client", handler.Client)
	http.HandleFunc("/orders", handler.Order)
	http.HandleFunc("/order_product", handler.OrderProduct)
}
