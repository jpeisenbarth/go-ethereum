package enode

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type selfNode struct {
	node *LocalNode
}

var singleInstanceSelfNode *selfNode

func NewInstanceSelfNode(n *LocalNode) bool{
	if singleInstanceSelfNode != nil {
		return false
	}
	singleInstanceSelfNode = &selfNode{node: n}
	return true
}

func GetInstanceSelfNode() *selfNode {
	if singleInstanceSelfNode == nil {
		return nil
	}
	return singleInstanceSelfNode
}

// TODO stage
func (n selfNode) IsClose(b common.Hash) bool {
	a := crypto.Keccak256Hash(n.node.ID().Bytes())
	return IsClose(a, b)
}