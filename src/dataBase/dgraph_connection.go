package database

import (
	"log"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
)

func newClient() *dgo.Dgraph {
	// This example uses dgo
	conn, err := dgo.DialCloud("https://blue-surf-591055.us-east-1.aws.cloud.dgraph.io/graphql", "YTZlYmRiZmIxZDg4MTM1MWFkMTQzNDZiYzcxNjE4MWM=")
	if err != nil {
		log.Fatal(err)
	}
	//defer conn.Close()
	dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	return dgraphClient
}
