package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/csrf"
	"github.com/gorilla/schema"
	"net/http"
	"os"
	"time"
	"umy/requests"
	"umy/services"
	"umy/templates"
	"umy/templates/pages"
)

var decoder = schema.NewDecoder()
var validate = validator.New(validator.WithRequiredStructEnabled())

func (c *Controller) RenderLogin(w http.ResponseWriter, r *http.Request) {
	login := templates.Login(csrf.Token(r))
	err := login.Render(r.Context(), w)
	if err != nil {
		c.HandleServerError(w, err)
	}
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		c.HandleServerError(w, err)
	}

	r.PostForm.Del("gorilla.csrf.Token")

	var loginRequest requests.Login
	err = decoder.Decode(&loginRequest, r.PostForm)
	if err != nil {
		c.HandleServerError(w, err)
	}

	err = validate.Struct(loginRequest)
	if err != nil {
		c.HandleServerError(w, err)
	}

	adminEmail := os.Getenv("ADMIN_EMAIL")
	adminPass := os.Getenv("ADMIN_PASS")
	if loginRequest.Email == adminEmail && loginRequest.Password == adminPass {
		token, err := services.GenerateJwt()
		if err != nil {
			c.HandleServerError(w, err)
		}

		cookie := &http.Cookie{
			Name:     "accessToken",
			Value:    token,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
		}
		if loginRequest.RememberMe {
			cookie.Expires = time.Now().Add(24 * time.Hour)
		}
		http.SetCookie(w, cookie)

		http.Redirect(w, r, "/admin/dashboard", http.StatusFound)
		return
	}

	http.Error(w, "Invalid email/password", http.StatusUnprocessableEntity)
}

func (c *Controller) RenderDashboard(w http.ResponseWriter, r *http.Request) {
	dash := pages.Dashboard()
	err := dash.Render(r.Context(), w)
	if err != nil {
		c.HandleServerError(w, err)
	}
}
