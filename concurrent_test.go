/*
@Time : 2019-05-05 15:10 
@Author : xiaoyueya
@File : concurrent_test
@Software: GoLand
*/
package main

import (
	"github.com/vibrantbyte/go-trie/gotrie"
	"strconv"
	"testing"
)


var urlLibrary *gotrie.UrlLibrary

//init
func init(){
	urlLibrary = gotrie.CreateUrlLibrary("www.baidu.com")
}

func BenchmarkAddUrl(b *testing.B){
	for i := 0; i < b.N; i++ { //use b.N for looping
		urlLibrary.AddUrl("/t/"+ strconv.Itoa(i))
	}
}

func Benchmark_TimeConsumingFunction(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能

	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		urlLibrary.AddUrl("/t/"+ strconv.Itoa(i))
	}
}