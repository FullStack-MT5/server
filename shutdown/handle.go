package shutdown

// Handle is a configurable shutdown handle that executes a
// custom shutdown procedure.
type Handle struct {
	handleFunc func() error
}

func NewHandle(f func() error) *Handle {
	return &Handle{
		handleFunc: f,
	}
}

// Call calls the registered handlFunc on Handle.
// Invoking Call is noop if handlFunc is nil.
func (h *Handle) Call() error {
	if h.handleFunc == nil {
		return nil
	}
	return h.handleFunc()
}
