/*
@Time : 2019-04-29 18:40 
@Author : xiaoyueya
@File : pattern_util_test
@Software: GoLand
*/
package utils

import "testing"

func TestIsPattern(t *testing.T) {

	t.Log(GlobPattern.MatchString("{name}"))
	t.Log(GlobPattern.MatchString("*"))
	t.Log(GlobPattern.MatchString("**"))
	t.Log(GlobPattern.MatchString("?"))
	t.Log(GlobPattern.MatchString("home"))

}
