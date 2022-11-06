package loader_test

import (
	"github.com/stretchr/testify/assert"
	"go-crdt-load-test/gcounter"
	"go-crdt-load-test/internal/app"
	"strconv"
	"testing"

	"go-crdt-load-test/loader"
)

// TestLoad is E2E test
// TODO write unit tests
func TestLoad(t *testing.T) {
	// Call 9 Increments and 1 Count, 10 times.
	// 9*10 = 90 is a total Increments counts.
	// 1*10 = 10 is a total Count calls count.
	// (9+1)*10 = 100 is a total requests count.
	loaderConfig := &loader.Config{
		CountsCount:      [3]int{10, 10, 1},
		IncsPerCountCall: [3]int{9, 9, 1},
		StartPort:        8000,
		EndPort:          8002,
	}
	err := app.Run(loaderConfig)
	assert.Nil(t, err)

	gc := gcounter.NewHttp("http://localhost:" + strconv.Itoa(loaderConfig.StartPort))

	count, err := gc.GetCount()
	assert.Nil(t, err)
	assert.Equal(t, 90, count)
}

// TODO unit test with stub loader.Loader.counter
