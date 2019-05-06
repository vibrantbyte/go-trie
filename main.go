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
	urlLibrary.AddUrl("/home/c/c")
	urlLibrary.AddUrl("/home/c/c/c")
	urlLibrary.AddUrl("/home/c/c/*")


	urls := urlLibrary.Matcher("/home/c/1/1000")
	if urls != nil {
		for index:= range urls  {
			fmt.Println(*urls[index])
		}
	}


	fmt.Println("-------------------------")
	urls = urlLibrary.Matcher("/home/c/c/1000")
	if urls != nil {
		for index:= range urls  {
			fmt.Println(*urls[index])
		}
	}


	fmt.Println("-------------------------")
	urls = urlLibrary.Matcher("/home/c/c/c/this")
	if urls != nil {
		for index:= range urls  {
			fmt.Println(*urls[index])
		}
	}


	fmt.Println("-------------------------")
	urls = urlLibrary.Matcher("/home/g")
	if urls != nil {
		for index:= range urls  {
			fmt.Println(*urls[index])
		}
	}


	fmt.Println("-------------------------")
	urls = urlLibrary.Matcher("/home/a")
	if urls != nil {
		for index:= range urls  {
			fmt.Println(*urls[index])
		}
	}

	fmt.Print(urlLibrary)


}
