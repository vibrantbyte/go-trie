/*
@Time : 2019-04-29 15:31 
@Author : xiaoyueya
@File : main
@Software: GoLand
*/
package main

import (
	"fmt"
	"github.com/vibrantbyte/go-trie/gotrie"
)

func main(){
	fmt.Println("This is a trie.")

	urlLibrary := gotrie.CreateUrlLibrary("www.baidu.com")

	urlLibrary.AddUrl("/home/work/name/get")
	urlLibrary.AddUrl("/home/vibrant/name/get")
	urlLibrary.AddUrl("/home/vibrant/1/a")


	fmt.Print(urlLibrary)


}
