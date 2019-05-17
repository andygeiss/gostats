package stats

// Service ...
type Service interface {
	Measure() (result []Statistics, err error)
}

// Statistics ...
type Statistics struct {
	Ident      string
	Complexity int
	Nodes      int
}
