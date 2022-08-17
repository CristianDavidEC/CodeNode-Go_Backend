package database

import (
	"log"

	"context"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
)

type ClientDgraph struct {
	*dgo.Dgraph
}

/*Create a Dgraph client*/
func NewDgraphClient(conectionDb string, apiKey string) (*ClientDgraph, error) {
	conn, err := dgo.DialCloud(conectionDb, apiKey)
	if err != nil {
		log.Fatal(err)
	}
	dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))
	defer conn.Close()
	return &ClientDgraph{dgraphClient}, nil
}

/*Queries all programs stored in  the database*/
func (client *ClientDgraph) GetAllProgramsDb() ([]byte, error) {
	query := `
	{getAllPrograms(func: has(id)){
			id 
			name
			description
		}
	}
	`
	resp, err := client.NewTxn().Query(context.Background(), query)
	return resp.GetJson(), err
}

/*Query the program by id*/
func (client *ClientDgraph) GetProgramDb(id string) ([]byte, error) {
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
	resp, err := client.NewTxn().Query(context.Background(), query)
	return resp.GetJson(), err
}

/*Save the program in the db*/
func (client *ClientDgraph) SaveProgramDb(newProgram []byte) error {
	mutationdb := &api.Mutation{
		CommitNow: true,
		SetJson:   newProgram,
	}
	_, err := client.NewTxn().Mutate(context.Background(), mutationdb)
	return err
}
