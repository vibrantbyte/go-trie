/*
@Time : 2019-04-29 15:30 
@Author : xiaoyueya
@File : url_library
@Software: GoLand
*/
package gotrie

import (
	"github.com/vibrantbyte/go-trie/utils"
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
	if utils.HasText(url){
		//截取URL
		urlSegment := utils.TokenizeToStringArray(url,utils.Spliter,true,true)
		if urlSegment != nil {
			//判断根节点
			if lib.root == nil {
				lib.root = new(TrieNode)
			}
			temp := lib.root
			pType := NormalPath
			for index := range urlSegment {
				strAddress := urlSegment[index]
				if utils.GlobPattern.MatchString(*strAddress) {
					pType = AntPath
					break
				}
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
			data.PType = pType
			data.Len ++
		}
		lib.len ++
	}
}

//RemoveUrl 删除url
func (lib *UrlLibrary) RemoveUrl(url string) {
	matcher := lib.matcherUrl(url)
	if matcher == nil || matcher.Data == nil {
		return
	}
	//删除url
	data := matcher.Data.(*NodeData)
	lib.removeSlice(data.UrlList,url)
}

//Matcher
func (lib *UrlLibrary) Matcher(url string) []*string {
	matcher := lib.matcherUrl(url)
	if matcher == nil || matcher.Data == nil {
		return nil
	}
	data := matcher.Data.(*NodeData)
	return data.UrlList
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

//matcherUrl
func (lib *UrlLibrary) matcherUrl(url string) *TrieNode {
	if utils.HasText(url){
		//截取URL
		urlSegment := utils.TokenizeToStringArray(url,utils.Spliter,true,true)
		if urlSegment != nil {
			//判断根节点
			if lib.root == nil {
				return nil
			}
			temp := lib.root
			for index := range urlSegment {
				strAddress := urlSegment[index]
				if utils.GlobPattern.MatchString(*strAddress) {
					break
				}
				//child
				child := temp.Child[*strAddress]
				if child == nil {
					break
				}
				temp = temp.Child[*strAddress]
			}
			return temp
		}
	}
	return nil
}

//removeSlice
func (lib *UrlLibrary) removeSlice(source []*string,url string) []*string{
	if source == nil {
		return nil
	}
	remove := false
	len := len(source)
	for index := range source  {
		strAddress := source[index]
		if strings.EqualFold(*strAddress,url) {
			remove = true
		}
		if remove {
			if len == index+1 {
				source[index] = nil
				break
			}
			//移动元素
			if len > index {
				source[index] = source[index+1]
			}
		}

	}
	return source
}