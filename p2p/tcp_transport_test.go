package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAddr := ":4344"
	// Create a new TCPTransport instance
	tr := NewTCPTransport(listenAddr)
	assert.Equal(t, tr.listenAddress, listenAddr, "Expected listen address to be %s", listenAddr)

	// server
	assert.Nil(t, tr.ListenAndAccept(), "Expected no error when starting to listen")
	select {}
}
