package stellargoclient

import (
	"github.com/pkg/errors"
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/txnbuild"
)

func (c *Client) SignTransactionXdr(txXdr string) error {
	return nil
}

func (c *Client) SignAndSubmit(txn *txnbuild.Transaction) error {
	// Sign the transaction, and base 64 encode its XDR representation
	signedTx, err := txn.Sign(c.GetStellarNetworkPassphrase(), c.kp)
	if err != nil {
		return errors.Wrap(err, "failed to sign transaction")
	}

	txeBase64, err := signedTx.Base64()
	if err != nil {
		return errors.Wrap(err, "failed to base64 encode transaction")
	}

	// Submit the transaction
	_, err = c.horizon.SubmitTransactionXDR(txeBase64)
	if err != nil {
		hError := err.(*horizonclient.Error)
		return hError
	}

	return nil
}
