module explorer

import freeflowuniverse.crystallib.actionparser { Action }
import threefoldtech.threebot.tfgrid { Limit, StatsFilter }

fn (mut h ExplorerHandler) counters(action Action) ! {
	match action.name {
		'get' {
			status := action.params.get_default('status', '')!

			req := StatsFilter{
				status: status,
			}

			res := h.explorer.counters(req)!
			h.logger.info('counters: ${res}')
		}
		else {
			return error('unknown action: ${action.name}')
		}
	}
}