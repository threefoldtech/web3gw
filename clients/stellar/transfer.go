package stellargoclient

import (
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/support/errors"
	"github.com/stellar/go/txnbuild"
)

func (c *Client) Transfer(destination, memo string, amount string) error {
	accountRequest := horizonclient.AccountRequest{AccountID: c.kp.Address()}
	hAccount, err := c.horizon.AccountDetail(accountRequest)
	if err != nil {
		return errors.Wrap(err, "account does not exist")
	}

	if !hasTftTrustline(hAccount) {
		return errors.New("source account does not have trustline")
	}

	destAccountRequest := horizonclient.AccountRequest{AccountID: destination}
	destHAccount, err := c.horizon.AccountDetail(destAccountRequest)
	if err != nil {
		return errors.Wrap(err, "account does not exist")
	}

	if !hasTftTrustline(destHAccount) {
		return errors.New("destination account does not have trustline")
	}

	transferTx := txnbuild.Payment{
		Destination:   destination,
		Amount:        amount,
		Asset:         c.GetTftAsset(),
		SourceAccount: c.kp.Address(),
	}

	params := txnbuild.TransactionParams{
		SourceAccount:        &hAccount,
		IncrementSequenceNum: true,
		Operations:           []txnbuild.Operation{&transferTx},
		BaseFee:              txnbuild.MinBaseFee,
		Memo:                 txnbuild.MemoText(memo),
		Preconditions: txnbuild.Preconditions{
			TimeBounds: txnbuild.NewInfiniteTimeout(),
		},
	}
	tx, err := txnbuild.NewTransaction(params)
	if err != nil {
		return err
	}

	return c.SignAndSubmit(tx)
}
