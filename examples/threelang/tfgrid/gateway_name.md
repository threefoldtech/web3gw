# Gatewya Name Example

- This example deployes, gets, and deletes a gateway name workload on the tfgrid.

!!tfgrid.core.login
	mnemonic: 'YOUR MNEMONICS'
	network: dev

!!tfgrid.gateway_name.create 
	name: hamadagateway
	backend: http://1.1.1.1:9000

!!tfgrid.gateway_name.get
	name: hamadagateway

!!tfgrid.gateway_name.delete
	name: hamadagateway