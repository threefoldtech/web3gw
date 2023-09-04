package pkg

type (
	// ErrClientNotConnected indicates a client is not yet connected to an node and or the client does not have a private key loaded yet.
	ErrClientNotConnected struct{}
)

// Error implements Error interface
func (e ErrClientNotConnected) Error() string {
	return "client not connected yet"
}
