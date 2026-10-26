package replenishment

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	store *Store
}

func NewHandler(store *Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	replenishmentRequests, err := h.store.GetAll(r.Context())
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(replenishmentRequests)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Items []ReplenishmentItem `json:"items"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	req, err := h.store.Create(r.Context(), body.Items)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(req)
}
func (h *Handler) AddItems(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Items []ReplenishmentItem `json:"items"`
	}
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	err = h.store.AddItems(r.Context(), id, body.Items)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return

	}
	w.WriteHeader(http.StatusOK)

}
func (h *Handler) NextLocation(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid request id", http.StatusBadRequest)
		return
	}
	warehouse_id, err := strconv.Atoi(r.URL.Query().Get("warehouse_id"))
	if err != nil {
		http.Error(w, "invalid warehouse id ", http.StatusBadRequest)
		return
	}
	location, err := h.store.GetNextLocation(r.Context(), warehouse_id, id)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(location)

}
