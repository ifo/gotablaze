package main

import (
	"log"
	"os"

	r "github.com/dancannon/gorethink"
)

// TODO load configuration instead of hard coding constants
const (
	dbAddressEnv = "DBADDRESS"
	dbNameEnv    = "DBNAME"
	gameTableEnv = "GAMETABLE"
)

type DBConfig struct {
	Address   string
	Name      string
	GameTable string
	Indexes   []string
}

type DBSession struct {
	S      *r.Session
	Config DBConfig
}

func defaultDBConfig() DBConfig {
	return DBConfig{
		Address:   "localhost:28015",
		Name:      "test",
		GameTable: "games",
		Indexes:   []string{"timestamp", "match_id"},
	}
}

func (c *DBConfig) Setup() *DBSession {
	if addr := os.Getenv(dbAddressEnv); addr != "" {
		c.Address = addr
	}
	if name := os.Getenv(dbNameEnv); name != "" {
		c.Name = name
	}
	if table := os.Getenv(gameTableEnv); table != "" {
		c.GameTable = table
	}

	sess, err := r.Connect(r.ConnectOpts{
		Address:  c.Address,
		Database: c.Name,
	})
	if err != nil {
		log.Fatalf("RethinkDB Connection error: %s\n", err.Error())
	}

	// set session in DBConfig
	dbs := &DBSession{S: sess, Config: *c}

	// ensure named gameTable exists in database ...
	dbs.CreateTable(c.GameTable)

	// ... with the correct indexes
	dbs.CreateIndexes(c.GameTable, c.Indexes)

	return dbs
}

func (s *DBSession) CreateTable(name string, tableOpts ...r.TableCreateOpts) error {
	opts := r.TableCreateOpts{}
	if len(tableOpts) > 1 {
		log.Fatalln("createTable only takes 0 or 1 arguments")
	} else if len(tableOpts) == 1 {
		opts = tableOpts[0]
	}

	return r.Db(s.Config.Name).TableCreate(name, opts).Exec(s.S)
}

func (s *DBSession) CreateIndexes(name string, indexes []string) []error {
	errs := []error{}
	for _, index := range indexes {
		err := r.Table(s.Config.GameTable).IndexCreate(index, r.IndexCreateOpts{}).Exec(s.S)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func (s *DBSession) SaveGame(game Game) error {
	return r.Table(s.Config.GameTable).Insert(game).Exec(s.S)
}
