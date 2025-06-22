package main

import "log"

func main() {
	if err := run(); err != nil {
		log.Fatalf("runtime error: %v", err)
	}
}

// run is used in place of main as it is a normal go function that can return an error.
// having explicit inputs also allows for easier testing of the whole application.
func run() error {
	return nil
}
