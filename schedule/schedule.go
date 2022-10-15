// Package schedule contains schedule algorithms for access to distributed systems.
package schedule

type Scheduler[T any] interface {
	Next() T
}
