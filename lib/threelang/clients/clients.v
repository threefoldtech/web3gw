module clients

import threefoldtech.threebot.tfgrid { TFGridClient }
import threefoldtech.threebot.tfchain { TfChainClient }
import threefoldtech.threebot.stellar { StellarClient }
import threefoldtech.threebot.eth { EthClient }
import threefoldtech.threebot.btc { BtcClient }

[heap]
pub struct Clients {
pub mut:
	tfg_client TFGridClient
	tfc_client TfChainClient
	str_client StellarClient
	eth_client EthClient
	btc_client BtcClient
}

