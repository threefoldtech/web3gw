module threelang

import freeflowuniverse.crystallib.actionsparser
import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution { SolutionHandler }
import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import threefoldtech.threebot.explorer

fn (mut r Runner) core_actions(mut actions actionsparser.ActionsParser, mut myclient RpcWsClient) ! {
	mut actions2 := actions.filtersort(actor: 'core', book: 'tfgrid')!
	for action in actions2 {
		p := action.params
		if action.name == 'login' {
			mnemonic := action.params.get_default('mnemonic', '')!
			netstring := action.params.get_default('net', 'dev')!

			mut tfgrid_client := tfgrid.new(mut myclient)
			mut exp := explorer.new(mut myclient)

			tfgrid_client.load(tfgrid.Credentials{
				mnemonic: mnemonic
				network: netstring
			})!

			exp.load(netstring)!

			mut sl := SolutionHandler{
				tfclient: tfgrid_client
				explorer: exp
			}

			r.handler = sl
		}
	}
}
