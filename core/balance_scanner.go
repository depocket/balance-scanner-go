// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package core

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var BalanceScannerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"contractIBEP20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getBalances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

type CallArg struct {
	From     common.Address   `json:"from"`
	To       common.Address  `json:"to"`
	Gas      string         `json:"gas"`
	GasPrice string         `json:"gasPrice"`
	Value    string          `json:"value"`
	Data     string           `json:"data"`
	Nonce    string          `json:"nonce"`
}

var BalanceScannerABI = BalanceScannerMetaData.ABI

type BalanceScanner struct {
	Client          *ethclient.Client
	RpcClient          *rpc.Client
	ContractAddress common.Address
	_abi             abi.ABI
}

func NewBalanceScanner(address string, client *ethclient.Client, rpcClient *rpc.Client) (*BalanceScanner, error) {
	_abi, _ := abi.JSON(strings.NewReader(BalanceScannerABI))
	return &BalanceScanner{
		_abi: _abi,
		Client: client,
		RpcClient: rpcClient,
		ContractAddress: common.HexToAddress(address),
	}, nil
}

func (_BalanceScanner *BalanceScanner) GetBalances(account common.Address, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	callData, err := _BalanceScanner._abi.Pack("getBalances", account, tokens)
	if err != nil {
		return *new([]*big.Int), err
	}
	resp, err := _BalanceScanner.Client.CallContract(context.Background(), ethereum.CallMsg{To: &_BalanceScanner.ContractAddress, Data: callData}, nil)
	if err != nil {
		return *new([]*big.Int), err
	}
	out, err = _BalanceScanner._abi.Unpack("getBalances", resp)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, err
}

func (_BalanceScanner *BalanceScanner) TokenMapBalances(address common.Address, tokens []common.Address) (map[string]*big.Int, error) {
	balances := make(map[string]*big.Int)
	res, err := _BalanceScanner.GetBalances(address, tokens)
	if err != nil {
		return balances, nil
	}
	for i, token := range tokens {
		balances[strings.ToLower(token.String())] = res[i]
	}
	return balances, nil
}

func (_BalanceScanner *BalanceScanner) GetBalancesCallData(account common.Address, tokens []common.Address) ([]byte, error) {
	return _BalanceScanner._abi.Pack("getBalances", account, tokens)
}

func (_BalanceScanner *BalanceScanner) DecodeCallResponse(data hexutil.Bytes) ([]interface{}, error) {
	return _BalanceScanner._abi.Unpack("getBalances", data)
}

func (_BalanceScanner *BalanceScanner) BatchCallBalances(account common.Address, tokenChunks [][]common.Address) ([]*big.Int, error) {
	var res []*big.Int
	var calls []rpc.BatchElem
	var callSize = len(tokenChunks)
	var callResults = make([]hexutil.Bytes, callSize)
	for index, chunk := range tokenChunks {
		callData, err := _BalanceScanner.GetBalancesCallData(account, chunk)
		if err != nil {
			return []*big.Int{}, err
		}
		elem := rpc.BatchElem {
			Method: "eth_call",
			Args: toCallArgs(ethereum.CallMsg{To: &_BalanceScanner.ContractAddress, Data: callData}),
			Result: &callResults[index],
		}
		calls = append(calls, elem)
	}
	err := _BalanceScanner.RpcClient.BatchCall(calls)
	if err != nil {
		return []*big.Int{}, err
	}
	for _, result := range callResults {
		r, err := _BalanceScanner.DecodeCallResponse(result)
		if err != nil {
			return []*big.Int{}, err
		}
		res = append(res, r[0].([]*big.Int)...)
	}
	return res, nil
}

func (_BalanceScanner *BalanceScanner) BatchCallBalancesWithContext(ctx context.Context, account common.Address, tokenChunks [][]common.Address) ([]*big.Int, error) {
	var res []*big.Int
	var calls []rpc.BatchElem
	var callSize = len(tokenChunks)
	var callResults = make([]hexutil.Bytes, callSize)
	for index, chunk := range tokenChunks {
		callData, err := _BalanceScanner.GetBalancesCallData(account, chunk)
		if err != nil {
			return []*big.Int{}, err
		}
		elem := rpc.BatchElem {
			Method: "eth_call",
			Args: toCallArgs(ethereum.CallMsg{To: &_BalanceScanner.ContractAddress, Data: callData}),
			Result: &callResults[index],
		}
		calls = append(calls, elem)
	}
	err := _BalanceScanner.RpcClient.BatchCallContext(ctx, calls)
	if err != nil {
		return []*big.Int{}, err
	}
	for _, result := range callResults {
		r, err := _BalanceScanner.DecodeCallResponse(result)
		if err != nil {
			return []*big.Int{}, err
		}
		res = append(res, r[0].([]*big.Int)...)
	}
	return res, nil
}

func toCallArgs(msg ethereum.CallMsg) []interface{} {
	arg := []interface{}{CallArg{
		To:       *msg.To,
		Gas: hexutil.EncodeUint64(1000000000), // 1000000000 was used to avoid the need to config gas limit on view call
		GasPrice: hexutil.EncodeUint64(0),
		Value: hexutil.EncodeUint64(0),
		Nonce: hexutil.EncodeUint64(0),
		Data: fmt.Sprintf("0x%s", common.Bytes2Hex(msg.Data)),
	}, toBlockNumArg(nil)}
	return arg
}

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	pending := big.NewInt(-1)
	if number.Cmp(pending) == 0 {
		return "pending"
	}
	return hexutil.EncodeBig(number)
}