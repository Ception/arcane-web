package web

import (
	"goarcane"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(store goarcane.AccountStore) *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}

	h.Use(middleware.Logger)
	h.Route("/accounts", func(r chi.Router) {
		r.Get("/", h.AccountsList())
	})

	return h
}

type Handler struct {
	*chi.Mux

	store goarcane.AccountStore
}

const accountsListHTML = `	<h1>Accounts:</h1>
							{{range .Accounts}}
								<dt><strong>{{.Username}}</strong></dt>
								<dd>{{.Email}}</dd>
							{{end}}`

func (h *Handler) AccountsList() http.HandlerFunc {
	type data struct {
		Accounts []goarcane.Account
	}

	tmpl := template.Must(template.New("").Parse(accountsListHTML))
	return func(w http.ResponseWriter, r *http.Request) {
		tt, err := h.store.Accounts()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, data{Accounts: tt})
	}
}
