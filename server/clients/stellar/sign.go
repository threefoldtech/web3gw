package stellargoclient

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/txnbuild"
)

func (c *Client) SignTransactionXdr(txXdr string) error {
	txn, err := txnbuild.TransactionFromXDR(txXdr)
	if err != nil {
		return errors.Wrap(err, "failed to create transaction from xdr")
	}
	tx, ok := txn.Transaction()
	if !ok {
		return errors.Wrap(err, "failed converting general transaction to transaction")
	}
	return c.SignAndSubmit(tx)
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

func (c *Client) SignFundAndSubmitTransaction(tx *txnbuild.Transaction) error {
	tx, err := tx.Sign(c.getNetworkPassPhrase(), c.kp)
	if err != nil {
		return errors.Wrap(err, "failed to sign transaction")
	}
	xdr, err := tx.Base64()
	if err != nil {
		return errors.Wrap(err, "failed to convert transaction to xdr")
	}

	url := c.GetTransactionFundingUrlFromNetwork() + "/fund_transaction"
	binaryPostdata, err := json.Marshal(map[string]string{
		"transaction": xdr,
	})
	if err != nil {
		return errors.Wrap(err, "failed to marshal data")
	}
	postDataReader := bytes.NewBuffer(binaryPostdata)

	resp, err := http.Post(url, "application/json", postDataReader)
	if err != nil {
		return errors.Wrap(err, "failed to send http post request")
	}
	data := map[string]string{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return errors.Wrap(err, "failed to decode response from http post")
	}
	if errorMsg, errorPresent := data["error"]; errorPresent {
		return errors.New(errorMsg)
	}
	return nil
}
