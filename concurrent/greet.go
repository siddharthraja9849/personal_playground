package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	websites := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
	}

	for _, web := range websites {
		getStatusCode(web)
	}

}

// func greeting(s string) {
// 	for i := 0; i < 6; i++ {
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Print("OOPS")
	}

	fmt.Printf("%d status code for %s %s\n", res.StatusCode, res.Status, endpoint)
}
