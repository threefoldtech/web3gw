!!web3gw.keys.define
    tfc_mnemonic:'mom picnic deliver again rug night rabbit music motion hole lion where'
    str_secret:'SAKRB7EE6H23EF733WFU76RPIYOPEWVOMBBUXDQYQ3OF4NF6ZY6B6VLW'


!!web3gw.money.send
    channel:tfchain
    to:28
    amount:100

!!web3gw.money.balance
    channel:tfchain
    currency:tft

!!web3gw.money.send
    channel:stellar
    bridge_to:tfchain
    to:28
    amount:100
