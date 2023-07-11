!!chain.client.load
    network:dev
    mnemonic:'riot umbrella corn width treat before mention forest prison clarify fitness wise bird march such indoor swap rebuild flush office drastic enjoy grunt later'

!!chain.service_contract.create
    service:5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY
    consumer:5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY

!!chain.service_contract.approve
    contract_id:1015

!!chain.service_contract.reject
    contract_id:1015

!!chain.service_contract.bill
    contract_id: 2016
    variable_amount: 100

!!chain.service_contract.set
    contract_id:2015
    base_fee:100
    variable_fee:200

!!chain.service_contract.set
    metadata:'update metadata'

!!chain.service_contract.cancel
    contract_id:2015