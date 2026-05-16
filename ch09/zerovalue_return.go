package main

import (
	"errors"
	"fmt"
)

type config struct {
	host string
}

func parseConfig(payload string) (*config, error) {
	if payload == "" {
		return &config{}, errors.New("empty payload")

	}
	return &config{host: payload}, nil
}

func pingServer() (string, interface{}) {
	return "ping", nil
}

func validateConfig() error {
	return errors.New("not implemented")
}

func fetchConfig() (error, error) {
	return errors.New("not implemented"), nil
}

func main() {
	_, err := parseConfig("")
	if err != nil {
		fmt.Println(err)
	}
}
