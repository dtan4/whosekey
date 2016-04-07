package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

const Usage = `
Usage:
  $ whosekey AWS_ACCESS_KEY_ID

Options:
  -h, --help Show this usage
`

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, Usage)
		os.Exit(1)
	}

	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println(Usage)
		return
	}

	accessKeyId := os.Args[1]
	svc := iam.New(session.New(), &aws.Config{})

	resp, err := svc.ListAccessKeys(nil)
	if err != nil {
		panic(err)
	}

	for _, accessKey := range resp.AccessKeyMetadata {
		if *accessKey.AccessKeyId == accessKeyId {
			fmt.Println(*accessKey.UserName)
			return
		}
	}

	fmt.Fprintln(os.Stderr, "Nobody has the specified access key.")
	os.Exit(1)
}
