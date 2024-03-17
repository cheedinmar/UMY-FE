package main

import (
	"fmt"
	"github.com/gorilla/csrf"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
	"umy/controllers"
	"umy/middleware"
)

const (
	ConnHost = "localhost"
	ConnPort = "8080"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	controller := controllers.NewController()

	router := mux.NewRouter()
	router.HandleFunc("/auth/login", controller.RenderLogin).Methods("GET")
	router.HandleFunc("/auth/login", controller.Login).Methods("POST")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./resources/public"))))

	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.Use(middleware.Admin)
	adminRouter.HandleFunc("/dashboard", controller.RenderDashboard).Methods("GET")

	compressed := handlers.CompressHandler(router)
	loggedRouter := handlers.LoggingHandler(os.Stdout, compressed)
	csrfMiddleware := csrf.Protect([]byte(os.Getenv("APP_KEY")))

	log.Printf("Server starting on port %v\n", ConnPort)
	srv := &http.Server{
		Handler:      csrfMiddleware(loggedRouter),
		Addr:         fmt.Sprintf("%s:%s", ConnHost, ConnPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
