package enode

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type selfNode struct {
	node *LocalNode
}

var singleInstance *selfNode

func NewInstance(n *LocalNode) bool{
	if singleInstance != nil {
		return false
	}
	singleInstance = &selfNode{node: n}
	return true
}

func GetInstance() *selfNode {
	if singleInstance == nil {
		return nil
	}
	return singleInstance
}

// TODO stage
func (n selfNode) IsClose(b common.Hash) bool {
	a := crypto.Keccak256Hash(n.node.ID().Bytes())
	return IsClose(a, b)
}