package main

import "fmt"

type ValidateVpaResponse struct {
	Status       string `json:"status"`
	IsSuccess    bool   `json:"success"`
	ErrorMessage string `json:"error"`
}

func main() {
	resp := ValidateVpaResponse{
		IsSuccess: false,
	}
	fmt.Printf("%+v", resp)
}
