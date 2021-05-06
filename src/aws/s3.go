package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/giustech/dumper/src/variable"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetLatestSnapshot(sess *session.Session, envi variable.Environments) string {
	result, err:=	listFiles(envi.BucketName, envi.DatabaseDialect, sess)
	if err != nil {
		fmt.Printf("Error to list files on bucket %s", envi.BucketName)
		os.Exit(1)
	}
	if len(result.Contents) == 0 {
		fmt.Printf("There's no snapshots in bucket %s", envi.BucketName)
		os.Exit(1)
	}
	index:=*result.KeyCount-1
	return aws.StringValue(result.Contents[index].Key)
}

func ClearBasedOnSnapshotMax(sess *session.Session, envi variable.Environments){
	result, err:=	listFiles(envi.BucketName, envi.DatabaseDialect, sess)
	if err != nil {
		fmt.Printf("Error to list files on bucket %s", envi.BucketName)
		os.Exit(1)
	}
	var maxExcludeIndex int
	maxExcludeIndex= int(*result.KeyCount) - envi.MaxSnapshots
	if maxExcludeIndex > 0 {
		awsS3:=s3.New(sess)
		for i := 0; i < maxExcludeIndex; i++ {
			req := &s3.DeleteObjectInput{
				Bucket:    aws.String(envi.BucketName),
				Key:       result.Contents[i].Key,
			}
			_, err := awsS3.DeleteObject(req)
			if err != nil {
				fmt.Printf("error: %s\n", err)
			}
		}
	}

}

func Upload(path string, fileName string, sess *session.Session, envi variable.Environments) error {
	parentPath, _ := os.Getwd()
	fullPath := filepath.Join(parentPath, path)
	file, err := os.Open(fullPath)
	if err != nil {
		return err
	}
	uploader := s3manager.NewUploader(sess)
	_, uploadError := uploader.Upload(&s3manager.UploadInput{
		Bucket: &envi.BucketName,
		//ACL:    aws.String("public-read"),
		Key:  &fileName,
		Body: file,
	})
	defer file.Close()
	return uploadError
}

func ListFiles(bucketName string, preffix string) (*s3.ListObjectsV2Output, error) {
	fmt.Printf("Download file = %s in Bucket %s", preffix, bucketName)
	fmt.Println("")
	sess := ConnectAws()
	return listFiles(bucketName, preffix, sess)
}

func listFiles(bucketName string, preffix string, sess *session.Session) (*s3.ListObjectsV2Output, error) {
	awsS3:=s3.New(sess)
	return awsS3.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
		Prefix: aws.String(preffix),
	})
}

func DonwLoad(key string){
	sess := ConnectAws()
	downloader := s3manager.NewDownloader(sess)
	envi:=variable.GetEnvironments()

	buff := &aws.WriteAtBuffer{}
	_, err := downloader.Download(buff, &s3.GetObjectInput{
		Bucket: aws.String(envi.BucketName),
		Key:    aws.String(key),
	})

	if err != nil {
		fmt.Printf("Error to download file %s on bucket %s", key, envi.BucketName)
		os.Exit(3)
	}

	data := buff.Bytes()

	createFileErr := ioutil.WriteFile("restore.dump", data, 0600)
	if createFileErr != nil {
		fmt.Println("Unable to write file")
		fmt.Println(err)
	}

}