package gcounter

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Http is GCounter implementation that use HTTP requests to use counter https://github.com/Dimedrolity/gcounter-crdt.
// The zero value is ready to use.
type Http struct{}

// GetCount gets actual counter value synchronized with all nodes.
func (Http) GetCount(address string) (int, error) {
	resp, err := http.Get(address + "/gcounter/count")
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

// Inc increments counter value on one node.
func (Http) Inc(address string) error {
	resp, err := http.Get(address + "/gcounter/increment")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
