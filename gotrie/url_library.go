/*
@Time : 2019-04-29 15:30 
@Author : xiaoyueya
@File : url_library
@Software: GoLand
*/
package gotrie

import (
	"strings"
)

//UrlLibrary url存储库
type UrlLibrary struct {
	//root 根节点
	root *TrieNode
	//存储器的大小
	len int
	//存储的域名
	host string
}

//CreateUrlLibrary
func CreateUrlLibrary(host string) *UrlLibrary{
	urlLibrary := new(UrlLibrary)
	urlLibrary.host = host
	urlLibrary.len = 0
	urlLibrary.root = nil
	return urlLibrary
}

//AddUrl 添加url
func (lib *UrlLibrary) AddUrl(url string) {
	if HasText(url){
		//截取URL
		urlSegment := TokenizeToStringArray(url,Spliter,true,true)
		if urlSegment != nil {
			//判断根节点
			if lib.root == nil {
				lib.root = new(TrieNode)
			}
			temp := lib.root
			for index := range urlSegment {
				temp = lib.recursionInsertUrl(temp,urlSegment[index])
			}
			//存储数据
			if temp.Data == nil {
				data := new(NodeData)
				data.Len = 0
				data.UrlList = make([]*string,0)
				temp.Data = data
			}
			//转化
			data := temp.Data.(*NodeData)
			data.UrlList = append(data.UrlList,&url)
			data.PType = NormalPath
			data.Len ++
		}
		lib.len ++
	}
}

//GetLen 获取存储大小
func (lib *UrlLibrary) GetLen() int{
	return lib.len
}

//GetHost 获取host
func (lib *UrlLibrary) GetHost() string{
	return lib.host
}

//recursionInsertUrl
func (lib *UrlLibrary) recursionInsertUrl(node *TrieNode,segment *string) *TrieNode {
	if node.Degree > 0 {
		//开始循环判断是否存在
		for k,n := range node.Child {
			if strings.EqualFold(*segment,k) {
				return n
			}
		}
	}else{
		//直接创建
		node.Child = make(map[string]*TrieNode)
	}
	tmp := new(TrieNode)
	//如果不存在创建
	node.Child[*segment] = tmp
	node.Degree ++
	return tmp
}