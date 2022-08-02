package database

import (
	"context"

	"github.com/dgraph-io/dgo/v210/protos/api"
)

func GetAllProgramsDb() (api.Response, error) {
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

func GetProgramDb(id string) (api.Response, error) {
	dbClient := newClient()
	txn := dbClient.NewTxn()
	query := `
	{getProgram(func: eq(id, ` + id + `)){
			id
			name
			description
			nodes
			drawflow
		}
	}
	`
	resp, err := txn.Query(context.Background(), query)
	return *resp, err
}

func SaveProgramDb(newProgram []byte) error {
	dbClient := newClient()
	txn := dbClient.NewTxn()

	mutationdb := &api.Mutation{
		CommitNow: true,
		SetJson:   newProgram,
	}
	_, err := txn.Mutate(context.Background(), mutationdb)
	return err
}
