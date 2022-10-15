package client_test

import (
	"testing"

	"go-crdt-load-test/client"

	"github.com/stretchr/testify/assert"
)

const address = "http://localhost:8001"

// TODO write unit tests. Now it is integration tests

func TestGetCount(t *testing.T) {
	count, err := client.GetCount(address)
	assert.Nil(t, err)
	assert.Equal(t, 1, count)
}

func TestInc(t *testing.T) {
	err := client.Inc(address)
	assert.Nil(t, err)
}
