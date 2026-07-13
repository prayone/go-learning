package main

import "fmt"

var day = 8

func getLevel(score int) string {
	if score >= 90 {
		return "专家"
	}
	if score >= 70 {
		return "进阶"
	}
	if score >= 40 {
		return "入门"
	}
	return "新手"
}
func getStudyMessage(days int) string {
	if days >= 7 {
		return "已经坚持学习一周了"
	}
	return "继续坚持学习"
}

func main() {
	fmt.Println("Hello, World!")
	name := "王萍"
	age := 30
	fmt.Printf("姓名：%s，年龄：%d\n", name, age)
	result()
	switch day {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	default:
		fmt.Println("其他日期")
	}
	res := getStudyMessage(day)
	fmt.Println(res)

}

func getUser() (string, int) {
	return "王萍", 30
}
func result() {
	name, age := getUser()
	fmt.Printf("getUser:"+"姓名：%s，年龄：%d\n", name, age)
}
