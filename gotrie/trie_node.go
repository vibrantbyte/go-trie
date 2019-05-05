/*
@Time : 2019-04-29 15:30 
@Author : xiaoyueya
@File : trie_node
@Software: GoLand
*/
package gotrie

import "sync"

//TrieNode 树节点
type TrieNode struct {
	//孩子节点
	Child *sync.Map
	//孩子节点的度
	Degree int64
	//节点数据
	Data interface{}
}