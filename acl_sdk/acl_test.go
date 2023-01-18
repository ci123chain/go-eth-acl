package acl_sdk

import (
	"testing"
)

const ACLAddress = "0xb2Db94b130819c57E030080489746892558fe7D9"

var chain = "https://weechain1.gw106.oneitfarm.com"
var vote_contract = "0xF1682F35353d404B0FfdE7A763474D3D91494bDB"
var address = "0x3F43E75Aaba2c2fD6E227C10C6E7DC125A93DE3c"
var create_vote_permission = "e7dcd7275292e064d090fbc5f3bd7995be23b502c1fed5cd94cfddbbdcd32bbc"
var modify_support_permission = "da3972983e62bdf826c4b807c4c9c2b8a941e1f83dfa76d53d6aeac11e1be650"

func TestHasPermission(t *testing.T) {
	//true
	hasCreateVotePermission, err := HasPermission(ACLAddress, chain, vote_contract, address, create_vote_permission)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(hasCreateVotePermission)

	//false
	hasModifySupportPermission, err := HasPermission(ACLAddress, chain, vote_contract, address, modify_support_permission)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(hasModifySupportPermission)
}