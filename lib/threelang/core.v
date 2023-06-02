module threelang

import freeflowuniverse.crystallib.actionsparser
import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution { Capacity, SolutionHandler, VM }
import log

fn (mut r Runner) core_actions(actions actionsparser.ActionsParser) ! {
	mut actions2 := actions.filtersort(actor: 'core', domain:'tfgrid')!
	for action in actions2 {
		p:=action.params
		if action.name == 'login' {
			//remember  on runner the client and login info for explorer/tfchain

			mnemonic := action.params.get_default('mnemonic', '')!
			//check env called TFGRID_MNEMONIC ??? if set use that one if 

			netstring := action.params.get_default('net', 'dev')!
			//check env called TFGRID_NET ??? if set use that one if 

		}

	}
}