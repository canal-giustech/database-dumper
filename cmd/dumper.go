package main

import (
	"fmt"
	"github.com/giustech/dumper/src/aws"
	"github.com/giustech/dumper/src/dump"
	"github.com/giustech/dumper/src/shell"
	"github.com/giustech/dumper/src/variable"
	"time"
)

func main() {
	env:=variable.GetEnvironments()
	database:=dump.GetDataBase(env.DatabaseDialect)
	now:=time.Now();
	if database != nil {
		sess:=aws.ConnectAws()
		fmt.Println(fmt.Sprintf("StartDump"))
		fmt.Println("Dialect Dumper = " + env.DatabaseDialect)
		dumpFileName, err := database.Dump()
		fmt.Println("Generate Md5sum")
		md5sum, errMd5:=shell.Md5SumFile(dumpFileName)
		if errMd5 != nil {
			fmt.Println("Error to Generate Md5Sum")
		}
		fileName:=fmt.Sprintf("%s/%s-%s.dump", env.DatabaseDialect, now.Format("2006_01_02_150405"), md5sum)
		fmt.Printf("Filename: %s", fileName)
		fmt.Println("")
		if err == nil {
			aws.ClearBasedOnSnapshotMax(sess, env)
			fmt.Println("Uploading Dump")
			s3Error:=aws.Upload(dumpFileName, fileName, sess, env)
			fmt.Sprintf("Error = %v", s3Error)
		}
	}
}


