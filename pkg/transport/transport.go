package transport

// Node represents the remote node
type Node interface {
}

// Transport that handles the connection between nodes in the network
// this can be (TCP, UDP, WebRTC)
type Transport interface {
	ListenAndAccept() error
}
