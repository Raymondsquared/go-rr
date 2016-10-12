// +build example

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Lists all objects in a bucket using pagination
//
// Usage:
// listObjects <bucket>
func main() {

	if len(os.Args) < 5 {
		fmt.Println("you must specify region, bucket, name and data")
		return
	}

	sess, errSess := session.NewSession()
	if errSess != nil {
		fmt.Println("failed to create session", errSess)
		return
	}

	svc := s3.New(sess, aws.NewConfig().WithRegion(os.Args[1]))

	fileBytes, errFile := ioutil.ReadFile(os.Args[4])
	if errFile != nil {
		fmt.Println("failed to read file", errFile)
		return
	}

	params := &s3.PutObjectInput{
		Bucket: &os.Args[2],
		Key:    &os.Args[3],
		Body:   bytes.NewReader(fileBytes),
	}

	resp, errPutObj := svc.PutObject(params)
	if errPutObj != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(errPutObj.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
