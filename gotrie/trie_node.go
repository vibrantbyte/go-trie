/*
@Time : 2019-04-29 15:30 
@Author : xiaoyueya
@File : trie_node
@Software: GoLand
*/
package gotrie

//TrieNode 树节点
type TrieNode struct {
	//孩子节点
	Child map[string]*TrieNode
	//孩子节点的度
	Degree int
	//节点数据
	Data interface{}
}

//Add
func (node *TrieNode) Add(key string,trie *TrieNode){
	if node.Degree == 0 {
		node.Child = make(map[string]*TrieNode)
	}
	node.Child[key] = trie
	node.Degree ++
}