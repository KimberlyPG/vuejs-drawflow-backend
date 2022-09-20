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

		resp, err := txn.Query(context.Background(), getNodeData)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(resp.Json)
	})

	type Data struct {
		Result int `json:"result,omitempty"`
	}
	type Connections1 struct {
		Node  string `json:"node"`
		Input string `json:"input"`
	}
	type Input1 struct {
		Connections []Connections1 `json:"connections"`
	}
	type Input2 struct {
		Connections []Connections1 `json:"connections"`
	}
	type Inputs struct {
		Input1 Input1 `json:"input_1"`
		Input2 Input2 `json:"input_2"`
	}
	type Connections2 struct {
		Node   string `json:"node"`
		Output string `json:"output"`
	}
	type Output1 struct {
		Connections []Connections2 `json:"connections"`
	}
	type Outputs struct {
		Output1 Output1 `json:"output_1"`
	}
	type DataNodes struct {
		Nid      int     `json:"id"`
		Name     string  `json:"name"`
		Data     Data    `json:"data"`
		Class    string  `json:"class"`
		HTML     string  `json:"html"`
		Typenode string  `json:"typenode"`
		Inputs   Inputs  `json:"inputs"`
		Outputs  Outputs `json:"outputs"`
		PosX     float64 `json:"pos_x"`
		PosY     float64 `json:"pos_y"`
	}

	type AutoGenerated struct {
		Data []DataNodes `json:"drawflow,omitempty"`
	}

	router.Post("/setAllPrograms", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var p AutoGenerated

		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "Decode error! please check your JSON formating.")
		}

		fmt.Printf("Inputed name: %s", p)
		dgClient := newClient()
		txn := dgClient.NewTxn()

		ctx := context.Background()

		mu := &api.Mutation{
			CommitNow: true,
		}

		pb, err := json.Marshal(p)
		if err != nil {
			log.Fatal(err)
		}

		mu.SetJson = pb

		log.Println("mutation object:", mu)

		resp, err := txn.Mutate(ctx, mu)

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
