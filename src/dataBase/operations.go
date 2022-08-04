package database

import (
	"context"

	"github.com/dgraph-io/dgo/v210/protos/api"
)

func GetAllProgramsDb() (api.Response, error) {
	txnClient := newClient()
	query := `
	{getAllPrograms(func: has(id)){
			id 
			name
			description
		}
	}
	`
	resp, err := txnClient.Query(context.Background(), query)
	return *resp, err
}

func GetProgramDb(id string) (api.Response, error) {
	txnClient := newClient()
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
	resp, err := txnClient.Query(context.Background(), query)
	return *resp, err
}

func SaveProgramDb(newProgram []byte) error {
	txnClient := newClient()

	mutationdb := &api.Mutation{
		CommitNow: true,
		SetJson:   newProgram,
	}
	_, err := txnClient.Mutate(context.Background(), mutationdb)
	return err
}
