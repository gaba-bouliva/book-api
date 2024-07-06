package book

import (
	"book-api/types"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	store types.BookStore
}

func NewBookHandler(store types.BookStore) *Handler{
	return &Handler{
		store,
	}
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {
	router.Get("/", h.handleGetBooks)
	router.Get("/test", h.handleTest)
}

func (h *Handler) handleTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "it works api is live!")
}

func (h *Handler) handleGetBooks(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleGetBook(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleUpdateBook(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleNewBook(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleDeleteBook(w http.ResponseWriter, r *http.Request) {

}