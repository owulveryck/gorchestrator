package orchestrator

// Executor is a type that describes an executor backend
type Executor struct {
	Url string // The url of the executor
}

// SetUrl sets the executor URL but does not check for its validity
func (e *Executor) SetUrl(u string) error {
	return nil
}
