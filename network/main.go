package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://lewiskori.com/")

	if err != nil {
		fmt.Println("Error is: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println((string(bs)))
	return len(bs), nil
}
