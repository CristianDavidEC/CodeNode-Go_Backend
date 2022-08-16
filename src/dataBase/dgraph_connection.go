package database

import (
	"fmt"
	"log"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
)

//var txnClient *dgo.Txn

/* Create the connection with the BD */
func NewClient(connectionDb string, apiKey string) *dgo.Dgraph {
	fmt.Println("Connecting to the database...")
	conn, err := dgo.DialCloud(connectionDb, apiKey)
	if err != nil {
		log.Fatal(err)
	}
	dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))
	return dgraphClient
}
