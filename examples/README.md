# Examples

This section of the repository contains examples of V scripts that use the V client library. They teach you how to use the V library. Inside each script you can find a function called $execute_rpcs$ in which you can add your calls to the specific V client. The first call should always be the one that creates the V client. You should give it the rpc websocket client.

```v
fn execute_rpcs(mut client RpcWsClient, mut logger log.Logger, secret string, network string) ! {
 mut stellar_client := stellar.new(mut client)

 stellar_client.load(secret: secret, network: network)!

 balance := stellar_client.balance("")! // fill in your address
 logger.info("My balance is: ${balance}")
}
```

## Running an example

`v -cg run <example>.v <args>`

## Adding new examples

You can use the file [main.v](main.v) as a template. Don't modify it though! The CI will build all v scripts inside the examples folder. CI should always be green so that we make sure that development builds at all times. 
