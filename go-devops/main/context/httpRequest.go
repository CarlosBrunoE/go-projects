package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Sends request to golang.org, and the context is cancelled if we does not have
// response in 3 seconds.
func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.golang.org", nil)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	// Attach to request
	req = req.WithContext(ctx)

	// Get our response
	resp, err := client.Do(req)
	cancel()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	// Print to stdout
	io.Copy(os.Stdout, resp.Body)
}
