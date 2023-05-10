package goethclient

import (
	"context"
	"math/big"
	"time"

	coreEntities "github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/pkg/errors"

	"github.com/daoleno/uniswapv3-sdk/constants"
	"github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/daoleno/uniswapv3-sdk/examples/contract"
	"github.com/daoleno/uniswapv3-sdk/examples/helper"
	"github.com/daoleno/uniswapv3-sdk/periphery"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
)

const (
	SlippageAmount = 3000
	GasLimit       = 210000

	GoerliChainId  = 5
	MainnetChainId = 1
)

var (
	GoerliWethContract  = common.HexToAddress("0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6")
	MainnetWethContract = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")

	GoerliTft  = coreEntities.NewToken(GoerliChainId, GoerliEthTftContractAddress, 7, "TFT", "TFT on Ethereum")
	GoerliWeth = coreEntities.NewToken(GoerliChainId, GoerliWethContract, 18, "WETH", "Wrapped Ether")

	MainnetTft  = coreEntities.NewToken(MainnetChainId, MainnetEthTftContractAddress, 7, "TFT", "TFT on Ethereum")
	MainnetWeth = coreEntities.NewToken(MainnetChainId, MainnetWethContract, 18, "WETH", "Wrapped Ether")

	SwapRouter = common.HexToAddress(helper.ContractV3SwapRouterV1)
)

func (c *Client) QuoteEthForTft(ctx context.Context, amount string) (int64, error) {
	tft, err := c.GetTftTokenContract()
	if err != nil {
		return 0, err
	}

	weth, err := c.GetWethTokenContract()
	if err != nil {
		return 0, err
	}

	return c.quoteTokens(ctx, amount, weth, tft)
}

func (c *Client) QuoteTftForEth(ctx context.Context, amount string) (int64, error) {
	tft, err := c.GetTftTokenContract()
	if err != nil {
		return 0, err
	}

	weth, err := c.GetWethTokenContract()
	if err != nil {
		return 0, err
	}

	return c.quoteTokens(ctx, amount, tft, weth)
}

func (c *Client) quoteTokens(ctx context.Context, input string, token0 *coreEntities.Token, token1 *coreEntities.Token) (int64, error) {
	quoterContract, err := contract.NewUniswapv3Quoter(common.HexToAddress(helper.ContractV3Quoter), c.Eth)
	if err != nil {
		log.Err(err).Msg("failed to create quoter contract")
		return 0, err
	}
	// 0.03% slippage
	fee := big.NewInt(SlippageAmount)

	amountIn := helper.FloatStringToBigInt(input, int(token0.Decimals()))
	sqrtPriceLimitX96 := big.NewInt(0)

	var out []interface{}
	rawCaller := &contract.Uniswapv3QuoterRaw{Contract: quoterContract}
	err = rawCaller.Call(nil, &out, "quoteExactInputSingle", token0.Address, token1.Address,
		fee, amountIn, sqrtPriceLimitX96)
	if err != nil {
		log.Err(err).Msg("failed to call quoteExactInputSingle")
		return 0, err
	}

	log.Debug().Msgf("Quote: input: %s, output: %s", input, out[0].(*big.Int).String())

	return out[0].(*big.Int).Int64(), nil
}

func (c *Client) SwapEthForTft(ctx context.Context, amountIn string) (string, error) {
	tft, err := c.GetTftTokenContract()
	if err != nil {
		return "", err
	}

	weth, err := c.GetWethTokenContract()
	if err != nil {
		return "", err
	}

	return c.makeSwap(ctx, amountIn, weth, tft)
}

func (c *Client) SwapTftForEth(ctx context.Context, amountIn string) (string, error) {
	tft, err := c.GetTftTokenContract()
	if err != nil {
		return "", err
	}

	weth, err := c.GetWethTokenContract()
	if err != nil {
		return "", err
	}

	return c.makeSwap(ctx, amountIn, tft, weth)
}

func (c *Client) makeSwap(ctx context.Context, input string, token0 *coreEntities.Token, token1 *coreEntities.Token) (string, error) {
	pool, err := helper.ConstructV3Pool(c.Eth, token0, token1, uint64(constants.FeeMedium))
	if err != nil {
		log.Err(err).Msg("failed to construct pool")
		return "", err
	}

	//0.01%
	slippageTolerance := coreEntities.NewPercent(big.NewInt(1), big.NewInt(SlippageAmount))
	//after 5 minutes
	d := time.Now().Add(time.Minute * time.Duration(15)).Unix()
	deadline := big.NewInt(d)

	// single trade input
	// single-hop exact input
	r, err := entities.NewRoute([]*entities.Pool{pool}, token0, token1)
	if err != nil {
		log.Err(err).Msg("failed to create route")
		return "", err
	}

	swapValue := helper.FloatStringToBigInt(input, int(token0.Decimals()))

	trade, err := entities.FromRoute(r, coreEntities.FromRawAmount(token0, swapValue), coreEntities.ExactInput)
	if err != nil {
		log.Err(err).Msg("failed to create trade")
		return "", err
	}

	log.Printf("Swap: input %v ouput %v", trade.Swaps[0].InputAmount.Quotient(), trade.Swaps[0].OutputAmount.Wrapped().Quotient())
	params, err := periphery.SwapCallParameters([]*entities.Trade{trade}, &periphery.SwapOptions{
		SlippageTolerance: slippageTolerance,
		Recipient:         c.Address,
		Deadline:          deadline,
	})
	if err != nil {
		log.Err(err).Msg("failed to get swap call parameters")
		return "", err
	}

	gasPrice, err := c.Eth.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	log.Debug().Msgf("gasLimit %d gasprice %d", GasLimit, gasPrice.Uint64())
	nounc, err := c.Eth.NonceAt(context.Background(), c.Address, nil)
	if err != nil {
		return "", err
	}

	tx := types.NewTransaction(nounc, SwapRouter, swapValue,
		GasLimit, gasPrice, params.Calldata)

	s, err := c.sendTransaction(tx)
	if err != nil {
		return "", err
	}

	res, err := bind.WaitMined(ctx, c.Eth, tx)
	if err != nil {
		return "", err
	}

	log.Debug().Msgf("Swap tx mined: %s, block %d, gas: %d, status: %d", tx.Hash().Hex(), res.BlockNumber, res.GasUsed, res.Status)

	return s, nil
}

func (c *Client) GetTftTokenContract() (*coreEntities.Token, error) {
	chainID, err := c.Eth.NetworkID(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get chainID")
	}

	if chainID.Cmp(big.NewInt(EthMainnetId)) == 0 {
		return MainnetTft, nil
	} else if chainID.Cmp(big.NewInt(EthGoerliId)) == 0 {
		return GoerliTft, nil
	} else {
		return nil, errors.New("unsupported chainID")
	}
}

func (c *Client) GetWethTokenContract() (*coreEntities.Token, error) {
	chainID, err := c.Eth.NetworkID(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get chainID")
	}

	if chainID.Cmp(big.NewInt(EthMainnetId)) == 0 {
		return MainnetWeth, nil
	} else if chainID.Cmp(big.NewInt(EthGoerliId)) == 0 {
		return GoerliWeth, nil
	} else {
		return nil, errors.New("unsupported chainID")
	}
}
