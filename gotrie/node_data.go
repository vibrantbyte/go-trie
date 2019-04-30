/*
@Time : 2019-04-29 17:00 
@Author : xiaoyueya
@File : node_data
@Software: GoLand
*/
package gotrie

import "strings"

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
	//模式
	PType PatternType
	//url存储
	UrlList []*string
}

//AddUrl
func (data *NodeData) AddUrl(url string,pType PatternType){
	if data.UrlList == nil {
		data.UrlList = make([]*string,0)
		data.PType = pType
	}
	data.Len ++
	data.UrlList = append(data.UrlList,&url)
}

//RemoveUrl
func (data *NodeData) RemoveUrl(url string) bool{
	if data.UrlList == nil {
		return false
	}
	remove := false
	len := len(data.UrlList)
	for index := range data.UrlList  {
		strAddress := data.UrlList[index]
		if strings.EqualFold(*strAddress,url) {
			remove = true
		}
		if remove {
			if len == index+1 {
				data.UrlList[index] = nil
				break
			}
			//移动元素
			if len > index {
				data.UrlList[index] = data.UrlList[index+1]
			}
		}

	}
	return remove
}
