/*
@Time : 2019-05-05 15:10 
@Author : xiaoyueya
@File : concurrent_test
@Software: GoLand
*/
package main

import (
	"fmt"
	"github.com/vibrantbyte/go-trie/gotrie"
	"math/rand"
	"strconv"
	"testing"
	"time"
)


var urlLibrary *gotrie.UrlLibrary

//init
func init(){
	urlLibrary = gotrie.CreateUrlLibrary("www.baidu.com")
}

//BenchmarkAddUrl 目前测试结果 千分之一秒
func BenchmarkAddUrl(b *testing.B){
	for i := 0; i < b.N; i++ { //use b.N for looping
		urlLibrary.AddUrl("/t/"+ strconv.Itoa(i))
	}
	b.Log(urlLibrary.GetLen())
}

//Benchmark_TimeConsumingFunction目前测试结果 千分之一秒
func Benchmark_TimeConsumingFunction(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能

	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		urlLibrary.AddUrl("/t/"+ strconv.Itoa(i))
	}
	b.Log(urlLibrary.GetLen())
}

//Benchmark_SearchFunction 测试查找效率
func Benchmark_SearchFunction(b *testing.B){
	b.StopTimer() //调用该函数停止压力测试的时间计数

	for j := 0; j< 5000 ; j++  {
		urlLibrary.AddUrl("/t/a/b/"+ strconv.Itoa(j))
	}

	b.StartTimer() //重新开始时间
	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		urlLibrary.Matcher("/t/a/b/"+ strconv.Itoa(1024))
	}
	b.Log(urlLibrary.GetLen())
}

//TestConcurrentTest
func TestConcurrentTest(t *testing.T)  {
	//多协程测试
	t.Log("启动写入协程")
	for i := 0; i< 100000 ;i++  {
		go wirteLib(t,i)
	}

	t.Log("启动读取协程")
	for i := 0; i< 100000 ;i++  {
		go func() {
			url := "/test/a/"+ strconv.Itoa(i)
			matcherUrls := urlLibrary.Matcher(url)
			if matcherUrls != nil {
				for index := range matcherUrls {
					t.Log(*matcherUrls[index])
				}
			}
		}()
	}


	time.Sleep(5 * time.Second)
	t.Log("总共写入 " + fmt.Sprint(urlLibrary.GetLen()) +" 条数据")
}


func wirteLib(t *testing.T,i int){
	//url := "/test/" + string(Krand(1, KC_RAND_KIND_LOWER)) + "/"+ strconv.Itoa(i)
	url := "/test/a/" + strconv.Itoa(i)
	urlLibrary.AddUrl(url)
}



const (
	KC_RAND_KIND_NUM   = 0  // 纯数字
	KC_RAND_KIND_LOWER = 1  // 小写字母
	KC_RAND_KIND_UPPER = 2  // 大写字母
	KC_RAND_KIND_ALL   = 3  // 数字、大小写字母
)

// 随机字符串
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i :=0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base+rand.Intn(scope))
	}
	return result
}