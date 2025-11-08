package transport

// HandShakeFunc is handler for handshake
type HandShakeFunc func(node Node) error

func NOPHandShaker(node Node) error {
	return nil
}
