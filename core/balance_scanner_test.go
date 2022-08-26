package core

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestBalanceScanner(t *testing.T) {
	client, _ := ethclient.Dial("https://bsc-dataseed1.ninicoin.io/")
	bep20, _ := NewBalanceScanner(BinanceChain, client, nil)
	balances, _ := bep20.GetBalances(
		common.HexToAddress("0x7d99eda556388Ad7743A1B658b9C4FC67D7A9d74"), []common.Address{
			common.HexToAddress("0x7d99eda556388Ad7743A1B658b9C4FC67D7A9d74"),
		})
	assert.Equal(t, balances[0].String(), "34000000000000000000")
}

func TestBalanceScanner_TokenMapBalances(t *testing.T) {
	client, _ := ethclient.Dial("https://bsc-dataseed1.ninicoin.io/")
	bep20, _ := NewBalanceScanner(BinanceChain, client, nil)
	balances, _ := bep20.TokenMapBalances(
		common.HexToAddress("0x7d99eda556388Ad7743A1B658b9C4FC67D7A9d74"), []common.Address{
			common.HexToAddress("0x7d99eda556388Ad7743A1B658b9C4FC67D7A9d74"),
		})
	assert.Equal(t, balances[strings.ToLower("0x7d99eda556388Ad7743A1B658b9C4FC67D7A9d74")].String(), "34000000000000000000")
}

func TestBalanceScanner_BatchCall(t *testing.T) {
	client, _ := ethclient.Dial("https://bsc-dataseed.binance.org/")
	r, _ := rpc.Dial("https://bsc-dataseed.binance.org/")
	bep20, _ := NewBalanceScanner(BinanceChain, client, r)
	resultCount := 0
	ret, _ := bep20.BatchCallBalances(common.HexToAddress("0x7d99eda556388Ad7743A1B658b9C4FC67D7A9d74"), [][]common.Address{
		{
			common.HexToAddress("0x7d99eda556388Ad7743A1B658b9C4FC67D7A9d74"),
			common.HexToAddress("0x7d99eda556388Ad7743A1B658b9C4FC67D7A9d74"),
		},
		{
			common.HexToAddress("0x7d99eda556388Ad7743A1B658b9C4FC67D7A9d74"),
			common.HexToAddress("0x7d99eda556388Ad7743A1B658b9C4FC67D7A9d74"),
			common.HexToAddress("0x7d99eda556388Ad7743A1B658b9C4FC67D7A9d74"),
		},
		{
			common.HexToAddress("0x7d99eda556388Ad7743A1B658b9C4FC67D7A9d74"),
			common.HexToAddress("0x7d99eda556388Ad7743A1B658b9C4FC67D7A9d74"),
			common.HexToAddress("0x7d99eda556388Ad7743A1B658b9C4FC67D7A9d74"),
			common.HexToAddress("0x7d99eda556388Ad7743A1B658b9C4FC67D7A9d74"),
			common.HexToAddress("0x7d99eda556388Ad7743A1B658b9C4FC67D7A9d74"),
		},
	})
	for j := 0; j < len(ret); j++ {
		resultCount++
		assert.Equal(t, ret[j].String(), "34000000000000000000")
	}
	fmt.Println("Total result", resultCount)
}
