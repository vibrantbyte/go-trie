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

	urlLibrary.AddUrl("/home/a")
	urlLibrary.AddUrl("/home/b")
	urlLibrary.AddUrl("/home/c")
	urlLibrary.AddUrl("/home/e/*")
	urlLibrary.AddUrl("/home/f/**")
	urlLibrary.AddUrl("/home/c/?")
	urlLibrary.AddUrl("/home/c/{ss}/1000")
	urlLibrary.AddUrl("/home/c/{ss}/1000")


	urls := urlLibrary.Matcher("/home/c")
	if urls != nil {
		for index:= range urls  {
			fmt.Println(*urls[index])
		}
	}

	urlLibrary.RemoveUrl("/home/c/{ss}/1000")
	urls = urlLibrary.Matcher("/home/c")
	if urls != nil {
		for index:= range urls  {
			value := urls[index]
			if value != nil{
				fmt.Println(*value)
			}
		}
	}

	urlLibrary.AddUrl("/home/c/{ss}/1000")
	urls = urlLibrary.Matcher("/home/c")
	if urls != nil {
		for index:= range urls  {
			value := urls[index]
			if value != nil{
				fmt.Println(*value)
			}
		}
	}



	fmt.Print(urlLibrary)


}
