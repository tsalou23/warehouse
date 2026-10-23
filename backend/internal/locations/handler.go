package locations

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

func (h *Handler) GetLocationsByWarehouseId(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid warehouse id", http.StatusBadRequest)
		return
	}
	locations, err := h.store.GetLocationsByWarehouseId(r.Context(), id)
	if err != nil {
		http.Error(w, "something went wrong could not get the locations", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(locations)

}
