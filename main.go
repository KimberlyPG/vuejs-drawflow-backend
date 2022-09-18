package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgraph-io/dgo/v210/protos/api"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*", "*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)

	router.Get("/getAllPrograms", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		dgClient := newClient()
		fmt.Print(dgClient)
		txn := dgClient.NewTxn()

		resp, err := txn.Query(context.Background(), getUser)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(resp.Json)
	})

	type Data struct {
		number string
		result int
		total  int
		num1   string
		num2   string
		option string
		assign string
	}

	type Program struct {
		class    string
		html     string
		data     Data
		id       int
		pos_x    float64
		pos_y    float64
		typenode string
	}

	router.Post("/setAllPrograms", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var p Program
		newjson := json.NewDecoder(r.Body).Decode(&p)
		dgClient := newClient()
		txn := dgClient.NewTxn()

		fmt.Printf("body", r.Body)
		inputJson, err := json.Marshal(newjson)

		mu := &api.Mutation{
			SetJson:   inputJson,
			CommitNow: true,
		}
		log.Println("mutation object:", mu)

		resp, err := txn.Mutate(context.Background(), mu)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(resp.Json)
	})

	ServerGo := &http.Server{
		Addr:           ":5000",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := ServerGo.ListenAndServe()
	if err != nil {
		fmt.Println("Server error", err.Error())
	} else {
		fmt.Println("Server running in :5000")
	}
}
