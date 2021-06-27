package main

import (
	"context"
	"errors"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambdaHandler)
}

//	Handles the incoming request from trigger
func lambdaHandler(ctx context.Context, request sumRequest) (resp interface{}, err error) {
	total, err := sum(request.ParamOne, request.ParamTwo)
	if err != nil {
		return
	}
	return &sumResponse{
		Total: total,
	}, nil
}

//	Returns sum of the parameters provided and error otherwise.
func sum(paramOne int, paramTwo int) (total int, err error) {
	if paramOne < 1 || paramTwo < 1 {
		err = errors.New("Invalid parameter/s")
		return
	}

	total = paramOne + paramTwo
	return
}

type sumRequest struct {
	ParamOne int `json:"param_one"`
	ParamTwo int `json:"param_two"`
}

type sumResponse struct {
	Total int `json:"total"`
}
