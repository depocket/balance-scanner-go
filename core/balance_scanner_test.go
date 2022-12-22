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

func TestBalanceScanner_MultiBatchCall(t *testing.T) {
	client, _ := ethclient.Dial("https://bsc-dataseed.binance.org/")
	r, _ := rpc.Dial("https://bsc-dataseed.binance.org/")
	bep20, _ := NewBalanceScanner("0xcA11bde05977b3631167028862bE2a173976CA11", client, r)
	resultCount := 0
	ret, err := bep20.BatchCallBalancesWithNative(common.HexToAddress("0xca0C80122afA57c38BcAa14fC77E056b94288469"), [][]common.Address{
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
			common.HexToAddress("0x0000000000000000000000000000000000000000"),
		},
	}, "0x0000000000000000000000000000000000000000")
	if err != nil {
		fmt.Println(err)
	}
	for j := 0; j < len(ret); j++ {
		if j == len(ret)-1 {
			assert.Equal(t, ret[j].String(), "207924141787718960")
		} else {
			assert.Equal(t, ret[j].String(), "1081707507519149458")
		}
		resultCount++
	}
	fmt.Println("Total result", resultCount)
}

func TestBalanceScanner_MultiBatchCallERC20Only(t *testing.T) {
	client, _ := ethclient.Dial("https://bsc-dataseed.binance.org/")
	r, _ := rpc.Dial("https://bsc-dataseed.binance.org/")
	bep20, _ := NewBalanceScanner("0xcA11bde05977b3631167028862bE2a173976CA11", client, r)
	resultCount := 0
	ret, err := bep20.BatchCallBalancesWithNative(common.HexToAddress("0xca0C80122afA57c38BcAa14fC77E056b94288469"), [][]common.Address{
		{
			common.HexToAddress("0x7d99eda556388Ad7743A1B658b9C4FC67D7A9d74"),
			common.HexToAddress("0x7d99eda556388Ad7743A1B658b9C4FC67D7A9d74"),
			common.HexToAddress("0x7d99eda556388Ad7743A1B658b9C4FC67D7A9d74"),
			common.HexToAddress("0x7d99eda556388Ad7743A1B658b9C4FC67D7A9d74"),
			common.HexToAddress("0x7d99eda556388Ad7743A1B658b9C4FC67D7A9d74"),
		},
	}, "")
	if err != nil {
		fmt.Println(err)
	}
	for j := 0; j < len(ret); j++ {
		assert.Equal(t, ret[j].String(), "1081707507519149458")
		resultCount++
	}
	fmt.Println("Total result", resultCount)
}

func TestBalanceScanner_MultiBatchCallNativeOnly(t *testing.T) {
	client, _ := ethclient.Dial("https://bsc-dataseed.binance.org/")
	r, _ := rpc.Dial("https://bsc-dataseed.binance.org/")
	bep20, _ := NewBalanceScanner("0xcA11bde05977b3631167028862bE2a173976CA11", client, r)
	resultCount := 0
	ret, err := bep20.BatchCallBalancesWithNative(common.HexToAddress("0xca0C80122afA57c38BcAa14fC77E056b94288469"), [][]common.Address{
		{
			common.HexToAddress("0x0000000000000000000000000000000000000000"),
		},
	}, "0x0000000000000000000000000000000000000000")
	if err != nil {
		fmt.Println(err)
	}
	for j := 0; j < len(ret); j++ {
		assert.Equal(t, ret[j].String(), "207924141787718960")
		resultCount++
	}
	fmt.Println("Total result", resultCount)
}
