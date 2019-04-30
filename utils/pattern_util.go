/*
@Time : 2019-04-29 17:38 
@Author : xiaoyueya
@File : pattern_util
@Software: GoLand
*/
package utils

import (
	"regexp"
	"strings"
)

//GlobPattern
var GlobPattern *regexp.Regexp

//initial
func init() {
	//reg
	reg,err := regexp.Compile("\\?|\\*|\\{((?:\\{[^/]+?\\}|[^/{}]|\\\\[{}])+?)\\}")
	if err == nil {
		GlobPattern = reg
	}
}


//IsPattern 是pattern
func IsPattern(path string) bool {
	return strings.Index(path,"*")!=-1 || strings.Index(path,"?")!=-1
}
