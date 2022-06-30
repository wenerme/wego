package chanx

// IsClosed returns true if the channel is closed.
func IsClosed[T any](ch <-chan T) bool {
	select {
	case <-ch:
		return true
	default:
	}
	return false
}

// TryReceive receives a value from a channel and returns true if the channel is closed.
func TryReceive[T any](ch <-chan T) (value T, closed bool) {
	select {
	case value = <-ch:
	default:
		closed = true
	}
	return
}

// TryClose closes a channel and returns true if the channel is closed.
func TryClose[T any](ch chan T) (justClosed bool) {
	defer func() {
		if recover() != nil {
			justClosed = false
		}
	}()

	close(ch)
	return true
}

// TrySend sends a value to a channel and returns true if the channel is closed.
func TrySend[T any](ch chan<- T, value T) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = true
		}
	}()
	ch <- value
	return false
}
