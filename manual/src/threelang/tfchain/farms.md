# Farms
## Get
Get farms info by provideing on of the parameters.
- action name: !!chain.farms.get
- parameters:
    - `id`: [optional] farm id
    - `name`: [optional] the full farm name "case-sensitive"
- example:
    ```
    !!chain.farms.get
        id:1
    ```
    ```
    !!chain.farms.get
        name:freefarm
    ```
## Create
Create new farm on the chain. 
- action name: !!chain.farms.create
- parameters:
    - `name`: [required] the full farm name "case-sensitive"
    - `public_ips`: [optional] is a list of IP addresses in CIDR format xxx.xxx.xxx.xxx/xx. separated by comma.
    - `gateways`: [optional] is a list of Gateways for the IP in ipv4 format. separated by comma.
- example:
    ```md
    !!chain.farms.create
        name:newfarm
    ```
    ```
    !!chain.farms.create
        name:newfarmwithips
        public_ips: 185.206.122.152/16
        gateways: 185.206.122.152
    ```
    Values on both public_ips/gateways lists are mapped by their indices. so for example gateways[0] is the gateway for the ip public_ips[0] and so on.