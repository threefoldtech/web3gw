module nostr

pub struct Event {
	id        string
	pubkey    string
	created_at u64
	kind      int
	tags      []string
	content   string
	sig       string

	// extra map[string]any
}