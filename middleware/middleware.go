package middleware

import (
	"errors"
	"log"
	"net/http"
	"umy/services"
)

func Admin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cCookie, err := r.Cookie("accessToken")

		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
				return
			}

			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		_, err = services.ValidateJwt(cCookie.Value)
		if err != nil {
			log.Printf("error validating jwt: %v", err)
			http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}
