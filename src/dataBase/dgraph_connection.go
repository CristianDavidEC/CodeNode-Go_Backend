package database

import (
	"fmt"
	"log"
	"os"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
)

//var ClienteDb *dgo.Dgraph

/* Create the connection with the BD */
func NewClient() {
	DATABASE_CONNECTION := os.Getenv("DATABASE_CONNECTION")
	API_KEY := os.Getenv("API_KEY")
	fmt.Println(DATABASE_CONNECTION)
	fmt.Println(API_KEY)
	conn, err := dgo.DialCloud(DATABASE_CONNECTION, API_KEY)
	if err != nil {
		log.Fatal(err)
	}
	dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))
	setTxnClient(dgraphClient)
}

var txnClient *dgo.Txn

func setTxnClient(clienteDb *dgo.Dgraph) {
	txnClient = clienteDb.NewTxn()
}
