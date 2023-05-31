module solution

import threefoldtech.threebot.tfgrid { TFGridClient }
import threefoldtech.threebot.explorer { ExplorerClient }

pub struct SolutionHandler {
pub mut:
	tfclient &TFGridClient
	explorer &ExplorerClient
}
