package stats

// Service ...
type Service interface {
	Measure() (slice []Statistics, err error)
}

// Statistics ...
type Statistics struct {
	Ident string
	Complexity int
}
