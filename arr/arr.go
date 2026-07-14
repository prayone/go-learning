package arr

import "fmt"

func Arr() {
	var n [3]int
	var i, j int
	for i = 0; i < 3; i++ {
		n[i] = i + 100
	}
	for j = 0; j < 3; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}

func Arr2() {
	var score = [3]int{80, 90, 100} //数组：长度固定
	for index, score := range score {
		fmt.Println(index, score)
	}
}

// 切片
func Slice() {
	var names = []string{"王萍", "张三", "李四"} //切片：长度不固定
	names = append(names, "赵六")
	// fmt.Println(names[0])
	// fmt.Println(names[3])
	// fmt.Println(len(names))
	for _, name := range names {
		fmt.Println(name)
	}
}

func Map() {
	//map[string]string    表示：key 是 string value 也是 string
	user := map[string]string{"name": "wp", "job": "全栈开发"}
	user["city"] = "北京"

	fmt.Println(user["name"] + " " + user["job"] + " " + user["city"])
	delete(user, "city")
	fmt.Println(user["city"])
	val, exist := user["city"] //两个返回值  val 为值，exist 为是否存在
	fmt.Println(val, exist)

	for k, v := range user {
		fmt.Println(k, v)
	}

}

// 结构体
func Struct() {
	type User struct {
		Name string //注意大写 表示这个字段可以被其他包访问。
		Age  int
		Job  string
	}
	user := User{Name: "wp", Age: 30, Job: "全栈开发"}
	fmt.Println(user.Name, user.Age, user.Job)

}
