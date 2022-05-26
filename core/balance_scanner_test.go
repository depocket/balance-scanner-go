package core

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBalanceScanner(t *testing.T) {
	client, _ := ethclient.Dial("https://bsc-dataseed1.ninicoin.io/")
	bep20, _ := NewBalanceScanner(BinanceChain, client)
	balances, _ := bep20.GetBalances(
		&bind.CallOpts{},
		common.HexToAddress("0x7d99eda556388Ad7743A1B658b9C4FC67D7A9d74"), []common.Address{
			common.HexToAddress("0x7d99eda556388Ad7743A1B658b9C4FC67D7A9d74"),
		})
	assert.Equal(t, balances[0].String(), "34000000000000000000")
}
