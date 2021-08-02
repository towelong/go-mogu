package test

import (
	"fmt"
	"testing"
	"towelong/mogu/utils"
)

func TestMd5(t *testing.T) {
	// str := "102489283student3478cbbc33f84bd00d75d7dfa69e0daa"
	str2 := "AndroidSTARTa6870f0e4f8947f882f2d77fd302f0c6102489283北京市 · 北京市3478cbbc33f84bd00d75d7dfa69e0daa"
	md5Str := utils.CreateSign(str2)
	fmt.Println(md5Str)
	// planId a6870f0e4f8947f882f2d77fd302f0c6
}
