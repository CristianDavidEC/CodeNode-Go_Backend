package database

import (
	"context"

	"github.com/dgraph-io/dgo/v210/protos/api"
)

func AllProgramsDb() (api.Response, error) {
	dbClient := newClient()
	txn := dbClient.NewTxn()
	query := `
	{getAllPrograms(func: has(Program.id)){
			id: Program.id 
			name: Program.name
			description: Program.description
		}
	  }
	`
	resp, err := txn.Query(context.Background(), query)
	return *resp, err
}
