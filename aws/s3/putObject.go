// +build example

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Lists all objects in a bucket using pagination
//
// Usage:
// listObjects <bucket>
func main() {

	if len(os.Args) < 7 {
		fmt.Println("you must specify 1region, 2bucket, 3name, 4extension, 5data and 6total")
		return
	}

	sess, errSess := session.NewSession()
	if errSess != nil {
		fmt.Println("failed to create session", errSess)
		return
	}

	svc := s3.New(sess, aws.NewConfig().WithRegion(os.Args[1]))

	fileBytes, errFile := ioutil.ReadFile(os.Args[5])
	if errFile != nil {
		fmt.Println("failed to read file", errFile)
		return
	}

	intTotal, errintTotal := strconv.Atoi(os.Args[6])

	if errintTotal != nil {
		fmt.Println("failed to convert to integer", errintTotal)
		return
	}

	for i := 0; i < intTotal; i++ {

		nameExt := os.Args[3] + strconv.Itoa(i) + "." + os.Args[4]

		params := &s3.PutObjectInput{
			Bucket: &os.Args[2],
			Key:    &nameExt,
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
}
