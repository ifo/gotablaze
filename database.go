package main

import (
	"log"
	"os"

	r "github.com/dancannon/gorethink"
)

// TODO pass these as context instead of using a global
var (
	session   *r.Session = nil
	dbAddress            = "localhost:28015"
	dbName               = "test"
	gameTable            = "games"
	indexes              = []string{"timestamp", "match_id"}
)

// TODO load configuration instead of hard coding constants
const (
	dbAddressEnv = "DBADDRESS"
	dbNameEnv    = "DBNAME"
	gameTableEnv = "GAMETABLE"
)

func setupRethinkDB() *r.Session {
	if os.Getenv(dbAddressEnv) != "" {
		dbAddress = os.Getenv(dbAddressEnv)
	}
	if os.Getenv(dbNameEnv) != "" {
		dbName = os.Getenv(dbNameEnv)
	}
	if os.Getenv(gameTableEnv) != "" {
		gameTable = os.Getenv(gameTableEnv)
	}

	sess, err := r.Connect(r.ConnectOpts{
		Address:  dbAddress,
		Database: dbName,
	})
	if err != nil {
		log.Fatalf("RethinkDB Connection error: %s\n", err.Error())
	}

	// ensure named gameTable exists in database
	createTable(gameTable, sess)
	// with the correct indexes
	createIndexes(gameTable, indexes, sess)

	// TODO not use a global
	session = sess
	return session
}

func createTable(name string, s *r.Session, tableOpts ...r.TableCreateOpts) error {
	opts := r.TableCreateOpts{}
	if len(tableOpts) > 1 {
		log.Fatalln("createTable only takes 0 or 1 arguments")
	} else if len(tableOpts) == 1 {
		opts = tableOpts[0]
	}

	// TODO return the object in a useable format instead of just the error
	err := r.Db(dbName).TableCreate(name, opts).Exec(s)
	return err
}

func createIndexes(name string, indexes []string, s *r.Session) []error {
	errs := []error{}
	for _, index := range indexes {
		err := r.Db(dbName).IndexCreate(index, r.IndexCreateOpts{}).Exec(s)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

// TODO rewrite this to not use a global
func gameTableQuery() r.Term {
	return r.Table(gameTable)
}

func saveGame(game Game, s *r.Session) error {
	err := gameTableQuery().Insert(game).Exec(s)
	return err
}
