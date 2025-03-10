package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/felippemozer/telelaudo-teste-golang-api/internal/domain/user"
	"github.com/felippemozer/telelaudo-teste-golang-api/internal/domain/weather"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error on get env files: ", err)
		return
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/weather", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			city := r.URL.Query().Get("city")
			if city == "" {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Please assign city to query"))
				return
			}

			weather, err := weather.GetBy(city)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			response, err := json.Marshal(weather)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(fmt.Sprint("Error on marshal weather data to bytes", err.Error())))
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(response)
		})
	})

	r.Route("/user", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			id := r.URL.Query().Get("id")
			if id == "" {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("ID not set correctly"))
				return
			}

			idInt, err := strconv.Atoi(id)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("ID is not a number"))
				return
			}

			user, err := user.GetBy(int8(idInt))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}

			response, err := json.Marshal(user)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(fmt.Sprint("Error on marshal user data to bytes", err.Error())))
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(response)
		})
	})

	fmt.Println("Ready to Serve")
	http.ListenAndServe(":4000", r)
}
