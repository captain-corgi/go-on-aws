package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
)

func main() {
	fmt.Println("Hello, AWS!")

	// String pointer
	aStrPt := aws.String("Hi")
	fmt.Println(aStrPt)
	fmt.Println(*aStrPt)
}
