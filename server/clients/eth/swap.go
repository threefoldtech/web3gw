package goethclient

import (
	"context"
	"math/big"
	"time"

	coreEntities "github.com/daoleno/uniswap-sdk-core/entities"

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
	GoerliWethContract = "0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6"

	SlippageAmount = 3000
	GasLimit       = 210000

	GoerliChainId  = 5
	MainnetChainId = 1
)

var (
	TftGoerli = coreEntities.NewToken(GoerliChainId, common.HexToAddress(GoerliEthTftContractAddress), 7, "TFT", "TFT on Ethereum")
	Weth      = coreEntities.NewToken(GoerliChainId, common.HexToAddress(GoerliWethContract), 18, "WETH", "Wrapped Ether")

	SwapRouter = common.HexToAddress(helper.ContractV3SwapRouterV1)
)

func (c *Client) QuoteTftEth(ctx context.Context, amount string) (int64, error) {
	quoterContract, err := contract.NewUniswapv3Quoter(common.HexToAddress(helper.ContractV3Quoter), c.Eth)
	if err != nil {
		log.Err(err).Msg("failed to create quoter contract")
		return 0, err
	}

	weth := common.HexToAddress(GoerliWethContract)
	tft := common.HexToAddress(GoerliEthTftContractAddress)

	// 0.03% slippage
	fee := big.NewInt(SlippageAmount)

	amountIn := helper.FloatStringToBigInt(amount, 18)
	log.Debug().Str("amountIn", amountIn.String()).Msg("amountIn")
	sqrtPriceLimitX96 := big.NewInt(0)

	var out []interface{}
	rawCaller := &contract.Uniswapv3QuoterRaw{Contract: quoterContract}
	err = rawCaller.Call(nil, &out, "quoteExactInputSingle", weth, tft,
		fee, amountIn, sqrtPriceLimitX96)
	if err != nil {
		log.Err(err).Msg("failed to call quoteExactInputSingle")
		return 0, err
	}

	return out[0].(*big.Int).Int64(), nil
}

func (c *Client) SwapTftEth(ctx context.Context, amountIn string) (string, error) {
	// ether := coreEntities.EtherOnChain(1)

	pool, err := helper.ConstructV3Pool(c.Eth, Weth, TftGoerli, uint64(constants.FeeMedium))
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
	r, err := entities.NewRoute([]*entities.Pool{pool}, Weth, TftGoerli)
	if err != nil {
		log.Err(err).Msg("failed to create route")
		return "", err
	}

	swapValue := helper.FloatStringToBigInt(amountIn, 18)
	log.Debug().Str("amountIn", swapValue.String()).Msg("amountIn")

	trade, err := entities.FromRoute(r, coreEntities.FromRawAmount(Weth, swapValue), coreEntities.ExactInput)
	if err != nil {
		log.Err(err).Msg("failed to create trade")
		return "", err
	}

	log.Printf("input %v ouput %v\n", trade.Swaps[0].InputAmount.Quotient(), trade.Swaps[0].OutputAmount.Wrapped().Quotient())
	params, err := periphery.SwapCallParameters([]*entities.Trade{trade}, &periphery.SwapOptions{
		SlippageTolerance: slippageTolerance,
		Recipient:         c.Address,
		Deadline:          deadline,
	})
	if err != nil {
		log.Err(err).Msg("failed to get swap call parameters")
		return "", err
	}
	log.Printf("calldata = 0x%x", params.Value.String())

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

	// TODO: move this ?
	res, err := bind.WaitMined(ctx, c.Eth, tx)
	if err != nil {
		return "", err
	}

	log.Debug().Msgf("Swap tx mined: %s, block %d, gas: %d, status: %d", tx.Hash().Hex(), res.BlockNumber, res.GasUsed, res.Status)

	return s, nil
}
