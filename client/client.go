// Package client is HTTP-client for CRDT GCounters
package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
)

// TODO specify host and port

func GetCount() (int, error) {
	resp, err := http.Get("http://localhost:8001/gcounter/count")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	strBody := string(bytes.TrimSpace(body))
	intNumber, err := strconv.Atoi(strBody)
	if err != nil {
		return 0, err
	}
	return intNumber, nil
}

func Inc() error {
	resp, err := http.Get("http://localhost:8001/gcounter/increment")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
