/*
@Time : 2019-04-29 17:00 
@Author : xiaoyueya
@File : node_data
@Software: GoLand
*/
package gotrie

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
