package main

import (
	"fmt"
	"strings"
)

func main() {


	var str = " _hellowold_ "
	// 是否存在
	fmt.Println(strings.Contains(str, "do"))
	// 存在个数
	fmt.Println(strings.Count(str, "o"))
	// 不关心大小写的 是否相等
	fmt.Println(strings.EqualFold(str, "HELLOWOLD"))
	// 存在的下标
	fmt.Println(strings.Index(str,"d1"))
	// 替换 字符串 匹配的字符 替换的字符 替换的个数 -1为全部
	fmt.Println(strings.Replace(str, "l", "1", -1))
	// 切割成字符串数组
	fmt.Println(strings.Split(str, ""))
	// 切换成大写
	fmt.Println(strings.ToUpper(str))
	// 切换小写
	fmt.Println(strings.ToLower(strings.ToUpper(str)))
	// 去掉两端的空格
	fmt.Println(strings.TrimSpace(str))
	// 去掉两端指定的字符
	fmt.Println(strings.Trim(str, " _"))
	fmt.Println(strings.TrimLeft(str, " _"))
	fmt.Println(strings.TrimRight(str, "_ "))
	// 是否以什么开头
	fmt.Println(strings.HasPrefix(str, " _"))
	// 是否以什么结尾
	fmt.Println(strings.HasSuffix(str, "_ "))

}
