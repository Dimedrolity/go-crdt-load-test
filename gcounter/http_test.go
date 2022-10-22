package gcounter_test

import (
	"testing"

	"go-crdt-load-test/gcounter"

	"github.com/stretchr/testify/assert"
)

const address = "http://localhost:8001"

// TODO write unit tests. Now it is integration tests

func TestGetCount(t *testing.T) {
	count, err := gcounter.Http{}.GetCount(address)
	assert.Nil(t, err)
	assert.Equal(t, 1, count)
}

func TestInc(t *testing.T) {
	err := gcounter.Http{}.Inc(address)
	assert.Nil(t, err)
}
