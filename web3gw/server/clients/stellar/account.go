package stellargoclient

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/protocols/horizon"
	"github.com/stellar/go/txnbuild"
)

func (c *Client) Load(secret string) error {
	k, err := GetKeypairFromSeed(secret)
	if err != nil {
		return err
	}
	c.kp = k

	// check if account has trustline, if not add it
	hAccount, err := c.AccountData(k.Address())
	if err != nil {
		return errors.Wrap(err, "account does not exist")
	}

	if !hasTrustline(hAccount, c.GetTftBaseAsset()) {
		log.Debug().Msgf("Adding trustline for account %s", k.Address())
		c.setTrustLine()
	}

	return nil
}

func (c *Client) CreateAccount() (string, error) {
	kp, err := keypair.Random()
	if err != nil {
		return "", err
	}

	c.kp = kp

	err = c.activateAccount()
	if err != nil {
		return "", err
	}

	err = c.setTrustLine()
	if err != nil {
		return "", err
	}

	return kp.Seed(), nil
}

// Activates the account using the activation service https://github.com/threefoldfoundation/tft-stellar/tree/master/ThreeBotPackages/activation_service
func (c *Client) activateAccount() error {
	url := c.GetActivationServiceUrl()
	binaryPostdata, err := json.Marshal(map[string]string{
		"address": c.kp.Address(),
	})
	if err != nil {
		return errors.Wrap(err, "failed Mashal data")
	}
	postDataReader := bytes.NewBuffer(binaryPostdata)
	resp, err := http.Post(url+"/activate_account", "application/json", postDataReader)
	if err != nil {
		return errors.Wrap(err, "failed sending post")
	}
	responseData := ""
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		return err
	}
	responseDataJson := map[string]string{}
	err = json.Unmarshal([]byte(responseData), &responseDataJson)
	if err != nil {
		return errors.Wrap(err, "failed decoding http post response")
	}
	if errorMsg, errorPresent := responseDataJson["error"]; errorPresent {
		return errors.New(errorMsg)
	}

	xdr, ok := responseDataJson["activation_transaction"]
	if !ok {
		return errors.Errorf("activation_transaction not found in response")
	}
	return c.SignTransactionXdr(xdr)
}

// Sets trustline using the activation service https://github.com/threefoldfoundation/tft-stellar/tree/master/ThreeBotPackages/activation_service
func (c *Client) setTrustLine() error {
	url := c.GetActivationServiceUrl()
	asset := c.GetTftAsset()
	binaryPostdata, err := json.Marshal(map[string]string{
		"address": c.kp.Address(),
		"asset":   asset.Code + ":" + asset.Issuer,
	})
	if err != nil {
		return errors.Wrap(err, "failed Mashal data")
	}
	postDataReader := bytes.NewBuffer(binaryPostdata)
	resp, err := http.Post(url+"/fund_trustline", "application/json", postDataReader)
	if err != nil {
		return errors.Wrap(err, "failed sending post")
	}
	responseDataJson := map[string]string{}
	err = json.NewDecoder(resp.Body).Decode(&responseDataJson)
	if err != nil {
		return errors.Wrap(err, "failed decoding http post response")
	}
	if errorMsg, errorPresent := responseDataJson["error"]; errorPresent {
		return errors.New(errorMsg)
	}
	xdr, ok := responseDataJson["addtrustline_transaction"]
	if !ok {
		return errors.Errorf("addtrustline_transaction not found in response")
	}
	return c.SignTransactionXdr(xdr)
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
	hAccount, err := c.AccountData(masterKey.Address())
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

func (c *Client) AccountData(account string) (horizon.Account, error) {
	accountRequest := horizonclient.AccountRequest{
		AccountID: account,
	}

	return c.horizon.AccountDetail(accountRequest)
}
