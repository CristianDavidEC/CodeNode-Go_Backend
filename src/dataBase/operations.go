package database

import (
	"context"

	"github.com/dgraph-io/dgo/v210/protos/api"
)

/*Queries all programs stored in  the database*/
func GetAllProgramsDb() (api.Response, error) {
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

/*Query the program by id*/
func GetProgramDb(id string) (api.Response, error) {
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

/*Save the program in the db*/
func SaveProgramDb(newProgram []byte) error {
	mutationdb := &api.Mutation{
		CommitNow: true,
		SetJson:   newProgram,
	}
	_, err := txnClient.Mutate(context.Background(), mutationdb)
	return err
}
