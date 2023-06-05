module threelang

import freeflowuniverse.crystallib.actionsparser
import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution { SolutionHandler }
import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import threefoldtech.threebot.explorer

fn (mut r Runner) core_actions(mut actions actionsparser.ActionsParser) ! {
	mut actions2 := actions.filtersort(actor: 'core', book: 'tfgrid')!

	for action in actions2 {
		if action.name == 'login' {
			mnemonic := action.params.get_default('mnemonic', '')!
			netstring := action.params.get_default('net', 'dev')!

			r.handler.tfclient.load(tfgrid.Credentials{
				mnemonic: mnemonic
				network: netstring
			})!

			r.handler.explorer.load(netstring)!
		}
	}
}
