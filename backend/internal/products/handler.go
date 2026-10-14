package products

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct{
	store *Store
}

func NewHandler (store *Store)*Handler{
	return &Handler{store:store}
}

func (h *Handler)GetAll(w http.ResponseWriter ,r *http.Request){
	products,err:=h.store.GetAll(r.Context())
	if err!=nil{
		http.Error(w,"something went wrong",http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (h *Handler)GetProductVariants(w http.ResponseWriter, r *http.Request){
	id,err:=strconv.Atoi(chi.URLParam(r,"id"))
	if err!=nil{
		http.Error(w,"invalid id",http.StatusBadRequest)
		return
	}
	productVariants,err:=h.store.GetAllProductVariants(r.Context(),id)
	if err!=nil{
		http.Error(w,"something went wrong",http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(productVariants)
}