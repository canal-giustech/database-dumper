package dump

import (
	"github.com/giustech/dumper/src/dump/postgres"
)

var (
	Postgres Database = &postgres.PostegresDatabase{}
)

type Database interface {
	Dump() (string, error)
	Dropdatabase()
	Restore(snapshotVersion string) (string, error)
}

func GetDataBase(dialect string) Database {
	switch dialect {
		case "postgres":
			return Postgres
		}
	return nil
}