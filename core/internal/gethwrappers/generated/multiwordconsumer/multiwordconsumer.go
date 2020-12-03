// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package multiwordconsumer

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MultiwordConsumerABI is the input ABI used to generate the binding from.
const MultiwordConsumerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_link\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_specId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"price\",\"type\":\"bytes\"}],\"name\":\"RequestFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"usd\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"eur\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"jpy\",\"type\":\"bytes32\"}],\"name\":\"RequestMultipleFulfilled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"}],\"name\":\"addExternalRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_payment\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"_callbackFunctionId\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"_expiration\",\"type\":\"uint256\"}],\"name\":\"cancelRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentPrice\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eur\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_price\",\"type\":\"bytes\"}],\"name\":\"fulfillBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_usd\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_eur\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_jpy\",\"type\":\"bytes32\"}],\"name\":\"fulfillMultipleParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"jpy\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_currency\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_payment\",\"type\":\"uint256\"}],\"name\":\"requestEthereumPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_currency\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_payment\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_callback\",\"type\":\"address\"}],\"name\":\"requestEthereumPriceByCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_currency\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_payment\",\"type\":\"uint256\"}],\"name\":\"requestMultipleParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_specId\",\"type\":\"bytes32\"}],\"name\":\"setSpecID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usd\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MultiwordConsumerBin is the compiled bytecode used for deploying new contracts.
var MultiwordConsumerBin = "0x6080604052600160045534801561001557600080fd5b506040516113ec3803806113ec8339818101604052606081101561003857600080fd5b508051602082015160409092015190919061005283610066565b61005b82610088565b600655506100aa9050565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b600380546001600160a01b0319166001600160a01b0392909216919091179055565b611333806100b96000396000f3fe608060405234801561001057600080fd5b50600436106100df5760003560e01c80639d1b464a1161008c578063d63a6ccd11610066578063d63a6ccd14610454578063e89855ba1461045c578063e8d5359d14610504578063faa367611461053d576100df565b80639d1b464a146102fb578063a856ff6b14610378578063c2fb8523146103a7576100df565b806374961d4d116100bd57806374961d4d1461018a57806383db5cbc1461024b5780638dc654a2146102f3576100df565b8063501fdd5d146100e45780635591a608146101035780637439ae5914610170575b600080fd5b610101600480360360208110156100fa57600080fd5b5035610545565b005b610101600480360360a081101561011957600080fd5b5073ffffffffffffffffffffffffffffffffffffffff813516906020810135906040810135907fffffffff00000000000000000000000000000000000000000000000000000000606082013516906080013561054a565b610178610611565b60408051918252519081900360200190f35b610101600480360360608110156101a057600080fd5b8101906020810181356401000000008111156101bb57600080fd5b8201836020820111156101cd57600080fd5b803590602001918460018302840111640100000000831117156101ef57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550508235935050506020013573ffffffffffffffffffffffffffffffffffffffff16610617565b6101016004803603604081101561026157600080fd5b81019060208101813564010000000081111561027c57600080fd5b82018360208201111561028e57600080fd5b803590602001918460018302840111640100000000831117156102b057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610660915050565b61010161066f565b610303610839565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561033d578181015183820152602001610325565b50505050905090810190601f16801561036a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101016004803603608081101561038e57600080fd5b50803590602081013590604081013590606001356108e5565b610101600480360360408110156103bd57600080fd5b813591908101906040810160208201356401000000008111156103df57600080fd5b8201836020820111156103f157600080fd5b8035906020019184600183028401116401000000008311171561041357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610a08945050505050565b610178610bba565b6101016004803603604081101561047257600080fd5b81019060208101813564010000000081111561048d57600080fd5b82018360208201111561049f57600080fd5b803590602001918460018302840111640100000000831117156104c157600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610bc0915050565b6101016004803603604081101561051a57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060200135610c02565b610178610c0c565b600655565b604080517f6ee4d55300000000000000000000000000000000000000000000000000000000815260048101869052602481018590527fffffffff0000000000000000000000000000000000000000000000000000000084166044820152606481018390529051869173ffffffffffffffffffffffffffffffffffffffff831691636ee4d5539160848082019260009290919082900301818387803b1580156105f157600080fd5b505af1158015610605573d6000803e3d6000fd5b50505050505050505050565b60095481565b61061f6111d0565b60065461064d90837fc2fb852300000000000000000000000000000000000000000000000000000000610c12565b90506106598184610c37565b5050505050565b61066b828230610617565b5050565b6000610679610c65565b90508073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb338373ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b1580156106ff57600080fd5b505afa158015610713573d6000803e3d6000fd5b505050506040513d602081101561072957600080fd5b5051604080517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b16815273ffffffffffffffffffffffffffffffffffffffff909316600484015260248301919091525160448083019260209291908290030181600087803b15801561079f57600080fd5b505af11580156107b3573d6000803e3d6000fd5b505050506040513d60208110156107c957600080fd5b505161083657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f556e61626c6520746f207472616e736665720000000000000000000000000000604482015290519081900360640190fd5b50565b6007805460408051602060026001851615610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190941693909304601f810184900484028201840190925281815292918301828280156108dd5780601f106108b2576101008083540402835291602001916108dd565b820191906000526020600020905b8154815290600101906020018083116108c057829003601f168201915b505050505081565b600084815260056020526040902054849073ffffffffffffffffffffffffffffffffffffffff163314610963576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260288152602001806112d66028913960400191505060405180910390fd5b60008181526005602052604080822080547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555182917f7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a91a28284867f0ec0c13e44aa04198947078cb990660252870dd3363f4c4bb3cc780f808dabbe856040518082815260200191505060405180910390a450600892909255600955600a5550565b600082815260056020526040902054829073ffffffffffffffffffffffffffffffffffffffff163314610a86576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260288152602001806112d66028913960400191505060405180910390fd5b60008181526005602052604080822080547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555182917f7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a91a2816040518082805190602001908083835b60208310610b2f57805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe09092019160209182019101610af2565b5181516020939093036101000a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01801990911692169190911790526040519201829003822093508692507f1a111c5dcf9a71088bd5e1797fdfaf399fec2afbb24aca247e4e3e9f4b61df919160009150a38151610bb4906007906020850190611205565b50505050565b60085481565b610bc86111d0565b600654610bf690307fa856ff6b00000000000000000000000000000000000000000000000000000000610c12565b9050610bb48183610c37565b61066b8282610c81565b600a5481565b610c1a6111d0565b610c226111d0565b610c2e81868686610d68565b95945050505050565b600354600090610c5e9073ffffffffffffffffffffffffffffffffffffffff168484610dca565b9392505050565b60025473ffffffffffffffffffffffffffffffffffffffff1690565b600081815260056020526040902054819073ffffffffffffffffffffffffffffffffffffffff1615610d1457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f5265717565737420697320616c72656164792070656e64696e67000000000000604482015290519081900360640190fd5b50600090815260056020526040902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b610d706111d0565b610d808560800151610100610ff8565b505091835273ffffffffffffffffffffffffffffffffffffffff1660208301527fffffffff0000000000000000000000000000000000000000000000000000000016604082015290565b6004546040805130606090811b60208084019190915260348084018690528451808503909101815260549093018452825192810192909220908601939093526000838152600590915281812080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8816179055905182917fb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af991a26002805473ffffffffffffffffffffffffffffffffffffffff1690634000aea09086908590610eab908890611032565b6040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610f19578181015183820152602001610f01565b50505050905090810190601f168015610f465780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b158015610f6757600080fd5b505af1158015610f7b573d6000803e3d6000fd5b505050506040513d6020811015610f9157600080fd5b5051610fe8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260238152602001806112b36023913960400191505060405180910390fd5b6004805460010190559392505050565b611000611283565b60208206156110155760208206602003820191505b506020828101829052604080518085526000815290920101905290565b6060634042994660e01b6000808560000151866020015187604001518860600151888a6080015160000151604051602401808973ffffffffffffffffffffffffffffffffffffffff1681526020018881526020018781526020018673ffffffffffffffffffffffffffffffffffffffff168152602001857bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156111125781810151838201526020016110fa565b50505050905090810190601f16801561113f5780820380516001836020036101000a031916815260200191505b50604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909d169c909c17909b525098995050505050505050505092915050565b6040805160a081018252600080825260208201819052918101829052606081019190915260808101611200611283565b905290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061124657805160ff1916838001178555611273565b82800160010185558215611273579182015b82811115611273578251825591602001919060010190611258565b5061127f92915061129d565b5090565b604051806040016040528060608152602001600081525090565b5b8082111561127f576000815560010161129e56fe756e61626c6520746f207472616e73666572416e6443616c6c20746f206f7261636c65536f75726365206d75737420626520746865206f7261636c65206f66207468652072657175657374a264697066735822beefbeefbeefbeefbeefbeefbeefbeefbeefbeefbeefbeefbeefbeefbeefbeefbeef64736f6c6343decafe0033"

// DeployMultiwordConsumer deploys a new Ethereum contract, binding an instance of MultiwordConsumer to it.
func DeployMultiwordConsumer(auth *bind.TransactOpts, backend bind.ContractBackend, _link common.Address, _oracle common.Address, _specId [32]byte) (common.Address, *types.Transaction, *MultiwordConsumer, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiwordConsumerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MultiwordConsumerBin), backend, _link, _oracle, _specId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MultiwordConsumer{MultiwordConsumerCaller: MultiwordConsumerCaller{contract: contract}, MultiwordConsumerTransactor: MultiwordConsumerTransactor{contract: contract}, MultiwordConsumerFilterer: MultiwordConsumerFilterer{contract: contract}}, nil
}

// MultiwordConsumer is an auto generated Go binding around an Ethereum contract.
type MultiwordConsumer struct {
	MultiwordConsumerCaller     // Read-only binding to the contract
	MultiwordConsumerTransactor // Write-only binding to the contract
	MultiwordConsumerFilterer   // Log filterer for contract events
}

// MultiwordConsumerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiwordConsumerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiwordConsumerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiwordConsumerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiwordConsumerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiwordConsumerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiwordConsumerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiwordConsumerSession struct {
	Contract     *MultiwordConsumer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MultiwordConsumerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiwordConsumerCallerSession struct {
	Contract *MultiwordConsumerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// MultiwordConsumerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiwordConsumerTransactorSession struct {
	Contract     *MultiwordConsumerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// MultiwordConsumerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiwordConsumerRaw struct {
	Contract *MultiwordConsumer // Generic contract binding to access the raw methods on
}

// MultiwordConsumerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiwordConsumerCallerRaw struct {
	Contract *MultiwordConsumerCaller // Generic read-only contract binding to access the raw methods on
}

// MultiwordConsumerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiwordConsumerTransactorRaw struct {
	Contract *MultiwordConsumerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiwordConsumer creates a new instance of MultiwordConsumer, bound to a specific deployed contract.
func NewMultiwordConsumer(address common.Address, backend bind.ContractBackend) (*MultiwordConsumer, error) {
	contract, err := bindMultiwordConsumer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiwordConsumer{MultiwordConsumerCaller: MultiwordConsumerCaller{contract: contract}, MultiwordConsumerTransactor: MultiwordConsumerTransactor{contract: contract}, MultiwordConsumerFilterer: MultiwordConsumerFilterer{contract: contract}}, nil
}

// NewMultiwordConsumerCaller creates a new read-only instance of MultiwordConsumer, bound to a specific deployed contract.
func NewMultiwordConsumerCaller(address common.Address, caller bind.ContractCaller) (*MultiwordConsumerCaller, error) {
	contract, err := bindMultiwordConsumer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiwordConsumerCaller{contract: contract}, nil
}

// NewMultiwordConsumerTransactor creates a new write-only instance of MultiwordConsumer, bound to a specific deployed contract.
func NewMultiwordConsumerTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiwordConsumerTransactor, error) {
	contract, err := bindMultiwordConsumer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiwordConsumerTransactor{contract: contract}, nil
}

// NewMultiwordConsumerFilterer creates a new log filterer instance of MultiwordConsumer, bound to a specific deployed contract.
func NewMultiwordConsumerFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiwordConsumerFilterer, error) {
	contract, err := bindMultiwordConsumer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiwordConsumerFilterer{contract: contract}, nil
}

// bindMultiwordConsumer binds a generic wrapper to an already deployed contract.
func bindMultiwordConsumer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiwordConsumerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiwordConsumer *MultiwordConsumerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MultiwordConsumer.Contract.MultiwordConsumerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiwordConsumer *MultiwordConsumerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiwordConsumer.Contract.MultiwordConsumerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiwordConsumer *MultiwordConsumerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiwordConsumer.Contract.MultiwordConsumerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiwordConsumer *MultiwordConsumerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MultiwordConsumer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiwordConsumer *MultiwordConsumerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiwordConsumer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiwordConsumer *MultiwordConsumerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiwordConsumer.Contract.contract.Transact(opts, method, params...)
}

// CurrentPrice is a free data retrieval call binding the contract method 0x9d1b464a.
//
// Solidity: function currentPrice() view returns(bytes)
func (_MultiwordConsumer *MultiwordConsumerCaller) CurrentPrice(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _MultiwordConsumer.contract.Call(opts, out, "currentPrice")
	return *ret0, err
}

// CurrentPrice is a free data retrieval call binding the contract method 0x9d1b464a.
//
// Solidity: function currentPrice() view returns(bytes)
func (_MultiwordConsumer *MultiwordConsumerSession) CurrentPrice() ([]byte, error) {
	return _MultiwordConsumer.Contract.CurrentPrice(&_MultiwordConsumer.CallOpts)
}

// CurrentPrice is a free data retrieval call binding the contract method 0x9d1b464a.
//
// Solidity: function currentPrice() view returns(bytes)
func (_MultiwordConsumer *MultiwordConsumerCallerSession) CurrentPrice() ([]byte, error) {
	return _MultiwordConsumer.Contract.CurrentPrice(&_MultiwordConsumer.CallOpts)
}

// Eur is a free data retrieval call binding the contract method 0x7439ae59.
//
// Solidity: function eur() view returns(bytes32)
func (_MultiwordConsumer *MultiwordConsumerCaller) Eur(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MultiwordConsumer.contract.Call(opts, out, "eur")
	return *ret0, err
}

// Eur is a free data retrieval call binding the contract method 0x7439ae59.
//
// Solidity: function eur() view returns(bytes32)
func (_MultiwordConsumer *MultiwordConsumerSession) Eur() ([32]byte, error) {
	return _MultiwordConsumer.Contract.Eur(&_MultiwordConsumer.CallOpts)
}

// Eur is a free data retrieval call binding the contract method 0x7439ae59.
//
// Solidity: function eur() view returns(bytes32)
func (_MultiwordConsumer *MultiwordConsumerCallerSession) Eur() ([32]byte, error) {
	return _MultiwordConsumer.Contract.Eur(&_MultiwordConsumer.CallOpts)
}

// Jpy is a free data retrieval call binding the contract method 0xfaa36761.
//
// Solidity: function jpy() view returns(bytes32)
func (_MultiwordConsumer *MultiwordConsumerCaller) Jpy(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MultiwordConsumer.contract.Call(opts, out, "jpy")
	return *ret0, err
}

// Jpy is a free data retrieval call binding the contract method 0xfaa36761.
//
// Solidity: function jpy() view returns(bytes32)
func (_MultiwordConsumer *MultiwordConsumerSession) Jpy() ([32]byte, error) {
	return _MultiwordConsumer.Contract.Jpy(&_MultiwordConsumer.CallOpts)
}

// Jpy is a free data retrieval call binding the contract method 0xfaa36761.
//
// Solidity: function jpy() view returns(bytes32)
func (_MultiwordConsumer *MultiwordConsumerCallerSession) Jpy() ([32]byte, error) {
	return _MultiwordConsumer.Contract.Jpy(&_MultiwordConsumer.CallOpts)
}

// Usd is a free data retrieval call binding the contract method 0xd63a6ccd.
//
// Solidity: function usd() view returns(bytes32)
func (_MultiwordConsumer *MultiwordConsumerCaller) Usd(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MultiwordConsumer.contract.Call(opts, out, "usd")
	return *ret0, err
}

// Usd is a free data retrieval call binding the contract method 0xd63a6ccd.
//
// Solidity: function usd() view returns(bytes32)
func (_MultiwordConsumer *MultiwordConsumerSession) Usd() ([32]byte, error) {
	return _MultiwordConsumer.Contract.Usd(&_MultiwordConsumer.CallOpts)
}

// Usd is a free data retrieval call binding the contract method 0xd63a6ccd.
//
// Solidity: function usd() view returns(bytes32)
func (_MultiwordConsumer *MultiwordConsumerCallerSession) Usd() ([32]byte, error) {
	return _MultiwordConsumer.Contract.Usd(&_MultiwordConsumer.CallOpts)
}

// AddExternalRequest is a paid mutator transaction binding the contract method 0xe8d5359d.
//
// Solidity: function addExternalRequest(address _oracle, bytes32 _requestId) returns()
func (_MultiwordConsumer *MultiwordConsumerTransactor) AddExternalRequest(opts *bind.TransactOpts, _oracle common.Address, _requestId [32]byte) (*types.Transaction, error) {
	return _MultiwordConsumer.contract.Transact(opts, "addExternalRequest", _oracle, _requestId)
}

// AddExternalRequest is a paid mutator transaction binding the contract method 0xe8d5359d.
//
// Solidity: function addExternalRequest(address _oracle, bytes32 _requestId) returns()
func (_MultiwordConsumer *MultiwordConsumerSession) AddExternalRequest(_oracle common.Address, _requestId [32]byte) (*types.Transaction, error) {
	return _MultiwordConsumer.Contract.AddExternalRequest(&_MultiwordConsumer.TransactOpts, _oracle, _requestId)
}

// AddExternalRequest is a paid mutator transaction binding the contract method 0xe8d5359d.
//
// Solidity: function addExternalRequest(address _oracle, bytes32 _requestId) returns()
func (_MultiwordConsumer *MultiwordConsumerTransactorSession) AddExternalRequest(_oracle common.Address, _requestId [32]byte) (*types.Transaction, error) {
	return _MultiwordConsumer.Contract.AddExternalRequest(&_MultiwordConsumer.TransactOpts, _oracle, _requestId)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x5591a608.
//
// Solidity: function cancelRequest(address _oracle, bytes32 _requestId, uint256 _payment, bytes4 _callbackFunctionId, uint256 _expiration) returns()
func (_MultiwordConsumer *MultiwordConsumerTransactor) CancelRequest(opts *bind.TransactOpts, _oracle common.Address, _requestId [32]byte, _payment *big.Int, _callbackFunctionId [4]byte, _expiration *big.Int) (*types.Transaction, error) {
	return _MultiwordConsumer.contract.Transact(opts, "cancelRequest", _oracle, _requestId, _payment, _callbackFunctionId, _expiration)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x5591a608.
//
// Solidity: function cancelRequest(address _oracle, bytes32 _requestId, uint256 _payment, bytes4 _callbackFunctionId, uint256 _expiration) returns()
func (_MultiwordConsumer *MultiwordConsumerSession) CancelRequest(_oracle common.Address, _requestId [32]byte, _payment *big.Int, _callbackFunctionId [4]byte, _expiration *big.Int) (*types.Transaction, error) {
	return _MultiwordConsumer.Contract.CancelRequest(&_MultiwordConsumer.TransactOpts, _oracle, _requestId, _payment, _callbackFunctionId, _expiration)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x5591a608.
//
// Solidity: function cancelRequest(address _oracle, bytes32 _requestId, uint256 _payment, bytes4 _callbackFunctionId, uint256 _expiration) returns()
func (_MultiwordConsumer *MultiwordConsumerTransactorSession) CancelRequest(_oracle common.Address, _requestId [32]byte, _payment *big.Int, _callbackFunctionId [4]byte, _expiration *big.Int) (*types.Transaction, error) {
	return _MultiwordConsumer.Contract.CancelRequest(&_MultiwordConsumer.TransactOpts, _oracle, _requestId, _payment, _callbackFunctionId, _expiration)
}

// FulfillBytes is a paid mutator transaction binding the contract method 0xc2fb8523.
//
// Solidity: function fulfillBytes(bytes32 _requestId, bytes _price) returns()
func (_MultiwordConsumer *MultiwordConsumerTransactor) FulfillBytes(opts *bind.TransactOpts, _requestId [32]byte, _price []byte) (*types.Transaction, error) {
	return _MultiwordConsumer.contract.Transact(opts, "fulfillBytes", _requestId, _price)
}

// FulfillBytes is a paid mutator transaction binding the contract method 0xc2fb8523.
//
// Solidity: function fulfillBytes(bytes32 _requestId, bytes _price) returns()
func (_MultiwordConsumer *MultiwordConsumerSession) FulfillBytes(_requestId [32]byte, _price []byte) (*types.Transaction, error) {
	return _MultiwordConsumer.Contract.FulfillBytes(&_MultiwordConsumer.TransactOpts, _requestId, _price)
}

// FulfillBytes is a paid mutator transaction binding the contract method 0xc2fb8523.
//
// Solidity: function fulfillBytes(bytes32 _requestId, bytes _price) returns()
func (_MultiwordConsumer *MultiwordConsumerTransactorSession) FulfillBytes(_requestId [32]byte, _price []byte) (*types.Transaction, error) {
	return _MultiwordConsumer.Contract.FulfillBytes(&_MultiwordConsumer.TransactOpts, _requestId, _price)
}

// FulfillMultipleParameters is a paid mutator transaction binding the contract method 0xa856ff6b.
//
// Solidity: function fulfillMultipleParameters(bytes32 _requestId, bytes32 _usd, bytes32 _eur, bytes32 _jpy) returns()
func (_MultiwordConsumer *MultiwordConsumerTransactor) FulfillMultipleParameters(opts *bind.TransactOpts, _requestId [32]byte, _usd [32]byte, _eur [32]byte, _jpy [32]byte) (*types.Transaction, error) {
	return _MultiwordConsumer.contract.Transact(opts, "fulfillMultipleParameters", _requestId, _usd, _eur, _jpy)
}

// FulfillMultipleParameters is a paid mutator transaction binding the contract method 0xa856ff6b.
//
// Solidity: function fulfillMultipleParameters(bytes32 _requestId, bytes32 _usd, bytes32 _eur, bytes32 _jpy) returns()
func (_MultiwordConsumer *MultiwordConsumerSession) FulfillMultipleParameters(_requestId [32]byte, _usd [32]byte, _eur [32]byte, _jpy [32]byte) (*types.Transaction, error) {
	return _MultiwordConsumer.Contract.FulfillMultipleParameters(&_MultiwordConsumer.TransactOpts, _requestId, _usd, _eur, _jpy)
}

// FulfillMultipleParameters is a paid mutator transaction binding the contract method 0xa856ff6b.
//
// Solidity: function fulfillMultipleParameters(bytes32 _requestId, bytes32 _usd, bytes32 _eur, bytes32 _jpy) returns()
func (_MultiwordConsumer *MultiwordConsumerTransactorSession) FulfillMultipleParameters(_requestId [32]byte, _usd [32]byte, _eur [32]byte, _jpy [32]byte) (*types.Transaction, error) {
	return _MultiwordConsumer.Contract.FulfillMultipleParameters(&_MultiwordConsumer.TransactOpts, _requestId, _usd, _eur, _jpy)
}

// RequestEthereumPrice is a paid mutator transaction binding the contract method 0x83db5cbc.
//
// Solidity: function requestEthereumPrice(string _currency, uint256 _payment) returns()
func (_MultiwordConsumer *MultiwordConsumerTransactor) RequestEthereumPrice(opts *bind.TransactOpts, _currency string, _payment *big.Int) (*types.Transaction, error) {
	return _MultiwordConsumer.contract.Transact(opts, "requestEthereumPrice", _currency, _payment)
}

// RequestEthereumPrice is a paid mutator transaction binding the contract method 0x83db5cbc.
//
// Solidity: function requestEthereumPrice(string _currency, uint256 _payment) returns()
func (_MultiwordConsumer *MultiwordConsumerSession) RequestEthereumPrice(_currency string, _payment *big.Int) (*types.Transaction, error) {
	return _MultiwordConsumer.Contract.RequestEthereumPrice(&_MultiwordConsumer.TransactOpts, _currency, _payment)
}

// RequestEthereumPrice is a paid mutator transaction binding the contract method 0x83db5cbc.
//
// Solidity: function requestEthereumPrice(string _currency, uint256 _payment) returns()
func (_MultiwordConsumer *MultiwordConsumerTransactorSession) RequestEthereumPrice(_currency string, _payment *big.Int) (*types.Transaction, error) {
	return _MultiwordConsumer.Contract.RequestEthereumPrice(&_MultiwordConsumer.TransactOpts, _currency, _payment)
}

// RequestEthereumPriceByCallback is a paid mutator transaction binding the contract method 0x74961d4d.
//
// Solidity: function requestEthereumPriceByCallback(string _currency, uint256 _payment, address _callback) returns()
func (_MultiwordConsumer *MultiwordConsumerTransactor) RequestEthereumPriceByCallback(opts *bind.TransactOpts, _currency string, _payment *big.Int, _callback common.Address) (*types.Transaction, error) {
	return _MultiwordConsumer.contract.Transact(opts, "requestEthereumPriceByCallback", _currency, _payment, _callback)
}

// RequestEthereumPriceByCallback is a paid mutator transaction binding the contract method 0x74961d4d.
//
// Solidity: function requestEthereumPriceByCallback(string _currency, uint256 _payment, address _callback) returns()
func (_MultiwordConsumer *MultiwordConsumerSession) RequestEthereumPriceByCallback(_currency string, _payment *big.Int, _callback common.Address) (*types.Transaction, error) {
	return _MultiwordConsumer.Contract.RequestEthereumPriceByCallback(&_MultiwordConsumer.TransactOpts, _currency, _payment, _callback)
}

// RequestEthereumPriceByCallback is a paid mutator transaction binding the contract method 0x74961d4d.
//
// Solidity: function requestEthereumPriceByCallback(string _currency, uint256 _payment, address _callback) returns()
func (_MultiwordConsumer *MultiwordConsumerTransactorSession) RequestEthereumPriceByCallback(_currency string, _payment *big.Int, _callback common.Address) (*types.Transaction, error) {
	return _MultiwordConsumer.Contract.RequestEthereumPriceByCallback(&_MultiwordConsumer.TransactOpts, _currency, _payment, _callback)
}

// RequestMultipleParameters is a paid mutator transaction binding the contract method 0xe89855ba.
//
// Solidity: function requestMultipleParameters(string _currency, uint256 _payment) returns()
func (_MultiwordConsumer *MultiwordConsumerTransactor) RequestMultipleParameters(opts *bind.TransactOpts, _currency string, _payment *big.Int) (*types.Transaction, error) {
	return _MultiwordConsumer.contract.Transact(opts, "requestMultipleParameters", _currency, _payment)
}

// RequestMultipleParameters is a paid mutator transaction binding the contract method 0xe89855ba.
//
// Solidity: function requestMultipleParameters(string _currency, uint256 _payment) returns()
func (_MultiwordConsumer *MultiwordConsumerSession) RequestMultipleParameters(_currency string, _payment *big.Int) (*types.Transaction, error) {
	return _MultiwordConsumer.Contract.RequestMultipleParameters(&_MultiwordConsumer.TransactOpts, _currency, _payment)
}

// RequestMultipleParameters is a paid mutator transaction binding the contract method 0xe89855ba.
//
// Solidity: function requestMultipleParameters(string _currency, uint256 _payment) returns()
func (_MultiwordConsumer *MultiwordConsumerTransactorSession) RequestMultipleParameters(_currency string, _payment *big.Int) (*types.Transaction, error) {
	return _MultiwordConsumer.Contract.RequestMultipleParameters(&_MultiwordConsumer.TransactOpts, _currency, _payment)
}

// SetSpecID is a paid mutator transaction binding the contract method 0x501fdd5d.
//
// Solidity: function setSpecID(bytes32 _specId) returns()
func (_MultiwordConsumer *MultiwordConsumerTransactor) SetSpecID(opts *bind.TransactOpts, _specId [32]byte) (*types.Transaction, error) {
	return _MultiwordConsumer.contract.Transact(opts, "setSpecID", _specId)
}

// SetSpecID is a paid mutator transaction binding the contract method 0x501fdd5d.
//
// Solidity: function setSpecID(bytes32 _specId) returns()
func (_MultiwordConsumer *MultiwordConsumerSession) SetSpecID(_specId [32]byte) (*types.Transaction, error) {
	return _MultiwordConsumer.Contract.SetSpecID(&_MultiwordConsumer.TransactOpts, _specId)
}

// SetSpecID is a paid mutator transaction binding the contract method 0x501fdd5d.
//
// Solidity: function setSpecID(bytes32 _specId) returns()
func (_MultiwordConsumer *MultiwordConsumerTransactorSession) SetSpecID(_specId [32]byte) (*types.Transaction, error) {
	return _MultiwordConsumer.Contract.SetSpecID(&_MultiwordConsumer.TransactOpts, _specId)
}

// WithdrawLink is a paid mutator transaction binding the contract method 0x8dc654a2.
//
// Solidity: function withdrawLink() returns()
func (_MultiwordConsumer *MultiwordConsumerTransactor) WithdrawLink(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiwordConsumer.contract.Transact(opts, "withdrawLink")
}

// WithdrawLink is a paid mutator transaction binding the contract method 0x8dc654a2.
//
// Solidity: function withdrawLink() returns()
func (_MultiwordConsumer *MultiwordConsumerSession) WithdrawLink() (*types.Transaction, error) {
	return _MultiwordConsumer.Contract.WithdrawLink(&_MultiwordConsumer.TransactOpts)
}

// WithdrawLink is a paid mutator transaction binding the contract method 0x8dc654a2.
//
// Solidity: function withdrawLink() returns()
func (_MultiwordConsumer *MultiwordConsumerTransactorSession) WithdrawLink() (*types.Transaction, error) {
	return _MultiwordConsumer.Contract.WithdrawLink(&_MultiwordConsumer.TransactOpts)
}

// MultiwordConsumerChainlinkCancelledIterator is returned from FilterChainlinkCancelled and is used to iterate over the raw logs and unpacked data for ChainlinkCancelled events raised by the MultiwordConsumer contract.
type MultiwordConsumerChainlinkCancelledIterator struct {
	Event *MultiwordConsumerChainlinkCancelled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiwordConsumerChainlinkCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiwordConsumerChainlinkCancelled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiwordConsumerChainlinkCancelled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiwordConsumerChainlinkCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiwordConsumerChainlinkCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiwordConsumerChainlinkCancelled represents a ChainlinkCancelled event raised by the MultiwordConsumer contract.
type MultiwordConsumerChainlinkCancelled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChainlinkCancelled is a free log retrieval operation binding the contract event 0xe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c5.
//
// Solidity: event ChainlinkCancelled(bytes32 indexed id)
func (_MultiwordConsumer *MultiwordConsumerFilterer) FilterChainlinkCancelled(opts *bind.FilterOpts, id [][32]byte) (*MultiwordConsumerChainlinkCancelledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MultiwordConsumer.contract.FilterLogs(opts, "ChainlinkCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return &MultiwordConsumerChainlinkCancelledIterator{contract: _MultiwordConsumer.contract, event: "ChainlinkCancelled", logs: logs, sub: sub}, nil
}

// WatchChainlinkCancelled is a free log subscription operation binding the contract event 0xe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c5.
//
// Solidity: event ChainlinkCancelled(bytes32 indexed id)
func (_MultiwordConsumer *MultiwordConsumerFilterer) WatchChainlinkCancelled(opts *bind.WatchOpts, sink chan<- *MultiwordConsumerChainlinkCancelled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MultiwordConsumer.contract.WatchLogs(opts, "ChainlinkCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiwordConsumerChainlinkCancelled)
				if err := _MultiwordConsumer.contract.UnpackLog(event, "ChainlinkCancelled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChainlinkCancelled is a log parse operation binding the contract event 0xe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c5.
//
// Solidity: event ChainlinkCancelled(bytes32 indexed id)
func (_MultiwordConsumer *MultiwordConsumerFilterer) ParseChainlinkCancelled(log types.Log) (*MultiwordConsumerChainlinkCancelled, error) {
	event := new(MultiwordConsumerChainlinkCancelled)
	if err := _MultiwordConsumer.contract.UnpackLog(event, "ChainlinkCancelled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MultiwordConsumerChainlinkFulfilledIterator is returned from FilterChainlinkFulfilled and is used to iterate over the raw logs and unpacked data for ChainlinkFulfilled events raised by the MultiwordConsumer contract.
type MultiwordConsumerChainlinkFulfilledIterator struct {
	Event *MultiwordConsumerChainlinkFulfilled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiwordConsumerChainlinkFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiwordConsumerChainlinkFulfilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiwordConsumerChainlinkFulfilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiwordConsumerChainlinkFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiwordConsumerChainlinkFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiwordConsumerChainlinkFulfilled represents a ChainlinkFulfilled event raised by the MultiwordConsumer contract.
type MultiwordConsumerChainlinkFulfilled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChainlinkFulfilled is a free log retrieval operation binding the contract event 0x7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a.
//
// Solidity: event ChainlinkFulfilled(bytes32 indexed id)
func (_MultiwordConsumer *MultiwordConsumerFilterer) FilterChainlinkFulfilled(opts *bind.FilterOpts, id [][32]byte) (*MultiwordConsumerChainlinkFulfilledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MultiwordConsumer.contract.FilterLogs(opts, "ChainlinkFulfilled", idRule)
	if err != nil {
		return nil, err
	}
	return &MultiwordConsumerChainlinkFulfilledIterator{contract: _MultiwordConsumer.contract, event: "ChainlinkFulfilled", logs: logs, sub: sub}, nil
}

// WatchChainlinkFulfilled is a free log subscription operation binding the contract event 0x7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a.
//
// Solidity: event ChainlinkFulfilled(bytes32 indexed id)
func (_MultiwordConsumer *MultiwordConsumerFilterer) WatchChainlinkFulfilled(opts *bind.WatchOpts, sink chan<- *MultiwordConsumerChainlinkFulfilled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MultiwordConsumer.contract.WatchLogs(opts, "ChainlinkFulfilled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiwordConsumerChainlinkFulfilled)
				if err := _MultiwordConsumer.contract.UnpackLog(event, "ChainlinkFulfilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChainlinkFulfilled is a log parse operation binding the contract event 0x7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a.
//
// Solidity: event ChainlinkFulfilled(bytes32 indexed id)
func (_MultiwordConsumer *MultiwordConsumerFilterer) ParseChainlinkFulfilled(log types.Log) (*MultiwordConsumerChainlinkFulfilled, error) {
	event := new(MultiwordConsumerChainlinkFulfilled)
	if err := _MultiwordConsumer.contract.UnpackLog(event, "ChainlinkFulfilled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MultiwordConsumerChainlinkRequestedIterator is returned from FilterChainlinkRequested and is used to iterate over the raw logs and unpacked data for ChainlinkRequested events raised by the MultiwordConsumer contract.
type MultiwordConsumerChainlinkRequestedIterator struct {
	Event *MultiwordConsumerChainlinkRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiwordConsumerChainlinkRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiwordConsumerChainlinkRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiwordConsumerChainlinkRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiwordConsumerChainlinkRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiwordConsumerChainlinkRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiwordConsumerChainlinkRequested represents a ChainlinkRequested event raised by the MultiwordConsumer contract.
type MultiwordConsumerChainlinkRequested struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChainlinkRequested is a free log retrieval operation binding the contract event 0xb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af9.
//
// Solidity: event ChainlinkRequested(bytes32 indexed id)
func (_MultiwordConsumer *MultiwordConsumerFilterer) FilterChainlinkRequested(opts *bind.FilterOpts, id [][32]byte) (*MultiwordConsumerChainlinkRequestedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MultiwordConsumer.contract.FilterLogs(opts, "ChainlinkRequested", idRule)
	if err != nil {
		return nil, err
	}
	return &MultiwordConsumerChainlinkRequestedIterator{contract: _MultiwordConsumer.contract, event: "ChainlinkRequested", logs: logs, sub: sub}, nil
}

// WatchChainlinkRequested is a free log subscription operation binding the contract event 0xb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af9.
//
// Solidity: event ChainlinkRequested(bytes32 indexed id)
func (_MultiwordConsumer *MultiwordConsumerFilterer) WatchChainlinkRequested(opts *bind.WatchOpts, sink chan<- *MultiwordConsumerChainlinkRequested, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MultiwordConsumer.contract.WatchLogs(opts, "ChainlinkRequested", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiwordConsumerChainlinkRequested)
				if err := _MultiwordConsumer.contract.UnpackLog(event, "ChainlinkRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChainlinkRequested is a log parse operation binding the contract event 0xb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af9.
//
// Solidity: event ChainlinkRequested(bytes32 indexed id)
func (_MultiwordConsumer *MultiwordConsumerFilterer) ParseChainlinkRequested(log types.Log) (*MultiwordConsumerChainlinkRequested, error) {
	event := new(MultiwordConsumerChainlinkRequested)
	if err := _MultiwordConsumer.contract.UnpackLog(event, "ChainlinkRequested", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MultiwordConsumerRequestFulfilledIterator is returned from FilterRequestFulfilled and is used to iterate over the raw logs and unpacked data for RequestFulfilled events raised by the MultiwordConsumer contract.
type MultiwordConsumerRequestFulfilledIterator struct {
	Event *MultiwordConsumerRequestFulfilled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiwordConsumerRequestFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiwordConsumerRequestFulfilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiwordConsumerRequestFulfilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiwordConsumerRequestFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiwordConsumerRequestFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiwordConsumerRequestFulfilled represents a RequestFulfilled event raised by the MultiwordConsumer contract.
type MultiwordConsumerRequestFulfilled struct {
	RequestId [32]byte
	Price     common.Hash
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestFulfilled is a free log retrieval operation binding the contract event 0x1a111c5dcf9a71088bd5e1797fdfaf399fec2afbb24aca247e4e3e9f4b61df91.
//
// Solidity: event RequestFulfilled(bytes32 indexed requestId, bytes indexed price)
func (_MultiwordConsumer *MultiwordConsumerFilterer) FilterRequestFulfilled(opts *bind.FilterOpts, requestId [][32]byte, price [][]byte) (*MultiwordConsumerRequestFulfilledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}

	logs, sub, err := _MultiwordConsumer.contract.FilterLogs(opts, "RequestFulfilled", requestIdRule, priceRule)
	if err != nil {
		return nil, err
	}
	return &MultiwordConsumerRequestFulfilledIterator{contract: _MultiwordConsumer.contract, event: "RequestFulfilled", logs: logs, sub: sub}, nil
}

// WatchRequestFulfilled is a free log subscription operation binding the contract event 0x1a111c5dcf9a71088bd5e1797fdfaf399fec2afbb24aca247e4e3e9f4b61df91.
//
// Solidity: event RequestFulfilled(bytes32 indexed requestId, bytes indexed price)
func (_MultiwordConsumer *MultiwordConsumerFilterer) WatchRequestFulfilled(opts *bind.WatchOpts, sink chan<- *MultiwordConsumerRequestFulfilled, requestId [][32]byte, price [][]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}

	logs, sub, err := _MultiwordConsumer.contract.WatchLogs(opts, "RequestFulfilled", requestIdRule, priceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiwordConsumerRequestFulfilled)
				if err := _MultiwordConsumer.contract.UnpackLog(event, "RequestFulfilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequestFulfilled is a log parse operation binding the contract event 0x1a111c5dcf9a71088bd5e1797fdfaf399fec2afbb24aca247e4e3e9f4b61df91.
//
// Solidity: event RequestFulfilled(bytes32 indexed requestId, bytes indexed price)
func (_MultiwordConsumer *MultiwordConsumerFilterer) ParseRequestFulfilled(log types.Log) (*MultiwordConsumerRequestFulfilled, error) {
	event := new(MultiwordConsumerRequestFulfilled)
	if err := _MultiwordConsumer.contract.UnpackLog(event, "RequestFulfilled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MultiwordConsumerRequestMultipleFulfilledIterator is returned from FilterRequestMultipleFulfilled and is used to iterate over the raw logs and unpacked data for RequestMultipleFulfilled events raised by the MultiwordConsumer contract.
type MultiwordConsumerRequestMultipleFulfilledIterator struct {
	Event *MultiwordConsumerRequestMultipleFulfilled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiwordConsumerRequestMultipleFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiwordConsumerRequestMultipleFulfilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiwordConsumerRequestMultipleFulfilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiwordConsumerRequestMultipleFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiwordConsumerRequestMultipleFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiwordConsumerRequestMultipleFulfilled represents a RequestMultipleFulfilled event raised by the MultiwordConsumer contract.
type MultiwordConsumerRequestMultipleFulfilled struct {
	RequestId [32]byte
	Usd       [32]byte
	Eur       [32]byte
	Jpy       [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestMultipleFulfilled is a free log retrieval operation binding the contract event 0x0ec0c13e44aa04198947078cb990660252870dd3363f4c4bb3cc780f808dabbe.
//
// Solidity: event RequestMultipleFulfilled(bytes32 indexed requestId, bytes32 indexed usd, bytes32 indexed eur, bytes32 jpy)
func (_MultiwordConsumer *MultiwordConsumerFilterer) FilterRequestMultipleFulfilled(opts *bind.FilterOpts, requestId [][32]byte, usd [][32]byte, eur [][32]byte) (*MultiwordConsumerRequestMultipleFulfilledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var usdRule []interface{}
	for _, usdItem := range usd {
		usdRule = append(usdRule, usdItem)
	}
	var eurRule []interface{}
	for _, eurItem := range eur {
		eurRule = append(eurRule, eurItem)
	}

	logs, sub, err := _MultiwordConsumer.contract.FilterLogs(opts, "RequestMultipleFulfilled", requestIdRule, usdRule, eurRule)
	if err != nil {
		return nil, err
	}
	return &MultiwordConsumerRequestMultipleFulfilledIterator{contract: _MultiwordConsumer.contract, event: "RequestMultipleFulfilled", logs: logs, sub: sub}, nil
}

// WatchRequestMultipleFulfilled is a free log subscription operation binding the contract event 0x0ec0c13e44aa04198947078cb990660252870dd3363f4c4bb3cc780f808dabbe.
//
// Solidity: event RequestMultipleFulfilled(bytes32 indexed requestId, bytes32 indexed usd, bytes32 indexed eur, bytes32 jpy)
func (_MultiwordConsumer *MultiwordConsumerFilterer) WatchRequestMultipleFulfilled(opts *bind.WatchOpts, sink chan<- *MultiwordConsumerRequestMultipleFulfilled, requestId [][32]byte, usd [][32]byte, eur [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var usdRule []interface{}
	for _, usdItem := range usd {
		usdRule = append(usdRule, usdItem)
	}
	var eurRule []interface{}
	for _, eurItem := range eur {
		eurRule = append(eurRule, eurItem)
	}

	logs, sub, err := _MultiwordConsumer.contract.WatchLogs(opts, "RequestMultipleFulfilled", requestIdRule, usdRule, eurRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiwordConsumerRequestMultipleFulfilled)
				if err := _MultiwordConsumer.contract.UnpackLog(event, "RequestMultipleFulfilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequestMultipleFulfilled is a log parse operation binding the contract event 0x0ec0c13e44aa04198947078cb990660252870dd3363f4c4bb3cc780f808dabbe.
//
// Solidity: event RequestMultipleFulfilled(bytes32 indexed requestId, bytes32 indexed usd, bytes32 indexed eur, bytes32 jpy)
func (_MultiwordConsumer *MultiwordConsumerFilterer) ParseRequestMultipleFulfilled(log types.Log) (*MultiwordConsumerRequestMultipleFulfilled, error) {
	event := new(MultiwordConsumerRequestMultipleFulfilled)
	if err := _MultiwordConsumer.contract.UnpackLog(event, "RequestMultipleFulfilled", log); err != nil {
		return nil, err
	}
	return event, nil
}
