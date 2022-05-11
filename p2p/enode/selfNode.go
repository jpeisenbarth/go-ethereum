package enode

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type selfNode struct {
	node *Node
}

var singleInstance *selfNode

func NewInstance(n *Node) bool{
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

// TODO stage :
// - retrouner bool et non pas la distance
// - description
// - finir la contidition de distance
func (n selfNode) IsClose(hash common.Hash) bool {
	distance := 2
	var rawBytes = make([]byte,distance)
	idHash := crypto.Keccak256(n.node.id[:])
	for i := 0; i < distance; i++ {
		rawBytes[i] = idHash[i] ^ hash[i]
	}
	// return big.NewInt(0).SetBytes(rawBytes[:])
	return false
}
