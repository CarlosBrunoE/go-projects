package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

var ranOnce bool

const (
	UnknownCode     = 0
	UnreachableCode = 1
	AuthFailureCode = 2
)

// ErrNetwork is a custom network error that had error codes
// and a message.
type ErrNetwork struct {
	Code int
	Msg  string
}

func (e ErrNetwork) Error() string {
	return fmt.Sprintf("network error(%d): %s", e.Code, e.Msg)
}

// someFunc is a standin for some function that takes in some data.
// On the first call it will return an ErrNetwork and the second just
// a standard error.
func someFunc(data string) error {
	if !ranOnce {
		ranOnce = true
		return ErrNetwork{
			Code: AuthFailureCode,
			Msg:  "user unrecognized",
		}
	}
	return fmt.Errorf("another error")
}

func restCall(data string) error {
	if err := someFunc(data); err != nil {
		// %w is used to wrap the original error with our new one
		return fmt.Errorf("restCall %s had an error %w", data, err)
	}
	return nil
}

func main() {
	// This loop will continue as long as we have network errors that are
	// not code AuthFailureCode and will  terminate on any other error or success.
	for {
		if err := restCall("data"); err != nil {
			var netErr ErrNetwork
			if errors.As(err, &netErr) {
				log.Println("network error: ", err)
				time.Sleep(1 * time.Second)
				continue
			}
			log.Println("unrecoverable: ", err)
		}
	}
}
