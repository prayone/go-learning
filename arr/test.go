package arr

import "fmt"

// 创建 user
type User struct {
	ID    int
	Name  string
	job   string
	Score int
}

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

func GetUser() []User {
	users := []User{
		User{ID: 1, Name: "王萍", job: "全栈开发", Score: 90},
		User{ID: 2, Name: "李明", job: "全栈开发", Score: 80},
		User{ID: 3, Name: "小黄", job: "全栈开发", Score: 70},
		User{ID: 4, Name: "二名", job: "全栈开发", Score: 60},
		User{ID: 5, Name: "小李", job: "全栈开发", Score: 50},
	}
	users = append(users, User{ID: 6, Name: "小红", job: "全栈开发", Score: 40})
	return users
}

func FindUserByID(users []User, id int) (User, bool) {
	for _, user := range users {
		if user.ID == id {
			return user, true
		}
	}
	return User{}, false
}

func GoFind() {
	user, exist := FindUserByID(GetUser(), 1)
	if exist {
		fmt.Println(user)
	}
}
func FindHighScore() {
	users := GetUser()
	for _, user := range users {
		if user.Score >= 70 {
			fmt.Println(user)
		}
	}
}
func EditScore() {
	users := GetUser()
	for _, user := range users {
		if user.ID == 2 {
			fmt.Println(user)
			user.Score = 75
		}
	}
	fmt.Println(users)
}
func StatLeavel() {
	users := GetUser()
	leavelCount := map[string]int{}
	for _, user := range users {
		level := getLevel(user.Score)
		leavelCount[level]++ // map里面不存在这个key默认是0，直接++即可
	}
	fmt.Println(leavelCount)
}
