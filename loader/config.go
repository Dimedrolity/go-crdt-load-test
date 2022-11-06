package loader

// Config for Loader.
// Call IncsPerCountCall Increments and 1 Count, CountsCount times.
// IncsPerCountCall*CountsCount is a total Increments counts.
// CountsCount is a total Count calls count.
// (IncsPerCountCall+1)*CountsCount is a total requests count.
type Config struct {
	// [start, stop, step]
	CountsCount      [3]int `yaml:"counts_count"`
	IncsPerCountCall [3]int `yaml:"incs_per_count_call"`

	StartPort int `yaml:"start_port"`
	EndPort   int `yaml:"end_port"`
}
