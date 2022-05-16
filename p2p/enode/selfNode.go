package enode

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