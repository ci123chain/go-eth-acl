// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package acl_sdk

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

// ACLABI is the input ABI used to generate the binding from.
const ACLABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"msgsender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"thiscontract\",\"type\":\"address\"}],\"name\":\"ACLAddr\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"ChangePermissionManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"entity\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"SetPermission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"entity\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"paramsHash\",\"type\":\"bytes32\"}],\"name\":\"SetPermissionParams\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"ANY_ENTITY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BURN_ENTITY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CREATE_PERMISSIONS_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EMPTY_PARAM_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NO_PERMISSION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"_params\",\"type\":\"uint256[]\"}],\"name\":\"canPerform\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitializationBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isPetrified\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIACL\",\"name\":\"_acl\",\"type\":\"address\"}],\"name\":\"setACL\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permissionsCreator\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_entity\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_app\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"createPermission\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_entity\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_app\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_role\",\"type\":\"bytes32\"}],\"name\":\"grantPermission\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_entity\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_app\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"_params\",\"type\":\"uint256[]\"}],\"name\":\"grantPermissionP\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_entity\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_app\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_role\",\"type\":\"bytes32\"}],\"name\":\"revokePermission\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_app\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_role\",\"type\":\"bytes32\"}],\"name\":\"setPermissionManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_app\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_role\",\"type\":\"bytes32\"}],\"name\":\"removePermissionManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_app\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_role\",\"type\":\"bytes32\"}],\"name\":\"createBurnedPermission\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_app\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_role\",\"type\":\"bytes32\"}],\"name\":\"burnPermissionManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_entity\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_app\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_role\",\"type\":\"bytes32\"}],\"name\":\"getPermissionParamsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_entity\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_app\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getPermissionParam\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint240\",\"name\":\"\",\"type\":\"uint240\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_app\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_role\",\"type\":\"bytes32\"}],\"name\":\"getPermissionManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_where\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_what\",\"type\":\"bytes32\"}],\"name\":\"hasPermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_where\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"_how\",\"type\":\"uint256[]\"}],\"name\":\"hasPermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_where\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_what\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_how\",\"type\":\"bytes\"}],\"name\":\"hasPermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_paramsHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_where\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"_how\",\"type\":\"uint256[]\"}],\"name\":\"evalParams\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ACL is an auto generated Go binding around an Ethereum contract.
type ACL struct {
	ACLCaller     // Read-only binding to the contract
	ACLTransactor // Write-only binding to the contract
	ACLFilterer   // Log filterer for contract events
}

// ACLCaller is an auto generated read-only Go binding around an Ethereum contract.
type ACLCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ACLTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ACLTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ACLFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ACLFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ACLSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ACLSession struct {
	Contract     *ACL              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ACLCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ACLCallerSession struct {
	Contract *ACLCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ACLTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ACLTransactorSession struct {
	Contract     *ACLTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ACLRaw is an auto generated low-level Go binding around an Ethereum contract.
type ACLRaw struct {
	Contract *ACL // Generic contract binding to access the raw methods on
}

// ACLCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ACLCallerRaw struct {
	Contract *ACLCaller // Generic read-only contract binding to access the raw methods on
}

// ACLTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ACLTransactorRaw struct {
	Contract *ACLTransactor // Generic write-only contract binding to access the raw methods on
}

// NewACL creates a new instance of ACL, bound to a specific deployed contract.
func NewACL(address common.Address, backend bind.ContractBackend) (*ACL, error) {
	contract, err := bindACL(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ACL{ACLCaller: ACLCaller{contract: contract}, ACLTransactor: ACLTransactor{contract: contract}, ACLFilterer: ACLFilterer{contract: contract}}, nil
}

// NewACLCaller creates a new read-only instance of ACL, bound to a specific deployed contract.
func NewACLCaller(address common.Address, caller bind.ContractCaller) (*ACLCaller, error) {
	contract, err := bindACL(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ACLCaller{contract: contract}, nil
}

// NewACLTransactor creates a new write-only instance of ACL, bound to a specific deployed contract.
func NewACLTransactor(address common.Address, transactor bind.ContractTransactor) (*ACLTransactor, error) {
	contract, err := bindACL(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ACLTransactor{contract: contract}, nil
}

// NewACLFilterer creates a new log filterer instance of ACL, bound to a specific deployed contract.
func NewACLFilterer(address common.Address, filterer bind.ContractFilterer) (*ACLFilterer, error) {
	contract, err := bindACL(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ACLFilterer{contract: contract}, nil
}

// bindACL binds a generic wrapper to an already deployed contract.
func bindACL(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ACLABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ACL *ACLRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ACL.Contract.ACLCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ACL *ACLRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ACL.Contract.ACLTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ACL *ACLRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ACL.Contract.ACLTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ACL *ACLCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ACL.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ACL *ACLTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ACL.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ACL *ACLTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ACL.Contract.contract.Transact(opts, method, params...)
}

// ANYENTITY is a free data retrieval call binding the contract method 0xa5ed8bf8.
//
// Solidity: function ANY_ENTITY() view returns(address)
func (_ACL *ACLCaller) ANYENTITY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ACL.contract.Call(opts, &out, "ANY_ENTITY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ANYENTITY is a free data retrieval call binding the contract method 0xa5ed8bf8.
//
// Solidity: function ANY_ENTITY() view returns(address)
func (_ACL *ACLSession) ANYENTITY() (common.Address, error) {
	return _ACL.Contract.ANYENTITY(&_ACL.CallOpts)
}

// ANYENTITY is a free data retrieval call binding the contract method 0xa5ed8bf8.
//
// Solidity: function ANY_ENTITY() view returns(address)
func (_ACL *ACLCallerSession) ANYENTITY() (common.Address, error) {
	return _ACL.Contract.ANYENTITY(&_ACL.CallOpts)
}

// BURNENTITY is a free data retrieval call binding the contract method 0xf516bc0e.
//
// Solidity: function BURN_ENTITY() view returns(address)
func (_ACL *ACLCaller) BURNENTITY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ACL.contract.Call(opts, &out, "BURN_ENTITY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNENTITY is a free data retrieval call binding the contract method 0xf516bc0e.
//
// Solidity: function BURN_ENTITY() view returns(address)
func (_ACL *ACLSession) BURNENTITY() (common.Address, error) {
	return _ACL.Contract.BURNENTITY(&_ACL.CallOpts)
}

// BURNENTITY is a free data retrieval call binding the contract method 0xf516bc0e.
//
// Solidity: function BURN_ENTITY() view returns(address)
func (_ACL *ACLCallerSession) BURNENTITY() (common.Address, error) {
	return _ACL.Contract.BURNENTITY(&_ACL.CallOpts)
}

// CREATEPERMISSIONSROLE is a free data retrieval call binding the contract method 0x3d6ab68f.
//
// Solidity: function CREATE_PERMISSIONS_ROLE() view returns(bytes32)
func (_ACL *ACLCaller) CREATEPERMISSIONSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ACL.contract.Call(opts, &out, "CREATE_PERMISSIONS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CREATEPERMISSIONSROLE is a free data retrieval call binding the contract method 0x3d6ab68f.
//
// Solidity: function CREATE_PERMISSIONS_ROLE() view returns(bytes32)
func (_ACL *ACLSession) CREATEPERMISSIONSROLE() ([32]byte, error) {
	return _ACL.Contract.CREATEPERMISSIONSROLE(&_ACL.CallOpts)
}

// CREATEPERMISSIONSROLE is a free data retrieval call binding the contract method 0x3d6ab68f.
//
// Solidity: function CREATE_PERMISSIONS_ROLE() view returns(bytes32)
func (_ACL *ACLCallerSession) CREATEPERMISSIONSROLE() ([32]byte, error) {
	return _ACL.Contract.CREATEPERMISSIONSROLE(&_ACL.CallOpts)
}

// EMPTYPARAMHASH is a free data retrieval call binding the contract method 0xc513f66e.
//
// Solidity: function EMPTY_PARAM_HASH() view returns(bytes32)
func (_ACL *ACLCaller) EMPTYPARAMHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ACL.contract.Call(opts, &out, "EMPTY_PARAM_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EMPTYPARAMHASH is a free data retrieval call binding the contract method 0xc513f66e.
//
// Solidity: function EMPTY_PARAM_HASH() view returns(bytes32)
func (_ACL *ACLSession) EMPTYPARAMHASH() ([32]byte, error) {
	return _ACL.Contract.EMPTYPARAMHASH(&_ACL.CallOpts)
}

// EMPTYPARAMHASH is a free data retrieval call binding the contract method 0xc513f66e.
//
// Solidity: function EMPTY_PARAM_HASH() view returns(bytes32)
func (_ACL *ACLCallerSession) EMPTYPARAMHASH() ([32]byte, error) {
	return _ACL.Contract.EMPTYPARAMHASH(&_ACL.CallOpts)
}

// NOPERMISSION is a free data retrieval call binding the contract method 0x1d63ff2b.
//
// Solidity: function NO_PERMISSION() view returns(bytes32)
func (_ACL *ACLCaller) NOPERMISSION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ACL.contract.Call(opts, &out, "NO_PERMISSION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NOPERMISSION is a free data retrieval call binding the contract method 0x1d63ff2b.
//
// Solidity: function NO_PERMISSION() view returns(bytes32)
func (_ACL *ACLSession) NOPERMISSION() ([32]byte, error) {
	return _ACL.Contract.NOPERMISSION(&_ACL.CallOpts)
}

// NOPERMISSION is a free data retrieval call binding the contract method 0x1d63ff2b.
//
// Solidity: function NO_PERMISSION() view returns(bytes32)
func (_ACL *ACLCallerSession) NOPERMISSION() ([32]byte, error) {
	return _ACL.Contract.NOPERMISSION(&_ACL.CallOpts)
}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_ACL *ACLCaller) CanPerform(opts *bind.CallOpts, _sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	var out []interface{}
	err := _ACL.contract.Call(opts, &out, "canPerform", _sender, _role, _params)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_ACL *ACLSession) CanPerform(_sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	return _ACL.Contract.CanPerform(&_ACL.CallOpts, _sender, _role, _params)
}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_ACL *ACLCallerSession) CanPerform(_sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	return _ACL.Contract.CanPerform(&_ACL.CallOpts, _sender, _role, _params)
}

// EvalParams is a free data retrieval call binding the contract method 0x1b5e75be.
//
// Solidity: function evalParams(bytes32 _paramsHash, address _who, address _where, bytes32 _what, uint256[] _how) view returns(bool)
func (_ACL *ACLCaller) EvalParams(opts *bind.CallOpts, _paramsHash [32]byte, _who common.Address, _where common.Address, _what [32]byte, _how []*big.Int) (bool, error) {
	var out []interface{}
	err := _ACL.contract.Call(opts, &out, "evalParams", _paramsHash, _who, _where, _what, _how)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EvalParams is a free data retrieval call binding the contract method 0x1b5e75be.
//
// Solidity: function evalParams(bytes32 _paramsHash, address _who, address _where, bytes32 _what, uint256[] _how) view returns(bool)
func (_ACL *ACLSession) EvalParams(_paramsHash [32]byte, _who common.Address, _where common.Address, _what [32]byte, _how []*big.Int) (bool, error) {
	return _ACL.Contract.EvalParams(&_ACL.CallOpts, _paramsHash, _who, _where, _what, _how)
}

// EvalParams is a free data retrieval call binding the contract method 0x1b5e75be.
//
// Solidity: function evalParams(bytes32 _paramsHash, address _who, address _where, bytes32 _what, uint256[] _how) view returns(bool)
func (_ACL *ACLCallerSession) EvalParams(_paramsHash [32]byte, _who common.Address, _where common.Address, _what [32]byte, _how []*big.Int) (bool, error) {
	return _ACL.Contract.EvalParams(&_ACL.CallOpts, _paramsHash, _who, _where, _what, _how)
}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_ACL *ACLCaller) GetInitializationBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ACL.contract.Call(opts, &out, "getInitializationBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_ACL *ACLSession) GetInitializationBlock() (*big.Int, error) {
	return _ACL.Contract.GetInitializationBlock(&_ACL.CallOpts)
}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_ACL *ACLCallerSession) GetInitializationBlock() (*big.Int, error) {
	return _ACL.Contract.GetInitializationBlock(&_ACL.CallOpts)
}

// GetPermissionManager is a free data retrieval call binding the contract method 0xb1905727.
//
// Solidity: function getPermissionManager(address _app, bytes32 _role) view returns(address)
func (_ACL *ACLCaller) GetPermissionManager(opts *bind.CallOpts, _app common.Address, _role [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ACL.contract.Call(opts, &out, "getPermissionManager", _app, _role)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPermissionManager is a free data retrieval call binding the contract method 0xb1905727.
//
// Solidity: function getPermissionManager(address _app, bytes32 _role) view returns(address)
func (_ACL *ACLSession) GetPermissionManager(_app common.Address, _role [32]byte) (common.Address, error) {
	return _ACL.Contract.GetPermissionManager(&_ACL.CallOpts, _app, _role)
}

// GetPermissionManager is a free data retrieval call binding the contract method 0xb1905727.
//
// Solidity: function getPermissionManager(address _app, bytes32 _role) view returns(address)
func (_ACL *ACLCallerSession) GetPermissionManager(_app common.Address, _role [32]byte) (common.Address, error) {
	return _ACL.Contract.GetPermissionManager(&_ACL.CallOpts, _app, _role)
}

// GetPermissionParam is a free data retrieval call binding the contract method 0xa03c5832.
//
// Solidity: function getPermissionParam(address _entity, address _app, bytes32 _role, uint256 _index) view returns(uint8, uint8, uint240)
func (_ACL *ACLCaller) GetPermissionParam(opts *bind.CallOpts, _entity common.Address, _app common.Address, _role [32]byte, _index *big.Int) (uint8, uint8, *big.Int, error) {
	var out []interface{}
	err := _ACL.contract.Call(opts, &out, "getPermissionParam", _entity, _app, _role, _index)

	if err != nil {
		return *new(uint8), *new(uint8), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetPermissionParam is a free data retrieval call binding the contract method 0xa03c5832.
//
// Solidity: function getPermissionParam(address _entity, address _app, bytes32 _role, uint256 _index) view returns(uint8, uint8, uint240)
func (_ACL *ACLSession) GetPermissionParam(_entity common.Address, _app common.Address, _role [32]byte, _index *big.Int) (uint8, uint8, *big.Int, error) {
	return _ACL.Contract.GetPermissionParam(&_ACL.CallOpts, _entity, _app, _role, _index)
}

// GetPermissionParam is a free data retrieval call binding the contract method 0xa03c5832.
//
// Solidity: function getPermissionParam(address _entity, address _app, bytes32 _role, uint256 _index) view returns(uint8, uint8, uint240)
func (_ACL *ACLCallerSession) GetPermissionParam(_entity common.Address, _app common.Address, _role [32]byte, _index *big.Int) (uint8, uint8, *big.Int, error) {
	return _ACL.Contract.GetPermissionParam(&_ACL.CallOpts, _entity, _app, _role, _index)
}

// GetPermissionParamsLength is a free data retrieval call binding the contract method 0x15949ed7.
//
// Solidity: function getPermissionParamsLength(address _entity, address _app, bytes32 _role) view returns(uint256)
func (_ACL *ACLCaller) GetPermissionParamsLength(opts *bind.CallOpts, _entity common.Address, _app common.Address, _role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ACL.contract.Call(opts, &out, "getPermissionParamsLength", _entity, _app, _role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPermissionParamsLength is a free data retrieval call binding the contract method 0x15949ed7.
//
// Solidity: function getPermissionParamsLength(address _entity, address _app, bytes32 _role) view returns(uint256)
func (_ACL *ACLSession) GetPermissionParamsLength(_entity common.Address, _app common.Address, _role [32]byte) (*big.Int, error) {
	return _ACL.Contract.GetPermissionParamsLength(&_ACL.CallOpts, _entity, _app, _role)
}

// GetPermissionParamsLength is a free data retrieval call binding the contract method 0x15949ed7.
//
// Solidity: function getPermissionParamsLength(address _entity, address _app, bytes32 _role) view returns(uint256)
func (_ACL *ACLCallerSession) GetPermissionParamsLength(_entity common.Address, _app common.Address, _role [32]byte) (*big.Int, error) {
	return _ACL.Contract.GetPermissionParamsLength(&_ACL.CallOpts, _entity, _app, _role)
}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_ACL *ACLCaller) HasInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ACL.contract.Call(opts, &out, "hasInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_ACL *ACLSession) HasInitialized() (bool, error) {
	return _ACL.Contract.HasInitialized(&_ACL.CallOpts)
}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_ACL *ACLCallerSession) HasInitialized() (bool, error) {
	return _ACL.Contract.HasInitialized(&_ACL.CallOpts)
}

// HasPermission is a free data retrieval call binding the contract method 0x6d6712d8.
//
// Solidity: function hasPermission(address _who, address _where, bytes32 _what) view returns(bool)
func (_ACL *ACLCaller) HasPermission(opts *bind.CallOpts, _who common.Address, _where common.Address, _what [32]byte) (bool, error) {
	var out []interface{}
	err := _ACL.contract.Call(opts, &out, "hasPermission", _who, _where, _what)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPermission is a free data retrieval call binding the contract method 0x6d6712d8.
//
// Solidity: function hasPermission(address _who, address _where, bytes32 _what) view returns(bool)
func (_ACL *ACLSession) HasPermission(_who common.Address, _where common.Address, _what [32]byte) (bool, error) {
	return _ACL.Contract.HasPermission(&_ACL.CallOpts, _who, _where, _what)
}

// HasPermission is a free data retrieval call binding the contract method 0x6d6712d8.
//
// Solidity: function hasPermission(address _who, address _where, bytes32 _what) view returns(bool)
func (_ACL *ACLCallerSession) HasPermission(_who common.Address, _where common.Address, _what [32]byte) (bool, error) {
	return _ACL.Contract.HasPermission(&_ACL.CallOpts, _who, _where, _what)
}

// HasPermission0 is a free data retrieval call binding the contract method 0xf520b58d.
//
// Solidity: function hasPermission(address _who, address _where, bytes32 _what, uint256[] _how) view returns(bool)
func (_ACL *ACLCaller) HasPermission0(opts *bind.CallOpts, _who common.Address, _where common.Address, _what [32]byte, _how []*big.Int) (bool, error) {
	var out []interface{}
	err := _ACL.contract.Call(opts, &out, "hasPermission0", _who, _where, _what, _how)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPermission0 is a free data retrieval call binding the contract method 0xf520b58d.
//
// Solidity: function hasPermission(address _who, address _where, bytes32 _what, uint256[] _how) view returns(bool)
func (_ACL *ACLSession) HasPermission0(_who common.Address, _where common.Address, _what [32]byte, _how []*big.Int) (bool, error) {
	return _ACL.Contract.HasPermission0(&_ACL.CallOpts, _who, _where, _what, _how)
}

// HasPermission0 is a free data retrieval call binding the contract method 0xf520b58d.
//
// Solidity: function hasPermission(address _who, address _where, bytes32 _what, uint256[] _how) view returns(bool)
func (_ACL *ACLCallerSession) HasPermission0(_who common.Address, _where common.Address, _what [32]byte, _how []*big.Int) (bool, error) {
	return _ACL.Contract.HasPermission0(&_ACL.CallOpts, _who, _where, _what, _how)
}

// HasPermission1 is a free data retrieval call binding the contract method 0xfdef9106.
//
// Solidity: function hasPermission(address _who, address _where, bytes32 _what, bytes _how) view returns(bool)
func (_ACL *ACLCaller) HasPermission1(opts *bind.CallOpts, _who common.Address, _where common.Address, _what [32]byte, _how []byte) (bool, error) {
	var out []interface{}
	err := _ACL.contract.Call(opts, &out, "hasPermission1", _who, _where, _what, _how)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPermission1 is a free data retrieval call binding the contract method 0xfdef9106.
//
// Solidity: function hasPermission(address _who, address _where, bytes32 _what, bytes _how) view returns(bool)
func (_ACL *ACLSession) HasPermission1(_who common.Address, _where common.Address, _what [32]byte, _how []byte) (bool, error) {
	return _ACL.Contract.HasPermission1(&_ACL.CallOpts, _who, _where, _what, _how)
}

// HasPermission1 is a free data retrieval call binding the contract method 0xfdef9106.
//
// Solidity: function hasPermission(address _who, address _where, bytes32 _what, bytes _how) view returns(bool)
func (_ACL *ACLCallerSession) HasPermission1(_who common.Address, _where common.Address, _what [32]byte, _how []byte) (bool, error) {
	return _ACL.Contract.HasPermission1(&_ACL.CallOpts, _who, _where, _what, _how)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ACL *ACLCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ACL.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ACL *ACLSession) IsOwner() (bool, error) {
	return _ACL.Contract.IsOwner(&_ACL.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ACL *ACLCallerSession) IsOwner() (bool, error) {
	return _ACL.Contract.IsOwner(&_ACL.CallOpts)
}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_ACL *ACLCaller) IsPetrified(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ACL.contract.Call(opts, &out, "isPetrified")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_ACL *ACLSession) IsPetrified() (bool, error) {
	return _ACL.Contract.IsPetrified(&_ACL.CallOpts)
}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_ACL *ACLCallerSession) IsPetrified() (bool, error) {
	return _ACL.Contract.IsPetrified(&_ACL.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ACL *ACLCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ACL.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ACL *ACLSession) Owner() (common.Address, error) {
	return _ACL.Contract.Owner(&_ACL.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ACL *ACLCallerSession) Owner() (common.Address, error) {
	return _ACL.Contract.Owner(&_ACL.CallOpts)
}

// BurnPermissionManager is a paid mutator transaction binding the contract method 0x09699ff5.
//
// Solidity: function burnPermissionManager(address _app, bytes32 _role) returns()
func (_ACL *ACLTransactor) BurnPermissionManager(opts *bind.TransactOpts, _app common.Address, _role [32]byte) (*types.Transaction, error) {
	return _ACL.contract.Transact(opts, "burnPermissionManager", _app, _role)
}

// BurnPermissionManager is a paid mutator transaction binding the contract method 0x09699ff5.
//
// Solidity: function burnPermissionManager(address _app, bytes32 _role) returns()
func (_ACL *ACLSession) BurnPermissionManager(_app common.Address, _role [32]byte) (*types.Transaction, error) {
	return _ACL.Contract.BurnPermissionManager(&_ACL.TransactOpts, _app, _role)
}

// BurnPermissionManager is a paid mutator transaction binding the contract method 0x09699ff5.
//
// Solidity: function burnPermissionManager(address _app, bytes32 _role) returns()
func (_ACL *ACLTransactorSession) BurnPermissionManager(_app common.Address, _role [32]byte) (*types.Transaction, error) {
	return _ACL.Contract.BurnPermissionManager(&_ACL.TransactOpts, _app, _role)
}

// CreateBurnedPermission is a paid mutator transaction binding the contract method 0x0808343e.
//
// Solidity: function createBurnedPermission(address _app, bytes32 _role) returns()
func (_ACL *ACLTransactor) CreateBurnedPermission(opts *bind.TransactOpts, _app common.Address, _role [32]byte) (*types.Transaction, error) {
	return _ACL.contract.Transact(opts, "createBurnedPermission", _app, _role)
}

// CreateBurnedPermission is a paid mutator transaction binding the contract method 0x0808343e.
//
// Solidity: function createBurnedPermission(address _app, bytes32 _role) returns()
func (_ACL *ACLSession) CreateBurnedPermission(_app common.Address, _role [32]byte) (*types.Transaction, error) {
	return _ACL.Contract.CreateBurnedPermission(&_ACL.TransactOpts, _app, _role)
}

// CreateBurnedPermission is a paid mutator transaction binding the contract method 0x0808343e.
//
// Solidity: function createBurnedPermission(address _app, bytes32 _role) returns()
func (_ACL *ACLTransactorSession) CreateBurnedPermission(_app common.Address, _role [32]byte) (*types.Transaction, error) {
	return _ACL.Contract.CreateBurnedPermission(&_ACL.TransactOpts, _app, _role)
}

// CreatePermission is a paid mutator transaction binding the contract method 0xbe038478.
//
// Solidity: function createPermission(address _entity, address _app, bytes32 _role, address _manager) returns()
func (_ACL *ACLTransactor) CreatePermission(opts *bind.TransactOpts, _entity common.Address, _app common.Address, _role [32]byte, _manager common.Address) (*types.Transaction, error) {
	return _ACL.contract.Transact(opts, "createPermission", _entity, _app, _role, _manager)
}

// CreatePermission is a paid mutator transaction binding the contract method 0xbe038478.
//
// Solidity: function createPermission(address _entity, address _app, bytes32 _role, address _manager) returns()
func (_ACL *ACLSession) CreatePermission(_entity common.Address, _app common.Address, _role [32]byte, _manager common.Address) (*types.Transaction, error) {
	return _ACL.Contract.CreatePermission(&_ACL.TransactOpts, _entity, _app, _role, _manager)
}

// CreatePermission is a paid mutator transaction binding the contract method 0xbe038478.
//
// Solidity: function createPermission(address _entity, address _app, bytes32 _role, address _manager) returns()
func (_ACL *ACLTransactorSession) CreatePermission(_entity common.Address, _app common.Address, _role [32]byte, _manager common.Address) (*types.Transaction, error) {
	return _ACL.Contract.CreatePermission(&_ACL.TransactOpts, _entity, _app, _role, _manager)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x0a8ed3db.
//
// Solidity: function grantPermission(address _entity, address _app, bytes32 _role) returns()
func (_ACL *ACLTransactor) GrantPermission(opts *bind.TransactOpts, _entity common.Address, _app common.Address, _role [32]byte) (*types.Transaction, error) {
	return _ACL.contract.Transact(opts, "grantPermission", _entity, _app, _role)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x0a8ed3db.
//
// Solidity: function grantPermission(address _entity, address _app, bytes32 _role) returns()
func (_ACL *ACLSession) GrantPermission(_entity common.Address, _app common.Address, _role [32]byte) (*types.Transaction, error) {
	return _ACL.Contract.GrantPermission(&_ACL.TransactOpts, _entity, _app, _role)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x0a8ed3db.
//
// Solidity: function grantPermission(address _entity, address _app, bytes32 _role) returns()
func (_ACL *ACLTransactorSession) GrantPermission(_entity common.Address, _app common.Address, _role [32]byte) (*types.Transaction, error) {
	return _ACL.Contract.GrantPermission(&_ACL.TransactOpts, _entity, _app, _role)
}

// GrantPermissionP is a paid mutator transaction binding the contract method 0x6815c992.
//
// Solidity: function grantPermissionP(address _entity, address _app, bytes32 _role, uint256[] _params) returns()
func (_ACL *ACLTransactor) GrantPermissionP(opts *bind.TransactOpts, _entity common.Address, _app common.Address, _role [32]byte, _params []*big.Int) (*types.Transaction, error) {
	return _ACL.contract.Transact(opts, "grantPermissionP", _entity, _app, _role, _params)
}

// GrantPermissionP is a paid mutator transaction binding the contract method 0x6815c992.
//
// Solidity: function grantPermissionP(address _entity, address _app, bytes32 _role, uint256[] _params) returns()
func (_ACL *ACLSession) GrantPermissionP(_entity common.Address, _app common.Address, _role [32]byte, _params []*big.Int) (*types.Transaction, error) {
	return _ACL.Contract.GrantPermissionP(&_ACL.TransactOpts, _entity, _app, _role, _params)
}

// GrantPermissionP is a paid mutator transaction binding the contract method 0x6815c992.
//
// Solidity: function grantPermissionP(address _entity, address _app, bytes32 _role, uint256[] _params) returns()
func (_ACL *ACLTransactorSession) GrantPermissionP(_entity common.Address, _app common.Address, _role [32]byte, _params []*big.Int) (*types.Transaction, error) {
	return _ACL.Contract.GrantPermissionP(&_ACL.TransactOpts, _entity, _app, _role, _params)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _permissionsCreator) returns()
func (_ACL *ACLTransactor) Initialize(opts *bind.TransactOpts, _permissionsCreator common.Address) (*types.Transaction, error) {
	return _ACL.contract.Transact(opts, "initialize", _permissionsCreator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _permissionsCreator) returns()
func (_ACL *ACLSession) Initialize(_permissionsCreator common.Address) (*types.Transaction, error) {
	return _ACL.Contract.Initialize(&_ACL.TransactOpts, _permissionsCreator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _permissionsCreator) returns()
func (_ACL *ACLTransactorSession) Initialize(_permissionsCreator common.Address) (*types.Transaction, error) {
	return _ACL.Contract.Initialize(&_ACL.TransactOpts, _permissionsCreator)
}

// RemovePermissionManager is a paid mutator transaction binding the contract method 0xa885508a.
//
// Solidity: function removePermissionManager(address _app, bytes32 _role) returns()
func (_ACL *ACLTransactor) RemovePermissionManager(opts *bind.TransactOpts, _app common.Address, _role [32]byte) (*types.Transaction, error) {
	return _ACL.contract.Transact(opts, "removePermissionManager", _app, _role)
}

// RemovePermissionManager is a paid mutator transaction binding the contract method 0xa885508a.
//
// Solidity: function removePermissionManager(address _app, bytes32 _role) returns()
func (_ACL *ACLSession) RemovePermissionManager(_app common.Address, _role [32]byte) (*types.Transaction, error) {
	return _ACL.Contract.RemovePermissionManager(&_ACL.TransactOpts, _app, _role)
}

// RemovePermissionManager is a paid mutator transaction binding the contract method 0xa885508a.
//
// Solidity: function removePermissionManager(address _app, bytes32 _role) returns()
func (_ACL *ACLTransactorSession) RemovePermissionManager(_app common.Address, _role [32]byte) (*types.Transaction, error) {
	return _ACL.Contract.RemovePermissionManager(&_ACL.TransactOpts, _app, _role)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ACL *ACLTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ACL.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ACL *ACLSession) RenounceOwnership() (*types.Transaction, error) {
	return _ACL.Contract.RenounceOwnership(&_ACL.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ACL *ACLTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ACL.Contract.RenounceOwnership(&_ACL.TransactOpts)
}

// RevokePermission is a paid mutator transaction binding the contract method 0x9d0effdb.
//
// Solidity: function revokePermission(address _entity, address _app, bytes32 _role) returns()
func (_ACL *ACLTransactor) RevokePermission(opts *bind.TransactOpts, _entity common.Address, _app common.Address, _role [32]byte) (*types.Transaction, error) {
	return _ACL.contract.Transact(opts, "revokePermission", _entity, _app, _role)
}

// RevokePermission is a paid mutator transaction binding the contract method 0x9d0effdb.
//
// Solidity: function revokePermission(address _entity, address _app, bytes32 _role) returns()
func (_ACL *ACLSession) RevokePermission(_entity common.Address, _app common.Address, _role [32]byte) (*types.Transaction, error) {
	return _ACL.Contract.RevokePermission(&_ACL.TransactOpts, _entity, _app, _role)
}

// RevokePermission is a paid mutator transaction binding the contract method 0x9d0effdb.
//
// Solidity: function revokePermission(address _entity, address _app, bytes32 _role) returns()
func (_ACL *ACLTransactorSession) RevokePermission(_entity common.Address, _app common.Address, _role [32]byte) (*types.Transaction, error) {
	return _ACL.Contract.RevokePermission(&_ACL.TransactOpts, _entity, _app, _role)
}

// SetACL is a paid mutator transaction binding the contract method 0x76aad605.
//
// Solidity: function setACL(address _acl) returns()
func (_ACL *ACLTransactor) SetACL(opts *bind.TransactOpts, _acl common.Address) (*types.Transaction, error) {
	return _ACL.contract.Transact(opts, "setACL", _acl)
}

// SetACL is a paid mutator transaction binding the contract method 0x76aad605.
//
// Solidity: function setACL(address _acl) returns()
func (_ACL *ACLSession) SetACL(_acl common.Address) (*types.Transaction, error) {
	return _ACL.Contract.SetACL(&_ACL.TransactOpts, _acl)
}

// SetACL is a paid mutator transaction binding the contract method 0x76aad605.
//
// Solidity: function setACL(address _acl) returns()
func (_ACL *ACLTransactorSession) SetACL(_acl common.Address) (*types.Transaction, error) {
	return _ACL.Contract.SetACL(&_ACL.TransactOpts, _acl)
}

// SetPermissionManager is a paid mutator transaction binding the contract method 0xafd925df.
//
// Solidity: function setPermissionManager(address _newManager, address _app, bytes32 _role) returns()
func (_ACL *ACLTransactor) SetPermissionManager(opts *bind.TransactOpts, _newManager common.Address, _app common.Address, _role [32]byte) (*types.Transaction, error) {
	return _ACL.contract.Transact(opts, "setPermissionManager", _newManager, _app, _role)
}

// SetPermissionManager is a paid mutator transaction binding the contract method 0xafd925df.
//
// Solidity: function setPermissionManager(address _newManager, address _app, bytes32 _role) returns()
func (_ACL *ACLSession) SetPermissionManager(_newManager common.Address, _app common.Address, _role [32]byte) (*types.Transaction, error) {
	return _ACL.Contract.SetPermissionManager(&_ACL.TransactOpts, _newManager, _app, _role)
}

// SetPermissionManager is a paid mutator transaction binding the contract method 0xafd925df.
//
// Solidity: function setPermissionManager(address _newManager, address _app, bytes32 _role) returns()
func (_ACL *ACLTransactorSession) SetPermissionManager(_newManager common.Address, _app common.Address, _role [32]byte) (*types.Transaction, error) {
	return _ACL.Contract.SetPermissionManager(&_ACL.TransactOpts, _newManager, _app, _role)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ACL *ACLTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ACL.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ACL *ACLSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ACL.Contract.TransferOwnership(&_ACL.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ACL *ACLTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ACL.Contract.TransferOwnership(&_ACL.TransactOpts, newOwner)
}

// ACLACLAddrIterator is returned from FilterACLAddr and is used to iterate over the raw logs and unpacked data for ACLAddr events raised by the ACL contract.
type ACLACLAddrIterator struct {
	Event *ACLACLAddr // Event containing the contract specifics and raw log

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
func (it *ACLACLAddrIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ACLACLAddr)
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
		it.Event = new(ACLACLAddr)
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
func (it *ACLACLAddrIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ACLACLAddrIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ACLACLAddr represents a ACLAddr event raised by the ACL contract.
type ACLACLAddr struct {
	Msgsender    common.Address
	Thiscontract common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterACLAddr is a free log retrieval operation binding the contract event 0xa5b2085db613b901016a287a140e5615acb5e5abe779da5a2e1522607aa69a20.
//
// Solidity: event ACLAddr(address indexed msgsender, address indexed thiscontract)
func (_ACL *ACLFilterer) FilterACLAddr(opts *bind.FilterOpts, msgsender []common.Address, thiscontract []common.Address) (*ACLACLAddrIterator, error) {

	var msgsenderRule []interface{}
	for _, msgsenderItem := range msgsender {
		msgsenderRule = append(msgsenderRule, msgsenderItem)
	}
	var thiscontractRule []interface{}
	for _, thiscontractItem := range thiscontract {
		thiscontractRule = append(thiscontractRule, thiscontractItem)
	}

	logs, sub, err := _ACL.contract.FilterLogs(opts, "ACLAddr", msgsenderRule, thiscontractRule)
	if err != nil {
		return nil, err
	}
	return &ACLACLAddrIterator{contract: _ACL.contract, event: "ACLAddr", logs: logs, sub: sub}, nil
}

// WatchACLAddr is a free log subscription operation binding the contract event 0xa5b2085db613b901016a287a140e5615acb5e5abe779da5a2e1522607aa69a20.
//
// Solidity: event ACLAddr(address indexed msgsender, address indexed thiscontract)
func (_ACL *ACLFilterer) WatchACLAddr(opts *bind.WatchOpts, sink chan<- *ACLACLAddr, msgsender []common.Address, thiscontract []common.Address) (event.Subscription, error) {

	var msgsenderRule []interface{}
	for _, msgsenderItem := range msgsender {
		msgsenderRule = append(msgsenderRule, msgsenderItem)
	}
	var thiscontractRule []interface{}
	for _, thiscontractItem := range thiscontract {
		thiscontractRule = append(thiscontractRule, thiscontractItem)
	}

	logs, sub, err := _ACL.contract.WatchLogs(opts, "ACLAddr", msgsenderRule, thiscontractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ACLACLAddr)
				if err := _ACL.contract.UnpackLog(event, "ACLAddr", log); err != nil {
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

// ParseACLAddr is a log parse operation binding the contract event 0xa5b2085db613b901016a287a140e5615acb5e5abe779da5a2e1522607aa69a20.
//
// Solidity: event ACLAddr(address indexed msgsender, address indexed thiscontract)
func (_ACL *ACLFilterer) ParseACLAddr(log types.Log) (*ACLACLAddr, error) {
	event := new(ACLACLAddr)
	if err := _ACL.contract.UnpackLog(event, "ACLAddr", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ACLChangePermissionManagerIterator is returned from FilterChangePermissionManager and is used to iterate over the raw logs and unpacked data for ChangePermissionManager events raised by the ACL contract.
type ACLChangePermissionManagerIterator struct {
	Event *ACLChangePermissionManager // Event containing the contract specifics and raw log

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
func (it *ACLChangePermissionManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ACLChangePermissionManager)
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
		it.Event = new(ACLChangePermissionManager)
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
func (it *ACLChangePermissionManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ACLChangePermissionManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ACLChangePermissionManager represents a ChangePermissionManager event raised by the ACL contract.
type ACLChangePermissionManager struct {
	App     common.Address
	Role    [32]byte
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChangePermissionManager is a free log retrieval operation binding the contract event 0xf3addc8b8e25ee11528a61b0e65092cae0666ef0ec0c64cb303993c88d689b4d.
//
// Solidity: event ChangePermissionManager(address indexed app, bytes32 indexed role, address indexed manager)
func (_ACL *ACLFilterer) FilterChangePermissionManager(opts *bind.FilterOpts, app []common.Address, role [][32]byte, manager []common.Address) (*ACLChangePermissionManagerIterator, error) {

	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _ACL.contract.FilterLogs(opts, "ChangePermissionManager", appRule, roleRule, managerRule)
	if err != nil {
		return nil, err
	}
	return &ACLChangePermissionManagerIterator{contract: _ACL.contract, event: "ChangePermissionManager", logs: logs, sub: sub}, nil
}

// WatchChangePermissionManager is a free log subscription operation binding the contract event 0xf3addc8b8e25ee11528a61b0e65092cae0666ef0ec0c64cb303993c88d689b4d.
//
// Solidity: event ChangePermissionManager(address indexed app, bytes32 indexed role, address indexed manager)
func (_ACL *ACLFilterer) WatchChangePermissionManager(opts *bind.WatchOpts, sink chan<- *ACLChangePermissionManager, app []common.Address, role [][32]byte, manager []common.Address) (event.Subscription, error) {

	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _ACL.contract.WatchLogs(opts, "ChangePermissionManager", appRule, roleRule, managerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ACLChangePermissionManager)
				if err := _ACL.contract.UnpackLog(event, "ChangePermissionManager", log); err != nil {
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

// ParseChangePermissionManager is a log parse operation binding the contract event 0xf3addc8b8e25ee11528a61b0e65092cae0666ef0ec0c64cb303993c88d689b4d.
//
// Solidity: event ChangePermissionManager(address indexed app, bytes32 indexed role, address indexed manager)
func (_ACL *ACLFilterer) ParseChangePermissionManager(log types.Log) (*ACLChangePermissionManager, error) {
	event := new(ACLChangePermissionManager)
	if err := _ACL.contract.UnpackLog(event, "ChangePermissionManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ACLOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ACL contract.
type ACLOwnershipTransferredIterator struct {
	Event *ACLOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ACLOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ACLOwnershipTransferred)
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
		it.Event = new(ACLOwnershipTransferred)
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
func (it *ACLOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ACLOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ACLOwnershipTransferred represents a OwnershipTransferred event raised by the ACL contract.
type ACLOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ACL *ACLFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ACLOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ACL.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ACLOwnershipTransferredIterator{contract: _ACL.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ACL *ACLFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ACLOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ACL.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ACLOwnershipTransferred)
				if err := _ACL.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ACL *ACLFilterer) ParseOwnershipTransferred(log types.Log) (*ACLOwnershipTransferred, error) {
	event := new(ACLOwnershipTransferred)
	if err := _ACL.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ACLSetPermissionIterator is returned from FilterSetPermission and is used to iterate over the raw logs and unpacked data for SetPermission events raised by the ACL contract.
type ACLSetPermissionIterator struct {
	Event *ACLSetPermission // Event containing the contract specifics and raw log

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
func (it *ACLSetPermissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ACLSetPermission)
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
		it.Event = new(ACLSetPermission)
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
func (it *ACLSetPermissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ACLSetPermissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ACLSetPermission represents a SetPermission event raised by the ACL contract.
type ACLSetPermission struct {
	Entity  common.Address
	App     common.Address
	Role    [32]byte
	Allowed bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetPermission is a free log retrieval operation binding the contract event 0x759b9a74d5354b5801710a0c1b283cc9f0d32b607ac8ced10c83ac8e75c77d52.
//
// Solidity: event SetPermission(address indexed entity, address indexed app, bytes32 indexed role, bool allowed)
func (_ACL *ACLFilterer) FilterSetPermission(opts *bind.FilterOpts, entity []common.Address, app []common.Address, role [][32]byte) (*ACLSetPermissionIterator, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}
	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _ACL.contract.FilterLogs(opts, "SetPermission", entityRule, appRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &ACLSetPermissionIterator{contract: _ACL.contract, event: "SetPermission", logs: logs, sub: sub}, nil
}

// WatchSetPermission is a free log subscription operation binding the contract event 0x759b9a74d5354b5801710a0c1b283cc9f0d32b607ac8ced10c83ac8e75c77d52.
//
// Solidity: event SetPermission(address indexed entity, address indexed app, bytes32 indexed role, bool allowed)
func (_ACL *ACLFilterer) WatchSetPermission(opts *bind.WatchOpts, sink chan<- *ACLSetPermission, entity []common.Address, app []common.Address, role [][32]byte) (event.Subscription, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}
	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _ACL.contract.WatchLogs(opts, "SetPermission", entityRule, appRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ACLSetPermission)
				if err := _ACL.contract.UnpackLog(event, "SetPermission", log); err != nil {
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

// ParseSetPermission is a log parse operation binding the contract event 0x759b9a74d5354b5801710a0c1b283cc9f0d32b607ac8ced10c83ac8e75c77d52.
//
// Solidity: event SetPermission(address indexed entity, address indexed app, bytes32 indexed role, bool allowed)
func (_ACL *ACLFilterer) ParseSetPermission(log types.Log) (*ACLSetPermission, error) {
	event := new(ACLSetPermission)
	if err := _ACL.contract.UnpackLog(event, "SetPermission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ACLSetPermissionParamsIterator is returned from FilterSetPermissionParams and is used to iterate over the raw logs and unpacked data for SetPermissionParams events raised by the ACL contract.
type ACLSetPermissionParamsIterator struct {
	Event *ACLSetPermissionParams // Event containing the contract specifics and raw log

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
func (it *ACLSetPermissionParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ACLSetPermissionParams)
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
		it.Event = new(ACLSetPermissionParams)
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
func (it *ACLSetPermissionParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ACLSetPermissionParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ACLSetPermissionParams represents a SetPermissionParams event raised by the ACL contract.
type ACLSetPermissionParams struct {
	Entity     common.Address
	App        common.Address
	Role       [32]byte
	ParamsHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetPermissionParams is a free log retrieval operation binding the contract event 0x8dfee25d92d73b8c9b868f9fa3e215cc1981033f426e53803e3da4f09a2cfc30.
//
// Solidity: event SetPermissionParams(address indexed entity, address indexed app, bytes32 indexed role, bytes32 paramsHash)
func (_ACL *ACLFilterer) FilterSetPermissionParams(opts *bind.FilterOpts, entity []common.Address, app []common.Address, role [][32]byte) (*ACLSetPermissionParamsIterator, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}
	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _ACL.contract.FilterLogs(opts, "SetPermissionParams", entityRule, appRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &ACLSetPermissionParamsIterator{contract: _ACL.contract, event: "SetPermissionParams", logs: logs, sub: sub}, nil
}

// WatchSetPermissionParams is a free log subscription operation binding the contract event 0x8dfee25d92d73b8c9b868f9fa3e215cc1981033f426e53803e3da4f09a2cfc30.
//
// Solidity: event SetPermissionParams(address indexed entity, address indexed app, bytes32 indexed role, bytes32 paramsHash)
func (_ACL *ACLFilterer) WatchSetPermissionParams(opts *bind.WatchOpts, sink chan<- *ACLSetPermissionParams, entity []common.Address, app []common.Address, role [][32]byte) (event.Subscription, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}
	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _ACL.contract.WatchLogs(opts, "SetPermissionParams", entityRule, appRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ACLSetPermissionParams)
				if err := _ACL.contract.UnpackLog(event, "SetPermissionParams", log); err != nil {
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

// ParseSetPermissionParams is a log parse operation binding the contract event 0x8dfee25d92d73b8c9b868f9fa3e215cc1981033f426e53803e3da4f09a2cfc30.
//
// Solidity: event SetPermissionParams(address indexed entity, address indexed app, bytes32 indexed role, bytes32 paramsHash)
func (_ACL *ACLFilterer) ParseSetPermissionParams(log types.Log) (*ACLSetPermissionParams, error) {
	event := new(ACLSetPermissionParams)
	if err := _ACL.contract.UnpackLog(event, "SetPermissionParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
