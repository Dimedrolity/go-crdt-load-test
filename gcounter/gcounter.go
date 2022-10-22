// Package gcounter is a Grow-Only Counter CRDT.
package gcounter

// GCounter is abstraction of Grow-Only Counter.
type GCounter interface {
	GetCount(address string) (int, error)
	Inc(address string) error
}
