# Stellar

## Creating an account

You can create a stellar account through the rpc stelar.CreateAccount which only requires one argument: the network the account should be created on. The following rpc should be send to the server to create your account. 

```json
{
    "jsonrpc":"2.0",
    "method":"stellar.CreateAccount",
    "params":[
        "public"
    ],
    "id":"a_unique_id_here"
}
```

It will return the seed in the result argument of the json rpc 2.0 response:

```json
{
    "jsonrpc":"2.0",
    "result":"seed_will_be_here",
    "id":"id_send_in_request"
}
```

## Loading your key

Before you can execute any other rpc (except for CreateAccount) you have to call the stellar.Load rpc. This will start your session and requires you to specify your secret and the network you want to connect to. Below you can find an example of the json request to send to the server. 

```json
{
    "jsonrpc":"2.0",
    "method":"stellar.Load",
    "params":[{
        "network":"public",
        "secret":"SA33FBB67CPIMHWTZYVR489Q6UKHFUPLKTLPG9BKAVG89I2J3SZNMW21"
    }],
    "id":"a_unique_id_here"
}
```

If everything went well it will return an empty json rpc 2.0 response:

```json
{
    "jsonrpc":"2.0",
    "id":"id_send_in_request"
}
```

## Asking your public address

Once your secret key has been loaded you can ask for your public address via the stellar.Address rpc:

```json
{
    "jsonrpc":"2.0",
    "method":"stellar.Address",
    "params":[],
    "id":"a_unique_id_here"
}
```

This will return the public address in a json rpc 2.0 response. 

```json
{
    "jsonrpc":"2.0",
    "result":"public_address_will_be_here",
    "id":"id_send_in_request"
}
```

## Transfer tokens from one account to another

The rpc stellar.Transfer allows you to transfer tokens from one account to the other. It requires the amount (json string), the destination account and a memo. 

```json
{
    "jsonrpc":"2.0",
    "method":"stellar.Transfer",
    "params":[{
        "amount": "1520.0",
        "destination": "some_public_stellar_address",
        "memo": "your_memo_comes_here"
    }],
    "id":"a_unique_id_here"
}
```

This will return the hash of the transaction in a json rpc 2.0 response.

```json
{
    "jsonrpc":"2.0",
    "result":"hash_will_be_here",
    "id":"id_send_in_request"
}
```

## Swap tokens from one asset to the other

It is possible to swap your lumen tokens to tft tokens and visa versa with the rpc stellar.Swap. I requires the amount, the source asset and the destination asset. Both the source and destination assets should be one of tft or xlm meaning we can only swap tfts to lumen and visa versa. Here is how you can swap 5 tft on stellar to lumen:

```json
{
    "jsonrpc":"2.0",
    "method":"stellar.Swap",
    "params":[{
        "amount": "5.0",
        "source_asset": "tft",
        "destination_asset": "xlm"
    }],
    "id":"a_unique_id_here"
}
```

This will return the hash of the transaction in a json rpc 2.0 response.

```json
{
    "jsonrpc":"2.0",
    "result":"hash_will_be_here",
    "id":"id_send_in_request"
}
```

## Get the balance of an account

You can as for the balance of an account with an rpc to stellar.Balance. The account to ask the balance of can be send via the params. If it is not present in the params the balance will be returned of the account that is currently loaded. 

```json
{
    "jsonrpc":"2.0",
    "method":"stellar.Balance",
    "params":[
        "you_can_pass_public_address_here"
    ],
    "id":"a_unique_id_here"
}
```

The response will be:

```json
{
    "jsonrpc":"2.0",
    "result":"balance_will_be_here",
    "id":"id_send_in_request"
}
```

## Bridge stellar tft to ethereum

You can convert your TFTs on stellar to TFTs on ethereum with the call to stellar.BridgeToEth. Below you can find an example that will move 298 TFT to the ethereum account defined by the param destination. 

```json
{
    "jsonrpc":"2.0",
    "method":"stellar.BridgeToEth",
    "params":[{
        "amount": "298.0",
        "destination": "eth_public_address_here"
    ],
    "id":"a_unique_id_here"
}
```

The response of the server will contain the hash of the transaction on the bridge:

```json
{
    "jsonrpc":"2.0",
    "result":"hash_will_be_here",
    "id":"id_send_in_request"
}
```

## Bridge stellar tft to tfchain

Similarly to prior call you can convert your TFTs on stellar to TFTs on tfchain with the call to stellar.BridgeToTfchain. The example below transfers 21 TFT from your loades stellar account to the account on tfchain that belongs to the twin id 122. 

```json
{
    "jsonrpc":"2.0",
    "method":"stellar.BridgeToTfchain",
    "params":[{
        "amount": "21.0",
        "twin_id": 122
    ],
    "id":"a_unique_id_here"
}
```

The response of the server will contain the hash of the transaction on the bridge.

```json
{
    "jsonrpc":"2.0",
    "result":"hash_will_be_here",
    "id":"id_send_in_request"
}
```

## Waiting for a transaction on the Ethereum bridge

This is a usefull call in case you are bridging from ethereum to stellar (transfering tft from ethereum to stellar). Once the ethereum has been transferred to the bridge it's time to wait till the bridge executes a similar transaction on stellar (so that the stellar account is receiving the tfts send from the ethereum account). This is what the rpc stellar.AwaitTransactionOnEthBridge does. It requires the memo of the transaction which is predicatble when executing the call to bridge the tokens. 

```json
{
    "jsonrpc":"2.0",
    "method":"stellar.BridgeToTfchain",
    "params":[
        "provide_the_memo_here"
    ],
    "id":"a_unique_id_here"
}
```

If the transaction is found in the next 5 minutes it will return an empty result.

```json
{
    "jsonrpc":"2.0",
    "id":"id_send_in_request"
}
```

## Listing transactions

You can get the transactions of an account via the stellar.Transactions rpc. Just provide the account, the limit, whether or not to include failed transactions, the cursor and the order the transactions should be in. All of these params are optional. Leaving the account empty will result in showing the transactions of the account that is loaded. There is a default limit of 10 transactions, the default cursor is the top (latest), by default we do not show failed transactions and the transactions are by default in descending order. 

```json
{
    "jsonrpc":"2.0",
    "method":"stellar.Transactions",
    "params":[{
        "account": "some_account_here_or_leave_empty",
        "limit": 12,
        "include_failed": false,
        "cursor": "leave_empty_for_top",
        "ascending": false
    ],
    "id":"a_unique_id_here"
}
```

If the transaction is found in the next 5 minutes it will return an empty result.

```json
{
    "jsonrpc":"2.0",
    "result":[
        {
            "id": "some_id",
            // many more attributes see https://github.com/stellar/go/blob/01c7aa30745a56d7ffcc75bb8ededd38ba582a58/protocols/horizon/main.go#L484
        }
    ],
    "id":"id_send_in_request"
}
```

## Showing the data related to an account

You can get more information about your account (or any other account) through the rpc stellar.AccountData. It has only one parameter: the account to get the data for. 

```json
{
    "jsonrpc":"2.0",
    "method":"stellar.AccountData",
    "params":[
        "account_or_leave_empty_for_your_account"
    ],
    "id":"a_unique_id_here"
}
```

The account data will be in the result if the account exists:

```json
{
    "jsonrpc":"2.0",
    "result": {
        "id": "some_id",
        // many more attributes see https://github.com/stellar/go/blob/01c7aa30745a56d7ffcc75bb8ededd38ba582a58/protocols/horizon/main.go#L33
    },
    "id":"id_send_in_request"
}
```
