package controllers

import (
	"net/http"
	"umy/templates/pages"
)

func (c *Controller) RenderMediaForm(w http.ResponseWriter, r *http.Request) {
	media := pages.Media()
	err := media.Render(r.Context(), w)
	if err != nil {
		c.HandleServerError(w, err)
	}
}
