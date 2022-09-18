package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
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
