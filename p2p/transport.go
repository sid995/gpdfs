package p2p

// Peer represents a remote node in the network.
// It is used to identify and communicate with other nodes.
// The Peer interface is a placeholder for the actual implementation
// of a peer in the network. It can be extended to include methods
// for sending and receiving messages, managing connections, etc.
type Peer interface{}

// Transport represents the underlying transport layer used for
// communication between peers in the network.
// It is responsible for establishing connections, sending and
// receiving messages, and managing the transport protocol.
type Transport interface {
	ListenAndAccept() error
}
