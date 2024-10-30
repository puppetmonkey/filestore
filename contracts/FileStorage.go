// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"fileHashToFile\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"fileHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getFile\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_fileHash\",\"type\":\"string\"}],\"name\":\"getFileByHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserFileCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_fileHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ipfsHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_fileName\",\"type\":\"string\"}],\"name\":\"storeFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userFiles\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"fileHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5061141e8061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610060575f3560e01c80632bfda31314610064578063560be5c014610097578063718f38a6146100ca5780637dd77d2b146100fd578063cecaf69e14610119578063fe8712e71461014c575b5f5ffd5b61007e60048036038101906100799190610c05565b61016a565b60405161008e9493929190610caf565b60405180910390f35b6100b160048036038101906100ac9190610e33565b610435565b6040516100c19493929190610caf565b60405180910390f35b6100e460048036038101906100df9190610ed4565b61060b565b6040516100f49493929190610caf565b60405180910390f35b61011760048036038101906101129190610f12565b6107e2565b005b610133600480360381019061012e9190610e33565b610927565b6040516101439493929190610caf565b60405180910390f35b610154610b7b565b6040516101619190610fb6565b60405180910390f35b60608060605f5f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208054905085106101f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101e890611019565b60405180910390fd5b5f5f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2086815481106102405761023f611037565b5b905f5260205f2090600402016040518060800160405290815f8201805461026690611091565b80601f016020809104026020016040519081016040528092919081815260200182805461029290611091565b80156102dd5780601f106102b4576101008083540402835291602001916102dd565b820191905f5260205f20905b8154815290600101906020018083116102c057829003601f168201915b505050505081526020016001820180546102f690611091565b80601f016020809104026020016040519081016040528092919081815260200182805461032290611091565b801561036d5780601f106103445761010080835404028352916020019161036d565b820191905f5260205f20905b81548152906001019060200180831161035057829003601f168201915b5050505050815260200160028201805461038690611091565b80601f01602080910402602001604051908101604052809291908181526020018280546103b290611091565b80156103fd5780601f106103d4576101008083540402835291602001916103fd565b820191905f5260205f20905b8154815290600101906020018083116103e057829003601f168201915b505050505081526020016003820154815250509050805f01518160200151826040015183606001519450945094509450509193509193565b6001818051602081018201805184825260208301602085012081835280955050505050505f91509050805f01805461046c90611091565b80601f016020809104026020016040519081016040528092919081815260200182805461049890611091565b80156104e35780601f106104ba576101008083540402835291602001916104e3565b820191905f5260205f20905b8154815290600101906020018083116104c657829003601f168201915b5050505050908060010180546104f890611091565b80601f016020809104026020016040519081016040528092919081815260200182805461052490611091565b801561056f5780601f106105465761010080835404028352916020019161056f565b820191905f5260205f20905b81548152906001019060200180831161055257829003601f168201915b50505050509080600201805461058490611091565b80601f01602080910402602001604051908101604052809291908181526020018280546105b090611091565b80156105fb5780601f106105d2576101008083540402835291602001916105fb565b820191905f5260205f20905b8154815290600101906020018083116105de57829003601f168201915b5050505050908060030154905084565b5f602052815f5260405f208181548110610623575f80fd5b905f5260205f2090600402015f9150915050805f01805461064390611091565b80601f016020809104026020016040519081016040528092919081815260200182805461066f90611091565b80156106ba5780601f10610691576101008083540402835291602001916106ba565b820191905f5260205f20905b81548152906001019060200180831161069d57829003601f168201915b5050505050908060010180546106cf90611091565b80601f01602080910402602001604051908101604052809291908181526020018280546106fb90611091565b80156107465780601f1061071d57610100808354040283529160200191610746565b820191905f5260205f20905b81548152906001019060200180831161072957829003601f168201915b50505050509080600201805461075b90611091565b80601f016020809104026020016040519081016040528092919081815260200182805461078790611091565b80156107d25780601f106107a9576101008083540402835291602001916107d2565b820191905f5260205f20905b8154815290600101906020018083116107b557829003601f168201915b5050505050908060030154905084565b5f60405180608001604052808581526020018481526020018381526020014281525090505f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f01908161087c9190611261565b5060208201518160010190816108929190611261565b5060408201518160020190816108a89190611261565b50606082015181600301555050806001856040516108c6919061136a565b90815260200160405180910390205f820151815f0190816108e79190611261565b5060208201518160010190816108fd9190611261565b5060408201518160020190816109139190611261565b506060820151816003015590505050505050565b60608060605f5f60018660405161093e919061136a565b90815260200160405180910390206040518060800160405290815f8201805461096690611091565b80601f016020809104026020016040519081016040528092919081815260200182805461099290611091565b80156109dd5780601f106109b4576101008083540402835291602001916109dd565b820191905f5260205f20905b8154815290600101906020018083116109c057829003601f168201915b505050505081526020016001820180546109f690611091565b80601f0160208091040260200160405190810160405280929190818152602001828054610a2290611091565b8015610a6d5780601f10610a4457610100808354040283529160200191610a6d565b820191905f5260205f20905b815481529060010190602001808311610a5057829003601f168201915b50505050508152602001600282018054610a8690611091565b80601f0160208091040260200160405190810160405280929190818152602001828054610ab290611091565b8015610afd5780601f10610ad457610100808354040283529160200191610afd565b820191905f5260205f20905b815481529060010190602001808311610ae057829003601f168201915b5050505050815260200160038201548152505090505f815f01515103610b58576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b4f906113ca565b60405180910390fd5b805f01518160200151826040015183606001519450945094509450509193509193565b5f5f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2080549050905090565b5f604051905090565b5f5ffd5b5f5ffd5b5f819050919050565b610be481610bd2565b8114610bee575f5ffd5b50565b5f81359050610bff81610bdb565b92915050565b5f60208284031215610c1a57610c19610bca565b5b5f610c2784828501610bf1565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610c7282610c30565b610c7c8185610c3a565b9350610c8c818560208601610c4a565b610c9581610c58565b840191505092915050565b610ca981610bd2565b82525050565b5f6080820190508181035f830152610cc78187610c68565b90508181036020830152610cdb8186610c68565b90508181036040830152610cef8185610c68565b9050610cfe6060830184610ca0565b95945050505050565b5f5ffd5b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610d4582610c58565b810181811067ffffffffffffffff82111715610d6457610d63610d0f565b5b80604052505050565b5f610d76610bc1565b9050610d828282610d3c565b919050565b5f67ffffffffffffffff821115610da157610da0610d0f565b5b610daa82610c58565b9050602081019050919050565b828183375f83830152505050565b5f610dd7610dd284610d87565b610d6d565b905082815260208101848484011115610df357610df2610d0b565b5b610dfe848285610db7565b509392505050565b5f82601f830112610e1a57610e19610d07565b5b8135610e2a848260208601610dc5565b91505092915050565b5f60208284031215610e4857610e47610bca565b5b5f82013567ffffffffffffffff811115610e6557610e64610bce565b5b610e7184828501610e06565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610ea382610e7a565b9050919050565b610eb381610e99565b8114610ebd575f5ffd5b50565b5f81359050610ece81610eaa565b92915050565b5f5f60408385031215610eea57610ee9610bca565b5b5f610ef785828601610ec0565b9250506020610f0885828601610bf1565b9150509250929050565b5f5f5f60608486031215610f2957610f28610bca565b5b5f84013567ffffffffffffffff811115610f4657610f45610bce565b5b610f5286828701610e06565b935050602084013567ffffffffffffffff811115610f7357610f72610bce565b5b610f7f86828701610e06565b925050604084013567ffffffffffffffff811115610fa057610f9f610bce565b5b610fac86828701610e06565b9150509250925092565b5f602082019050610fc95f830184610ca0565b92915050565b7f496e76616c69642066696c6520696e6465782e000000000000000000000000005f82015250565b5f611003601383610c3a565b915061100e82610fcf565b602082019050919050565b5f6020820190508181035f83015261103081610ff7565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806110a857607f821691505b6020821081036110bb576110ba611064565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261111d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826110e2565b61112786836110e2565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61116261115d61115884610bd2565b61113f565b610bd2565b9050919050565b5f819050919050565b61117b83611148565b61118f61118782611169565b8484546110ee565b825550505050565b5f5f905090565b6111a6611197565b6111b1818484611172565b505050565b5b818110156111d4576111c95f8261119e565b6001810190506111b7565b5050565b601f821115611219576111ea816110c1565b6111f3846110d3565b81016020851015611202578190505b61121661120e856110d3565b8301826111b6565b50505b505050565b5f82821c905092915050565b5f6112395f198460080261121e565b1980831691505092915050565b5f611251838361122a565b9150826002028217905092915050565b61126a82610c30565b67ffffffffffffffff81111561128357611282610d0f565b5b61128d8254611091565b6112988282856111d8565b5f60209050601f8311600181146112c9575f84156112b7578287015190505b6112c18582611246565b865550611328565b601f1984166112d7866110c1565b5f5b828110156112fe578489015182556001820191506020850194506020810190506112d9565b8683101561131b5784890151611317601f89168261122a565b8355505b6001600288020188555050505b505050505050565b5f81905092915050565b5f61134482610c30565b61134e8185611330565b935061135e818560208601610c4a565b80840191505092915050565b5f611375828461133a565b915081905092915050565b7f46696c65206e6f7420666f756e642e00000000000000000000000000000000005f82015250565b5f6113b4600f83610c3a565b91506113bf82611380565b602082019050919050565b5f6020820190508181035f8301526113e1816113a8565b905091905056fea264697066735822122092dcd061388821c15d02cb2d0a6f6eb5c11778dada636f3ec68223c63b1e476e64736f6c634300081b0033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// FileHashToFile is a free data retrieval call binding the contract method 0x560be5c0.
//
// Solidity: function fileHashToFile(string ) view returns(string fileHash, string ipfsHash, string fileName, uint256 timestamp)
func (_Contract *ContractCaller) FileHashToFile(opts *bind.CallOpts, arg0 string) (struct {
	FileHash  string
	IpfsHash  string
	FileName  string
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "fileHashToFile", arg0)

	outstruct := new(struct {
		FileHash  string
		IpfsHash  string
		FileName  string
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FileHash = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.IpfsHash = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.FileName = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// FileHashToFile is a free data retrieval call binding the contract method 0x560be5c0.
//
// Solidity: function fileHashToFile(string ) view returns(string fileHash, string ipfsHash, string fileName, uint256 timestamp)
func (_Contract *ContractSession) FileHashToFile(arg0 string) (struct {
	FileHash  string
	IpfsHash  string
	FileName  string
	Timestamp *big.Int
}, error) {
	return _Contract.Contract.FileHashToFile(&_Contract.CallOpts, arg0)
}

// FileHashToFile is a free data retrieval call binding the contract method 0x560be5c0.
//
// Solidity: function fileHashToFile(string ) view returns(string fileHash, string ipfsHash, string fileName, uint256 timestamp)
func (_Contract *ContractCallerSession) FileHashToFile(arg0 string) (struct {
	FileHash  string
	IpfsHash  string
	FileName  string
	Timestamp *big.Int
}, error) {
	return _Contract.Contract.FileHashToFile(&_Contract.CallOpts, arg0)
}

// GetFile is a free data retrieval call binding the contract method 0x2bfda313.
//
// Solidity: function getFile(uint256 _index) view returns(string, string, string, uint256)
func (_Contract *ContractCaller) GetFile(opts *bind.CallOpts, _index *big.Int) (string, string, string, *big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getFile", _index)

	if err != nil {
		return *new(string), *new(string), *new(string), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetFile is a free data retrieval call binding the contract method 0x2bfda313.
//
// Solidity: function getFile(uint256 _index) view returns(string, string, string, uint256)
func (_Contract *ContractSession) GetFile(_index *big.Int) (string, string, string, *big.Int, error) {
	return _Contract.Contract.GetFile(&_Contract.CallOpts, _index)
}

// GetFile is a free data retrieval call binding the contract method 0x2bfda313.
//
// Solidity: function getFile(uint256 _index) view returns(string, string, string, uint256)
func (_Contract *ContractCallerSession) GetFile(_index *big.Int) (string, string, string, *big.Int, error) {
	return _Contract.Contract.GetFile(&_Contract.CallOpts, _index)
}

// GetFileByHash is a free data retrieval call binding the contract method 0xcecaf69e.
//
// Solidity: function getFileByHash(string _fileHash) view returns(string, string, string, uint256)
func (_Contract *ContractCaller) GetFileByHash(opts *bind.CallOpts, _fileHash string) (string, string, string, *big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getFileByHash", _fileHash)

	if err != nil {
		return *new(string), *new(string), *new(string), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetFileByHash is a free data retrieval call binding the contract method 0xcecaf69e.
//
// Solidity: function getFileByHash(string _fileHash) view returns(string, string, string, uint256)
func (_Contract *ContractSession) GetFileByHash(_fileHash string) (string, string, string, *big.Int, error) {
	return _Contract.Contract.GetFileByHash(&_Contract.CallOpts, _fileHash)
}

// GetFileByHash is a free data retrieval call binding the contract method 0xcecaf69e.
//
// Solidity: function getFileByHash(string _fileHash) view returns(string, string, string, uint256)
func (_Contract *ContractCallerSession) GetFileByHash(_fileHash string) (string, string, string, *big.Int, error) {
	return _Contract.Contract.GetFileByHash(&_Contract.CallOpts, _fileHash)
}

// GetUserFileCount is a free data retrieval call binding the contract method 0xfe8712e7.
//
// Solidity: function getUserFileCount() view returns(uint256)
func (_Contract *ContractCaller) GetUserFileCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getUserFileCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserFileCount is a free data retrieval call binding the contract method 0xfe8712e7.
//
// Solidity: function getUserFileCount() view returns(uint256)
func (_Contract *ContractSession) GetUserFileCount() (*big.Int, error) {
	return _Contract.Contract.GetUserFileCount(&_Contract.CallOpts)
}

// GetUserFileCount is a free data retrieval call binding the contract method 0xfe8712e7.
//
// Solidity: function getUserFileCount() view returns(uint256)
func (_Contract *ContractCallerSession) GetUserFileCount() (*big.Int, error) {
	return _Contract.Contract.GetUserFileCount(&_Contract.CallOpts)
}

// UserFiles is a free data retrieval call binding the contract method 0x718f38a6.
//
// Solidity: function userFiles(address , uint256 ) view returns(string fileHash, string ipfsHash, string fileName, uint256 timestamp)
func (_Contract *ContractCaller) UserFiles(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	FileHash  string
	IpfsHash  string
	FileName  string
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "userFiles", arg0, arg1)

	outstruct := new(struct {
		FileHash  string
		IpfsHash  string
		FileName  string
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FileHash = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.IpfsHash = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.FileName = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserFiles is a free data retrieval call binding the contract method 0x718f38a6.
//
// Solidity: function userFiles(address , uint256 ) view returns(string fileHash, string ipfsHash, string fileName, uint256 timestamp)
func (_Contract *ContractSession) UserFiles(arg0 common.Address, arg1 *big.Int) (struct {
	FileHash  string
	IpfsHash  string
	FileName  string
	Timestamp *big.Int
}, error) {
	return _Contract.Contract.UserFiles(&_Contract.CallOpts, arg0, arg1)
}

// UserFiles is a free data retrieval call binding the contract method 0x718f38a6.
//
// Solidity: function userFiles(address , uint256 ) view returns(string fileHash, string ipfsHash, string fileName, uint256 timestamp)
func (_Contract *ContractCallerSession) UserFiles(arg0 common.Address, arg1 *big.Int) (struct {
	FileHash  string
	IpfsHash  string
	FileName  string
	Timestamp *big.Int
}, error) {
	return _Contract.Contract.UserFiles(&_Contract.CallOpts, arg0, arg1)
}

// StoreFile is a paid mutator transaction binding the contract method 0x7dd77d2b.
//
// Solidity: function storeFile(string _fileHash, string _ipfsHash, string _fileName) returns()
func (_Contract *ContractTransactor) StoreFile(opts *bind.TransactOpts, _fileHash string, _ipfsHash string, _fileName string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "storeFile", _fileHash, _ipfsHash, _fileName)
}

// StoreFile is a paid mutator transaction binding the contract method 0x7dd77d2b.
//
// Solidity: function storeFile(string _fileHash, string _ipfsHash, string _fileName) returns()
func (_Contract *ContractSession) StoreFile(_fileHash string, _ipfsHash string, _fileName string) (*types.Transaction, error) {
	return _Contract.Contract.StoreFile(&_Contract.TransactOpts, _fileHash, _ipfsHash, _fileName)
}

// StoreFile is a paid mutator transaction binding the contract method 0x7dd77d2b.
//
// Solidity: function storeFile(string _fileHash, string _ipfsHash, string _fileName) returns()
func (_Contract *ContractTransactorSession) StoreFile(_fileHash string, _ipfsHash string, _fileName string) (*types.Transaction, error) {
	return _Contract.Contract.StoreFile(&_Contract.TransactOpts, _fileHash, _ipfsHash, _fileName)
}
