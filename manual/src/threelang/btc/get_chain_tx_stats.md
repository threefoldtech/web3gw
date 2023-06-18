# Get Chain Transaction Stats Action

> Returns statistics about the total number and rate of transactions in the chain. Providing the arguments will reduce the amount of blocks to calculate the statistics on.

- action name: !!btc.get.chain_tx_stats
- parameters:
  - amount_of_blocks [required]
    - provide statistics for amount_of_blocks blocks, if 0 for all blocks
  - block_hash_end [required]
    - provide statistics for amount_of_blocks blocks up until the block with the hash provided in block_hash_end

## Example

```md
    !!btc.get.chain_tx_stats
        amount_of_blocks: 123
        block_hash_end: 0x00000000000000000000000000000000
```
