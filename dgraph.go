package main

import (
	"log"
	"os"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func newClient() *dgo.Dgraph {
	dotenv := goDotEnvVariable("key")
	// fmt.Printf("godotenv : %s = %s \n", "KEY", dotenv)
	conn, err := dgo.DialCloud("https://nameless-brook-350164.eu-central-1.aws.cloud.dgraph.io/graphql", dotenv)
	if err != nil {
		log.Fatal(err)
	}
	// defer conn.Close()

	dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	return dgraphClient
}
