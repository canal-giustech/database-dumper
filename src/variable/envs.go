package variable

import (
	"os"
	"strconv"
)

type Environments struct {
	AccessKeyId string
	SecretAccessKey string
	Region          string
	BucketName      string
	Hostname        string
	Port            string
	Username        string
	Password        string
	DatabaseName    string
	DatabaseDialect string
	MaxSnapshots    int
	SnapshotVersion string
}

func GetEnvironments() Environments {

	var maxSnapshots int
	maxSnapshots, err:=strconv.Atoi(os.Getenv("MAX_SNAPSHOTS"))

	if err != nil {
		maxSnapshots = 10
	}

	envi:=Environments{
		AccessKeyId:     os.Getenv("AWS_ACCESS_KEY_ID"),
		SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		Region:          os.Getenv("AWS_DEFAULT_REGION"),
		BucketName:      os.Getenv("AWS_BUCKET"),
		Hostname:        os.Getenv("DB_HOST"),
		Port:            os.Getenv("DB_PORT"),
		Username:        os.Getenv("DB_USER"),
		Password:        os.Getenv("DB_PASSWORD"),
		DatabaseName:    os.Getenv("DB_NAME"),
		DatabaseDialect: os.Getenv("DUMPER_DIALECT"),
		MaxSnapshots:    maxSnapshots - 1,
		SnapshotVersion: os.Getenv("SNAPSHOT_VERSION"),
	}
	return envi

}

