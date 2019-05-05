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
	"sync"
	"sync/atomic"
)

//基于性能考虑，并且严格限定
const TreeMaxLength  = 1 << 16

//UrlLibrary url存储库
type UrlLibrary struct {
	//root 根节点
	root *TrieNode
	//存储器的大小
	len int64
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
	if lib.GetLen() > TreeMaxLength {
		return
	}
	if utils.HasText(url){
		//截取URL
		urlSegment := utils.TokenizeToStringArray(url,utils.Spliter,true,true)
		if urlSegment != nil {
			//判断根节点
			if lib.root == nil {
				lib.root = new(TrieNode)
			}
			//存储过程变量
			temp := lib.root
			pType := NormalPath
			final := 0

			for index := range urlSegment {
				final = index
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
				temp.Data = data
			}
			//转化
			data := temp.Data.(*NodeData)
			data.AddUrl(urlSegment[final:],pType)
		}
		atomic.AddInt64(&lib.len,1)
	}
}

//RemoveUrl 删除url
func (lib *UrlLibrary) RemoveUrl(url string) {
	matcher,_,other := lib.matcherUrl(url)
	if matcher == nil || matcher.Data == nil {
		return
	}
	//删除url
	data := matcher.Data.(*NodeData)
	if data.RemoveUrl(other){
		atomic.AddInt64(&lib.len,-1)
	}
}

//Matcher
func (lib *UrlLibrary) Matcher(url string) []*string {
	matcher,prefix,_ := lib.matcherUrl(url)
	if matcher == nil || matcher.Data == nil {
		return nil
	}
	data := matcher.Data.(*NodeData)
	return data.GetUrlList(prefix)
}

//GetLen 获取存储大小
func (lib *UrlLibrary) GetLen() int64{
	return atomic.LoadInt64(&lib.len)
}

//GetHost 获取host
func (lib *UrlLibrary) GetHost() string{
	return lib.host
}

//recursionInsertUrl
func (lib *UrlLibrary) recursionInsertUrl(node *TrieNode,segment *string) *TrieNode {
	var tmp *TrieNode
	if atomic.LoadInt64(&node.Degree) > 0 {
		//开始循环判断是否存在
		node.Child.Range(func(key, value interface{}) bool {
			if strings.EqualFold(*segment,key.(string)) {
				//定位到指定节点
				tmp = value.(*TrieNode)
				return false
			}
			return true
		})
		//如果定位到则返回
		if tmp != nil {
			return tmp
		}
	}else{
		//直接创建
		node.Child =  new(sync.Map)
	}
	tmp = new(TrieNode)
	//如果不存在创建
	node.Child.Store(*segment,tmp)
	atomic.AddInt64(&node.Degree,1)
	return tmp
}

//matcherUrl
func (lib *UrlLibrary) matcherUrl(url string) (*TrieNode,*string,[]*string) {
	prefixUrl := utils.EmptyString
	if utils.HasText(url){
		//截取URL
		urlSegment := utils.TokenizeToStringArray(url,utils.Spliter,true,true)
		if urlSegment != nil {
			//判断根节点
			if lib.root == nil {
				return nil,&prefixUrl,nil
			}
			temp := lib.root
			final := 0
			for index := range urlSegment {
				final = index
				strAddress := urlSegment[index]
				//处理本节点及后续节点
				if utils.GlobPattern.MatchString(*strAddress) {
					break
				}
				//child
				child,ok := temp.Child.Load(*strAddress)
				if !ok || child == nil {
					break
				}
				//进行前缀拼接
				prefixUrl += utils.Spliter
				prefixUrl += *strAddress

				//获取下一个节点
				temp = child.(*TrieNode)
			}

			//otherUrl
			return temp,&prefixUrl,urlSegment[final:]
		}
	}
	return nil,&prefixUrl,nil
}
