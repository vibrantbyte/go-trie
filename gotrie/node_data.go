/*
@Time : 2019-04-29 17:00 
@Author : xiaoyueya
@File : node_data
@Software: GoLand
*/
package gotrie

import (
	"github.com/vibrantbyte/go-trie/utils"
	"strings"
)

type PatternType int

const(
	//NormalPath 正常模式
	NormalPath PatternType = 1
	//AntPath 规则
	AntPath PatternType = 2
	//RegexPath 正则模式
	RegexPath PatternType = 3
)

//NodeData 节点数据
type NodeData struct {
	//长度
	Len int
	//url存储
	UrlMap map[string]*UrlData
}

//UrlData url 存储类
type UrlData struct {
	Url *string
	PType PatternType
}


//AddUrl
func (data *NodeData) AddUrl(urlSegment []*string,pType PatternType){
	if data.UrlMap == nil {
		data.UrlMap = make(map[string]*UrlData)
	}
	data.Len ++
	//存储对象
	uData := new(UrlData)
	uData.PType = pType
	uData.Url = data.toUrl(urlSegment)
	//存储url对象
	data.UrlMap[*uData.Url] = uData
}

//RemoveUrl
func (data *NodeData) RemoveUrl(urlSegment []*string) bool{
	if data.UrlMap == nil {
		return false
	}
	remove := false
	url := data.toUrl(urlSegment)
	urlData := data.UrlMap[*url]
	if urlData == nil {
		return false
	}

	for key := range data.UrlMap  {
		if strings.EqualFold(key,*urlData.Url) {
			remove = true
			//删除url
			delete(data.UrlMap,key)
			break
		}
	}
	return remove
}

//GetUrlList
func (data *NodeData) GetUrlList(prefix *string) []*string {
	urlList := make([]*string,0)
	if data.UrlMap != nil {
		for key := range data.UrlMap {
			urlData := data.UrlMap[key]
			switch urlData.PType {
				case NormalPath:
					urlList = append(urlList, prefix)
					break
				case AntPath:
					url := *prefix + *urlData.Url
					urlList = append(urlList, &url)
					break
				case RegexPath:
					break
				default:
					break
			}
		}
	}
	return urlList
}

//toUrl
func (data *NodeData) toUrl(urlSegment []*string) *string{
	result := ""
	if urlSegment != nil {
		for index := range urlSegment {
			result += utils.Spliter
			result += *urlSegment[index]
		}
	}
	return &result
}
