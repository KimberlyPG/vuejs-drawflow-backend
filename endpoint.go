package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dgraph-io/dgo/v210/protos/api"
)

const getUser = `
	{
		node(func: uid(0xfffd8d6ab2dd9f17)) {
			expand (_all_) {
				uid
		}
	  }
	}
`

func getAllPrograms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dgClient := newClient()
	fmt.Print(dgClient)
	txn := dgClient.NewTxn()

	resp, err := txn.Query(context.Background(), getUser)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(resp.Json)
}

var decode struct {
	All []struct {
		id   string
		name string
		age  string
	}
}

// const setNodes Nodes = `
// {
// 	"set":[
// 	  {
// 		"name": "Jojs",
//       	"age": 42,
// 	  }
// 	]
//   }
// `

func setAllPrograms(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	dgClient := newClient()
	txn := dgClient.NewTxn()

	// input := &model{
	// 	r.Body
	// }
	fmt.Print(r.Body)
	inputJson, err := json.Marshal(r.Body)

	mu := &api.Mutation{
		SetJson:   inputJson,
		CommitNow: true,
	}
	log.Println("mutation object:", mu)

	// defer txn.Discard(ctx)
	resp, err := txn.Mutate(context.Background(), mu)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(resp.Json)

	// txn := dgClient.NewTxn()

	// resp, err := txn.Query(context.Background(), setNodes)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// out, err := json.Marshal(decode.All)
	// if err != nil {
	//     log.Fatal(err)
	// }

	// _, err := txn.Mutate(context.Background(), &api.Mutation{SetJson: out})
}
