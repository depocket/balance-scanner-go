// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package core

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var BalanceScannerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"contractIBEP20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getBalances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

var BalanceScannerABI = BalanceScannerMetaData.ABI

type BalanceScanner struct {
	contract *bind.BoundContract
}

func NewBalanceScanner(address string, backend bind.ContractBackend) (*BalanceScanner, error) {
	contract, err := bindBalanceScanner(common.HexToAddress(address), backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalanceScanner{contract: contract}, nil
}

func bindBalanceScanner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalanceScannerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_BalanceScanner *BalanceScanner) GetBalances(opts *bind.CallOpts, account common.Address, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _BalanceScanner.contract.Call(opts, &out, "getBalances", account, tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err
}


func (_BalanceScanner *BalanceScanner) TokenMapBalances(address common.Address, tokens []common.Address) (map[string]*big.Int, error) {
	balances := make(map[string]*big.Int)
	res, err := _BalanceScanner.GetBalances(&bind.CallOpts{}, address, tokens)
	if err != nil {
		return balances, nil
	}
	for i, token := range tokens {
		balances[strings.ToLower(token.String())] = res[i]
	}
	return balances, nil
}