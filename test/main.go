package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Message string
	Ok      bool
}

func Handler() (Response, error) {
	fmt.Println("===> START: Handler()")
	return Response{
		Message: "success",
		Ok:      true,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
