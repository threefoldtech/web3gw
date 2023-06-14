# Service Contract

## Create
Create new service contract
- action name: !!tfchain.service_contract.create
- parameters:
    - `service`: [required] SS58 address
    - `consumer`: [required] SS58 address
- example:
    ```
    !!tfchain.service_contract.create
        service:  `string`
        consumer: `string`    
    ```

## Approve
Approve service contract
- action name: !!tfchain.service_contract.approve
- parameters:
    - `contract_id`: [required] contract id
- example:
    ```
    !!tfchain.service_contract.approve
        contract_id: 2015
    ```

## Reject
Reject service contract
- action name: !!tfchain.service_contract.reject
- parameters:
    - `contract_id`: [required] contract id
- example:
    ```
    !!tfchain.service_contract.reject
        contract_id:2015
    ```

## Bill
Bill service contract
- action name: !!tfchain.service_contract.bill
- parameters:
    - `contract_id`: [required] contract id
    - `variable_amount`: [required] amount of variable fee
- example:
    ```
    !!tfchain.service_contract.bill
        contract_id:2015     
        variable_amount:100
    ```
## Update
Update service contract by setting fees or metadata
- action name: !!tfchain.service_contract.set
- parameters:
    - `contract_id`: [required] contract id
    - `base_fee`: amount of base fee
    - `variable_fee`: amount of variable fee
    - `metadata`: metadata
- example:
    - Set fees for service contract
        ```
        !!tfchain.service_contract.set
            contract_id:2015
            base_fee:      `u64`
            variable_fee:  `u64`
        ```

    - Set metadata for service contract
        ```
        !!tfchain.service_contract.set
            metadata: string
        ```

## Cancel
Cancel service contract 
- action name: !!tfchain.service_contract.cancel
- parameters:
    - `contract_id`: [required] contract id
- example:
    ```
    !!tfchain.service_contract.cancel
        contract_id:2015
    ```