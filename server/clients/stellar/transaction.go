package stellargoclient

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/stellar/go/clients/horizonclient"
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

func (c *Client) AwaitTransactionWithMemo(ctx context.Context, account string, memo string, timeout int) error {
	memo = strings.TrimPrefix(memo, "0x")
	for i := 0; i < int(timeout); i++ {
		select {
		case <-time.After(1 * time.Second):
			transactions, err := c.Transactions(account, 10, false, "", horizonclient.OrderDesc)
			if err != nil {
				return err
			}
			for _, tx := range transactions {
				decodedMemo, err := base64.StdEncoding.DecodeString(tx.Memo)
				if err == nil {
					hexDecodedMemo := hex.EncodeToString(decodedMemo)
					if hexDecodedMemo == memo {
						return nil
					}
				}
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return errors.New("transaction not found")
}

func (c *Client) AwaitTransactionWithMemoOnEthBridge(ctx context.Context, memo string, timeout int) error {
	bridgeAddress, err := c.GetEthBridgeAddress()
	if err != nil {
		return err
	}
	return c.AwaitTransactionWithMemo(ctx, bridgeAddress, memo, timeout)
}

func (c *Client) AwaitForTransactionWithMemoOnTfchainBridge(ctx context.Context, memo string, timeout int) error {
	bridgeAddress, err := c.GetTfchainBridgeAddress()
	if err != nil {
		return err
	}
	return c.AwaitTransactionWithMemo(ctx, bridgeAddress, memo, timeout)
}

func (c *Client) Transactions(account string, limit uint, includeFailed bool, cursor string, order horizonclient.Order) ([]horizon.Transaction, error) {
	transactionRequest := horizonclient.TransactionRequest{
		ForAccount:    account,
		Limit:         limit,
		Order:         order,
		IncludeFailed: includeFailed,
		Cursor:        cursor,
	}

	txs, err := c.horizon.Transactions(transactionRequest)
	if err != nil {
		return []horizon.Transaction{}, err
	}
	return txs.Embedded.Records, nil
}

// Allows you to fund a transaction with tft instead of lumen by using the transaction funding service: https://github.com/threefoldfoundation/tft-stellar/tree/master/ThreeBotPackages/transactionfunding_service
func (c *Client) FundTransactionUsingTft(destination string, amount string) error {
	feeAccount, fee, err := c.getTFTTransactionFundingCondition()
	if err != nil {
		return errors.Wrap(err, "failed to get tft transaction funding condition")
	}
	payment := txnbuild.Payment{
		Destination: destination,
		Amount:      amount,
		Asset:       c.GetTftAsset(),
	}
	feePayment := txnbuild.Payment{
		Destination: feeAccount,
		Amount:      fee,
		Asset:       c.GetTftAsset(),
	}

	sourceAccount, err := c.AccountData(c.kp.Address())
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
