package stellargoclient

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/protocols/horizon"
	"github.com/stellar/go/txnbuild"
)

func (c *Client) GenerateAccount() (*keypair.Full, error) {
	kp, err := keypair.Random()
	if err != nil {
		return nil, err
	}

	payment := txnbuild.Payment{
		SourceAccount: kp.Address(),
		Destination:   kp.Address(),
		Asset:         c.GetTftAsset(),
		Amount:        "2",
	}
	params := txnbuild.TransactionParams{
		SourceAccount: &horizon.Account{
			AccountID: kp.Address(),
		},
		IncrementSequenceNum: true,
		Operations:           []txnbuild.Operation{&payment},
		BaseFee:              0,
		Memo:                 nil,
		Preconditions: txnbuild.Preconditions{
			TimeBounds: txnbuild.NewInfiniteTimeout(),
		},
	}
	tx, err := txnbuild.NewTransaction(params)
	if err != nil {
		return nil, err
	}

	xdrJson, err := tx.ToXDR().MarshalBinary()
	if err != nil {
		return nil, err
	}
	base64EncodedXDR := base64.StdEncoding.EncodeToString(xdrJson)
	url := c.GetTransactionFundingUrlFromNetwork(c.stellarNetwork)
	postBody, _ := json.Marshal(map[string]string{
		"transaction": base64EncodedXDR,
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := map[string]string{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	errorMsg, ok := data["error"]
	if ok {
		return nil, errors.Errorf("%s", errorMsg)
	}
	base64EncodedXDR, ok = data["transaction_xdr"]
	if !ok {
		return nil, errors.Errorf("transaction_xdr not found in response from funding")
	}

	log.Debug().Msgf("XDR %s", base64EncodedXDR)

	fundingTx, err := txnbuild.TransactionFromXDR(base64EncodedXDR)
	if err != nil {
		return nil, err
	}

	tx, ok = fundingTx.Transaction()
	if !ok {
		return nil, err
	}

	xdrJson, err = tx.ToXDR().MarshalBinary()
	if err != nil {
		return nil, err
	}
	base64EncodedXDR = base64.StdEncoding.EncodeToString(xdrJson)
	log.Debug().Msgf("XDR after: %s", base64EncodedXDR)

	c.kp = kp

	err = c.SignAndSubmit(tx)
	if err != nil {
		return nil, err
	}

	return kp, nil
}

// Generates and activates an account on the stellar testnet
func (c *Client) GenerateAccount2() (*keypair.Full, error) {
	kp, err := keypair.Random()
	if err != nil {
		return nil, err
	}

	// Don't activate account on public network
	if c.stellarNetwork == "public" {
		// Todo create account via other account
		return kp, nil
	}

	err = activateAccount(kp.Address())
	if err != nil {
		return nil, err
	}

	err = c.SetTrustLine(kp.Address())
	if err != nil {
		return nil, err
	}

	return kp, nil
}

func (c *Client) getTrustLineOperation(account string) (*txnbuild.Transaction, error) {
	createTftTrustlineOperation := txnbuild.ChangeTrust{
		Line: txnbuild.ChangeTrustAssetWrapper{
			Asset: c.GetTftAsset(),
		},
		Limit:         "",
		SourceAccount: account,
	}

	// Get information about the account we just created
	accountRequest := horizonclient.AccountRequest{AccountID: account}
	hAccount, err := c.horizon.AccountDetail(accountRequest)
	if err != nil {
		return nil, errors.Wrap(err, "failed getting account detail from horizon")
	}

	log.Debug().Msgf("Account is %s", hAccount.ID)

	params := txnbuild.TransactionParams{
		SourceAccount:        &hAccount,
		IncrementSequenceNum: true,
		Operations:           []txnbuild.Operation{&createTftTrustlineOperation},
		BaseFee:              txnbuild.MinBaseFee,
		Memo:                 nil,
		Preconditions: txnbuild.Preconditions{
			TimeBounds: txnbuild.NewInfiniteTimeout(),
		},
	}
	return txnbuild.NewTransaction(params)
}

func (c *Client) SetTrustLine(account string) error {
	tx, err := c.getTrustLineOperation(account)
	if err != nil {
		return err
	}

	return c.SignAndSubmit(tx)
}

// Generates and activates accounts on the stellar testnet
func (c *Client) GenerateAndActivateAccounts(count int) ([]keypair.Full, error) {
	accounts := make([]keypair.Full, 0)
	for i := 0; i < count; i++ {
		kp, err := c.GenerateAccount()
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, *kp)
	}

	return accounts, nil
}

func (c *Client) SetAccountOptions(keypairs []keypair.Full) error {
	masterKey := keypairs[0]
	majority := (len(keypairs) / 2) + 1

	setOptionsOperations := make([]txnbuild.Operation, 0)
	for i := 1; i < len(keypairs); i++ {
		activeSigner := keypairs[i]
		setOptions := txnbuild.SetOptions{
			InflationDestination: nil,
			ClearFlags:           nil,
			SetFlags:             nil,
			MasterWeight:         txnbuild.NewThreshold(*txnbuild.NewThreshold(1)),
			LowThreshold:         txnbuild.NewThreshold(txnbuild.Threshold(0)),
			MediumThreshold:      txnbuild.NewThreshold(txnbuild.Threshold(majority)),
			HighThreshold:        txnbuild.NewThreshold(txnbuild.Threshold(majority)),
			HomeDomain:           nil,
			Signer:               &txnbuild.Signer{Address: activeSigner.Address(), Weight: 1},
			SourceAccount:        masterKey.Address(),
		}

		setOptionsOperations = append(setOptionsOperations, &setOptions)
	}

	// Get information about the account we just created
	accountRequest := horizonclient.AccountRequest{AccountID: masterKey.Address()}
	hAccount, err := c.horizon.AccountDetail(accountRequest)
	if err != nil {
		return err
	}

	params := txnbuild.TransactionParams{
		SourceAccount:        &hAccount,
		IncrementSequenceNum: true,
		Operations:           setOptionsOperations,
		BaseFee:              txnbuild.MinBaseFee,
		Memo:                 nil,
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

// Activates an account on the stellar testnet
// This is done by sending a request to the friendbot
func activateAccount(addr string) error {
	_, err := http.Get("https://friendbot.stellar.org/?addr=" + addr)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) AccountData(account string) (horizon.Account, error) {
	accountRequest := horizonclient.AccountRequest{
		AccountID: account,
	}

	return c.horizon.AccountDetail(accountRequest)
}
