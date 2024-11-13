package db

import (
	"context"
	"log"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
)

const (
	dgraphAddress = "127.0.0.1:9180"
)

var database *dgo.Dgraph
var cancel CancelFunc
var lock = sync.Mutex{}

type CancelFunc func()

type Definition struct {
	id          int
	name        string
	description string
	phonetic    string
	refs        []int
	aliases     []int
}

func NewClient() (*dgo.Dgraph, CancelFunc) {
	if database == nil {
		db, _cancel := newClient()
		database = db
		cancel = _cancel
	}
	return database, cancel
}

func newClient() (*dgo.Dgraph, CancelFunc) {
	conn, err := grpc.NewClient(dgraphAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)
	ctx := context.Background()

	// Perform login call. If the Dgraph cluster does not have ACL and
	// enterprise features enabled, this call should be skipped.
	for {
		// Keep retrying until we succeed or receive a non-retriable error.
		err = dg.Login(ctx, "groot", "password")
		if err == nil || !strings.Contains(err.Error(), "Please retry") {
			break
		}
		time.Sleep(time.Second)
	}
	if err != nil {
		log.Fatalf("While trying to login %v", err.Error())
	}

	return dg, func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error while closing connection:%v", err)
		}
	}
}
