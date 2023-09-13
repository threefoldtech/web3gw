module tfgrid

import rand
import encoding.base64

const key_len = 32 // KeyLen is the expected key length for a WireGuard key.

// GenerateKey generates a Key suitable for use as a pre-shared secret key from
// a cryptographically safe source.
//
// The output Key should not be used as a private key; use GeneratePrivateKey
// instead.

pub fn generate_key() []u8 {
	mut res := []u8{len: tfgrid.key_len}
	rand.read(mut res)
	return res
}


// GeneratePrivateKey generates a Key suitable for use as a private key from a
// cryptographically safe source.
pub fn generate_private_key() []u8 {
	mut key := generate_key()

	// Modify random bytes using algorithm described at:
	// https://cr.yp.to/ecdh.html.

	key[0] &= 248
	key[31] &= 127
	key[31] |= 64

	return key
}

// ParseKey parses a Key from a base64-encoded string, as produced by the
// Key.String method.
pub fn parse_key(key string) []u8 {
	return base64.decode(key)
}
