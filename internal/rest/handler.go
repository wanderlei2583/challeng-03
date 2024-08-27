package rest

import (
	"encoding/json"
	"net/http"

	"orders/internal/orders"
)

type Handler struct {
	repo orders.Repository
}

func (h *Handler) ListOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.repo.ListOrders(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}
