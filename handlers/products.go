package handlers

import (
	"net/http"

	"github.com/brunompx/angula/model"
	"github.com/brunompx/angula/views"
)

func (h *Handler) HandleListProducts(w http.ResponseWriter, r *http.Request) {
	isAddingProduct := r.URL.Query().Get("isAddingProduct") == "true"

	products, err := h.store.GetProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var categories []model.Category
	if isAddingProduct {
		categories, err = h.store.GetCategories()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	views.OrderEdit(products, categories, isAddingProduct).Render(r.Context(), w)
}

/*
func (h *Handler) HandleAddProduct(w http.ResponseWriter, r *http.Request) {
	product := &model.Product{
		User:        "bruno",
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Price:       r.FormValue("price"),
		Category:    r.FormValue("category"),
	}
	if r.FormValue("active") == "on" {
		product.Active = true
	} else {
		product.Active = false
	}
	product.Stock, _ = strconv.Atoi(r.FormValue("stock"))
	product.CategoryId, _ = strconv.Atoi(r.FormValue("category"))

	newProduct, err := h.store.CreateProduct(product)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	components.ProductTile(newProduct).Render(r.Context(), w)
}

func (h *Handler) HandleSearchProduct(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("search")

	product, err := h.store.FindProduct(text)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	components.ProductsList(product).Render(r.Context(), w)
}
*/
