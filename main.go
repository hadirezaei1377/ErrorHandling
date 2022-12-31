package main

import (
	"fmt"
	"net/http"
	"time"
)

// Error handling strategies

// strategy 1, error_propagation
func error_propagation() (*http.Response, error) {

	var resp *http.Response

	// creat http Get by this anonymous function
	err := func() error {
		// call to virtual server by ftp protocol
		url := "ftp://anyhost.com"
		var err error
		resp, err = http.Get(url)
		if err != nil {
			// construct a new error message that includes reasoning and the underlying error

			return fmt.Errorf("failed to call the %s: %v", url, err)
		}
		return nil
	}()

	// more datails in another propagation
	if err != nil {
		// construct a new error message that includes reasoning and the underlying error
		return nil, fmt.Errorf("something went wrong: %v", err)
	}
	return resp, nil
}

// strategy 2, retry on error
func retry_on_error() (resp *http.Response, err error) {

	deadline := time.Now().Add(time.Second * 15)

	// for says about 15 seconds you must do works(make http get)    	for tries := 1; time.Now().Before(deadline); tries++ {
		resp, err = http.Get("ftp://anyhost.com")
		if err != nil {
			// exponential back-off
			time.Sleep(time.Second << uint(tries))
			continue
		}

		return resp, nil
	}

	// deadline is finished 
	return nil, fmt.Errorf("failed to call the API: %v", err)

	// strategy 3, log and exit
	func log_and_exit() *http.Response {
		resp, err := http.Get("ftp://anyhost.com")
		if err != nil {
			log.Fatalf("failed to call the API %v. exit", err)
		}
	
		return resp
	}
	