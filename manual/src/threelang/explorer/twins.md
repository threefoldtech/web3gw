# Twins Action
Query and filter twins on the chain.

## Filter twins
- action name: !!explor.twins.filter
- parameters:
    - `twin_id`: twin id
    - `account_id`: twin account address
    - `relay`: relay address of the twin
    - `public_key`: twin public key


    - `size`: size of the returned batch of the twins. default is 50
    - `page`: offset of the returned batch of the twins. default is 1
    - `randomize`: if true, the returned batch of twins will be random. default is false
    - `count`: if true, will return the number of twins that match the filter even if the size is set. default is false.

- examples:
    - get specific twin by it's id
        ```bash
        !!explor.twins.filter
            twin_id: 29
        ```
    - get twin by it's account address
        ```bash
        !!explor.twins.filter
            account_id:'5FiC58mQ3J8dbfpUwDvSxYAgnW5uibmubJoATMFwkT6tC2Sn'
        ```
