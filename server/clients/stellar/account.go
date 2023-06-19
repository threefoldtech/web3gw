package stellargoclient

import (
	"bytes"
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

func (c *Client) getTFTTransactionFundingCondition() (feeWallet, fee string, err error) {
	baseUrl := c.GetTransactionFundingUrlFromNetwork()
	resp, err := http.Get(baseUrl + "/conditions")
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	type Condition struct {
		Asset          string
		Fee_account_id string
		Fee_fixed      string
	}
	conditions := make([]Condition, 0)
	json.Unmarshal(body, &conditions)
	tftAsset := c.GetTftAsset()
	tftAssetString := tftAsset.GetCode() + ":" + tftAsset.GetIssuer()
	for _, c := range conditions {
		if c.Asset == tftAssetString {
			feeWallet = c.Fee_account_id
			fee = c.Fee_fixed
			return
		}
	}
	err = errors.New("Transaction funding condition for TFT not found")
	return
}
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

func (c *Client) GenerateAccount() (*keypair.Full, error) {
	kp, err := keypair.Random()
	if err != nil {
		return nil, err
	}

	c.kp = kp

	err = c.activateAccount()
	if err != nil {
		return nil, err
	}

	err = c.setTrustLine()
	if err != nil {
		return nil, err
	}
	return kp, nil
}

func (c *Client) FundTransactionUsingTft(amount string) error {
	feeAccount, fee, err := c.getTFTTransactionFundingCondition()
	if err != nil {
		return errors.Wrap(err, "failed to get tft transaction funding condition")
	}

	payment := txnbuild.Payment{
		Destination: c.kp.Address(),
		Amount:      amount,
		Asset:       c.GetTftAsset(),
	}
	feePayment := txnbuild.Payment{
		Destination: feeAccount,
		Amount:      fee,
		Asset:       c.GetTftAsset(),
	}
	sourceAccount, err := c.horizon.AccountDetail(horizonclient.AccountRequest{
		AccountID: c.kp.Address(),
	})
	if err != nil {
		return errors.Wrap(err, "failed to get accountdetail")
	}
	tx, err := txnbuild.NewTransaction(txnbuild.TransactionParams{
		SourceAccount:        &sourceAccount,
		IncrementSequenceNum: true,
		BaseFee:              0,
		Operations:           []txnbuild.Operation{&payment, &feePayment},
		Preconditions: txnbuild.Preconditions{
			TimeBounds: txnbuild.NewInfiniteTimeout(),
		},
	})
	if err != nil {
		return errors.Wrap(err, "failed to create transaction")
	}
	return c.SignFundAndSubmitTransaction(tx)
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
