# go-rr
Go RR AWS - S3 Utility

## Example

listObjects is an example using the AWS SDK for Go to list objects' key in a S3 bucket.


## Usage

The example uses the the bucket name provided, and lists all object keys in a bucket.

```sh
$ go run listObjects.go <region> <bucket>
$ go run putObject.go <region> <bucket> <filename> <extension> <file / template> <total>
$ go run putObject.go ap-southeast-2 rr-go jac txt ../../templates/jac.json 2
```

Output:
```
Page, 0
Object: test.txt
```
