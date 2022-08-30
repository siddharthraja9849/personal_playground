package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Errorf("%v+", err)
		os.Exit(1)
	}
	// bs := make([]byte, 99999)
	// N, err := resp.Body.Read(bs)
	// if err != nil {
	// 	fmt.Errorf("%v+", err)
	// }
	// fmt.Println(string(bs))
	// fmt.Println(N)

	myWriter := logWriter{}
	//io.Copy(os.Stdout, resp.Body)
	io.Copy(myWriter, resp.Body)

}
