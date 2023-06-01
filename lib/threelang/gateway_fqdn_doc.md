!!tfgrid.login
	mnemonic: 'route visual hundred rabbit wet crunch ice castle milk model inherit outside'
	network: dev

!!tfgrid.gateway_fqdn.create 
    name: hamadafqdn
	node_id: 11
    backend: http://1.1.1.1:9000
    fqdn: hamada1.3x0.me

!!tfgrid.gateway_fqdn.get
	name: hamadafqdn

!!tfgrid.gateway_fqdn.delete
	name: hamadafqdn
