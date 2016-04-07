package main

import (
	"fmt"
	"os"
)

const Usage = `
Usage:
  $ whosekey <AWS_ACCESS_KEY_ID>
`

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, Usage)
		os.Exit(1)
	}

	access_key_id := os.Args[1]
	fmt.Println(access_key_id)
}
