package main

import (
	"fmt"
	"github.com/giustech/dumper/src/aws"
	"github.com/giustech/dumper/src/dump"
	"github.com/giustech/dumper/src/variable"
	"os"
	"strings"
)

func main() {
	env:=variable.GetEnvironments()
	database:=dump.GetDataBase(env.DatabaseDialect)
	sess:=aws.ConnectAws()
	var version string
	if env.SnapshotVersion != "" {
		version = env.SnapshotVersion
	} else {
		version=strings.Replace(aws.GetLatestSnapshot(sess,env), fmt.Sprintf("%s/", env.DatabaseDialect), "", 1)
		version=strings.Split(version, "-")[0]
	}

	if version == "" {
		fmt.Println("SnapshotVersion is null")
		os.Exit(1)
	}

	list, err:=aws.ListFiles(env.BucketName, fmt.Sprintf("%s/%s", env.DatabaseDialect, version))
	if err != nil || len(list.Contents) == 0 {
		fmt.Printf("Error to get dump file %s/%s", env.DatabaseDialect, version)
		fmt.Println("")
		os.Exit(2)
	}

	if len(list.Contents) > 1 {
		fmt.Printf("There's 2 dump files with this configuration %s/%s", env.DatabaseDialect, version)
	}

	aws.DonwLoad(*list.Contents[0].Key)
	_, errorRestore:=database.Restore(version)
	if errorRestore != nil {
		fmt.Printf("Error %v\n", errorRestore)
		fmt.Printf("Error to Restore database %s/%s\n", env.DatabaseDialect, version)
	} else {
		fmt.Println("Dump Completed")
	}

}
