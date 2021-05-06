package postgres

import (
	"fmt"
	"github.com/giustech/dumper/src/shell"
	"github.com/giustech/dumper/src/variable"
	"io/ioutil"
)

type PostegresDatabase struct {

}

func (database *PostegresDatabase) Dropdatabase() {

}

func (database *PostegresDatabase) generatePgPass(envis variable.Environments) {
	pgpass:=fmt.Sprintf("%s:%s:%s:%s:%s", envis.Hostname, envis.Port, envis.DatabaseName, envis.Username, envis.Password)
	fmt.Println("Creating .Pgpass")
	err := ioutil.WriteFile("/.pgpass", []byte(pgpass), 0600)
	if err != nil {
		fmt.Println("Unable to write file")
		fmt.Println(err)
	}
}

func (database *PostegresDatabase) Restore(snapshotVersion string) (string, error) {
	envis:=variable.GetEnvironments()
	database.generatePgPass(envis)
	return shell.Execute("pg_restore", "-h", envis.Hostname, "-U", envis.Username, "-d", envis.DatabaseName, "restore.dump")
}

func (database *PostegresDatabase) Dump() (string, error) {
	envis:=variable.GetEnvironments()
	database.generatePgPass(envis)
	_, errExec :=shell.Execute("pg_dump", "-f", "dump_file", "-h", envis.Hostname, "-U", envis.Username, "-d", envis.DatabaseName, "-Fc")
	return "dump_file", errExec
}