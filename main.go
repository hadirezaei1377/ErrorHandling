package main

import (
	"fmt"
	"net/http"
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
