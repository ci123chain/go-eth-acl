package acl_sdk

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func HasPermission(ACLAddress, chain, contract, address, permission string) (bool, error) {
	//链地址
	conn, err := ethclient.Dial(chain)
	if err != nil {
		return false, err
	}

	//acl合约地址
	aclAddress := common.HexToAddress(ACLAddress)
	acl, err := NewACL(aclAddress, conn)
	if err != nil {
		return false, err
	}

	//查询用户地址
	user := common.HexToAddress(address)
	//权限点所在合约地址
	target := common.HexToAddress(contract)
	//权限点
	permissionBz, _ := hex.DecodeString(permission)

	var permission32 [32]byte
	copy(permission32[:], permissionBz)

	//调用acl合约查询是否有权限
	return acl.HasPermission(nil, user, target, permission32)
}
