package main

import (
	"testing"
)

func TestCheckLink(t *testing.T) {
	c := make(chan string)

	status, err := checkLink("http://google.com", c)
	if status != "up" || err != nil {
		t.Errorf("Expected 'up' for http://google.com but got %v with error: %v", status, err)
	}

	status, err = checkLink("http://invalid-link-for-test.xyz", c)
	if status != "down" || err == nil {
		t.Errorf("Expected 'down' for http://invalid-link-for-test.xyz but got %v", status)
	}
}
