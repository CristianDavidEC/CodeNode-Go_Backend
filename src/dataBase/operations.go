package database

import (
	"context"

	"github.com/dgraph-io/dgo/v210/protos/api"
)

func AllProgramsDb() (api.Response, error) {
	dbClient := newClient()
	txn := dbClient.NewTxn()
	query := `
	{getAllPrograms(func: has(id)){
			id 
			name
			description
		}
	}
	`
	resp, err := txn.Query(context.Background(), query)
	return *resp, err
}

func SaveProgramdb(newProgram []byte) error {
	dbClient := newClient()
	txn := dbClient.NewTxn()

	mutationdb := &api.Mutation{
		CommitNow: true,
		SetJson:   newProgram,
	}
	_, err := txn.Mutate(context.Background(), mutationdb)
	return err
}
