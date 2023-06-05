module threelang

import freeflowuniverse.crystallib.actionsparser

pub fn (mut r Runner) helper_actions(mut ap actionsparser.ActionsParser) ! {
	mut sshkey_action := ap.filtersort(actor: 'sshkeys', book: 'tfgrid')!

	for a in sshkey_action {
		name := a.params.get('name')!
		key := a.params.get('ssh_key')!
		r.ssh_keys[name] = key
	}
}
