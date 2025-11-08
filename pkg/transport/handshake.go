package transport

// HandShakeFunc is ... ?
type HandShakeFunc func(node Node) error

func NOPHandShaker(node Node) error {
	return nil
}
