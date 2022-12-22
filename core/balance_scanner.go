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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"contractIBEP20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getBalances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structDePocketCore.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structDePocketCore.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"blockAndAggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structDePocketCore.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structDePocketCore.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"tryAggregate\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structDePocketCore.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structDePocketCore.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"tryBlockAndAggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structDePocketCore.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

type MultiCall struct {
	Target   common.Address
	CallData []byte
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


/*
	GetBalances query the ERC20 token balances in batch
 */
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

/*
	TokenMapBalances query the ERC20 token balances in batch and return the result in mapping by token_address:balance
*/
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

/*
	BatchCallBalancesWithNative query the ERC20 token balances in batch of RPC and multi-call and return the result in mapping by token_address:balance
*/
func (_BalanceScanner *BalanceScanner) BatchCallBalancesWithNative(account common.Address, tokenChunks [][]common.Address, nativeAddress string) ([]*big.Int, error) {
	var calls []rpc.BatchElem
	var callResults = make([]hexutil.Bytes, len(tokenChunks))
	var nativeResult = new(hexutil.Big)
	var nativeIndex = -1 // Determinate the index of native token
	var skipResult = -1 // Determinate the balanceOf result that need to be skipped when decode the result
	var count = 0

	for index, chunk := range tokenChunks {
		var multiBalanceCalls []MultiCall
		if len(chunk) == 1 && strings.ToLower(chunk[0].String()) == nativeAddress {
			nativeIndex = count
			skipResult = index
			count++
			continue
		}
		for _, tk := range chunk {
			data, err := _BalanceScanner.BalanceOfCallData(account)
			if err != nil {
				return nil, err
			}
			if strings.ToLower(tk.String()) != nativeAddress {
				multiBalanceCalls = append(multiBalanceCalls, MultiCall{
					Target: tk,
					CallData: data,
				})
			} else {
				nativeIndex = count
			}
			count++
		}
		if len(multiBalanceCalls) > 0 {
			callData, err := _BalanceScanner.GetMulticallBalanceCallData(multiBalanceCalls)
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
	}

	if nativeIndex != -1 {
		nativeElem := rpc.BatchElem {
			Method: "eth_getBalance",
			Args: []interface{}{account, toBlockNumArg(nil)},
			Result: &nativeResult,
		}
		calls = append(calls, nativeElem)
	}

	err := _BalanceScanner.RpcClient.BatchCall(calls)

	if err != nil {
		return []*big.Int{}, err
	}

	var res = make([]*big.Int, count)
	resCount := 0
	for idx, result := range callResults {
		if idx == skipResult {
			continue
		}
		if resCount == nativeIndex {
			resCount ++
		}
		r, err := _BalanceScanner.DecodeMulticallBalanceResponse(result)
		if err != nil {
			return []*big.Int{}, err
		}
		for _, response := range r[0].([]struct {
			Success    bool   `json:"success"`
			ReturnData []byte `json:"returnData"`
		}) {
			if response.Success {
				balance, err := _BalanceScanner.DeCodeBalanceOf(response.ReturnData)
				if err != nil {
					return []*big.Int{}, err
				}
				res[resCount] = balance[0].(*big.Int)
			} else {
				res[resCount] = new(big.Int)
			}
			resCount++
		}
	}
	if nativeIndex != -1 {
		res[nativeIndex] = nativeResult.ToInt()
	}
	return res, nil
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

func (_BalanceScanner *BalanceScanner) GetBalancesCallData(account common.Address, tokens []common.Address) ([]byte, error) {
	return _BalanceScanner._abi.Pack("getBalances", account, tokens)
}

func (_BalanceScanner *BalanceScanner) BalanceOfCallData(account common.Address) ([]byte, error) {
	return _BalanceScanner._abi.Pack("balanceOf", account)
}

func (_BalanceScanner *BalanceScanner) DeCodeBalanceOf(data hexutil.Bytes) ([]interface{}, error) {
	return _BalanceScanner._abi.Unpack("balanceOf", data)
}

func (_BalanceScanner *BalanceScanner) DecodeCallResponse(data hexutil.Bytes) ([]interface{}, error) {
	return _BalanceScanner._abi.Unpack("getBalances", data)
}

func (_BalanceScanner *BalanceScanner) DecodeMulticallBalanceResponse(data hexutil.Bytes) ([]interface{}, error) {
	return _BalanceScanner._abi.Unpack("tryAggregate", data)
}

func (_BalanceScanner *BalanceScanner) GetMulticallBalanceCallData(calls []MultiCall) ([]byte, error) {
	return _BalanceScanner._abi.Pack("tryAggregate", false, calls)
}