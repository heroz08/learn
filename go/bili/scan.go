package main

import "fmt"


func main() {
	var age int
	var name string
	var isvip bool

	fmt.Println("输入姓名 年龄 和是否为vip 逗号分隔")
	
	// fmt.Println("输入姓名")
	// fmt.Scanln(&name)
	// fmt.Println("输入年龄")
	// fmt.Scanln(&age)
	// fmt.Println("输入是否为vip")
	// fmt.Scanln(&isvip)

	fmt.Scanf("%s-%d-%t",&name,&age,&isvip)

	fmt.Printf("姓名：%v 年龄： %v 是否为VIP：%v", name, age, isvip)
}